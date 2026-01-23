package swap

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/concurrent/wg"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"
)

func NewUniApolloClient() *UniApolloClient {
	return &UniApolloClient{
		//https://thegraph.com/hosted-service/subgraph/uniswap/uniswap-v3
		EntrypointFormat: "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-%s",
	}
}

type UniApolloClient struct {
	EntrypointFormat string
}

type GetTokenPriceParam struct {
	BlockNumber int64
	Token       string
}

type GetTokenPriceParams []GetTokenPriceParam

func (ps GetTokenPriceParams) Tokens() []string {

	tmp := mapset.NewSet[string]()

	for _, p := range ps {
		tmp.Add(p.Token)
	}

	return tmp.ToSlice()
}

func (ps GetTokenPriceParams) GroupByBlock() map[int64]GetTokenPriceParams {

	rsp := make(map[int64]GetTokenPriceParams, len(ps))

	for _, p := range ps {
		rsp[p.BlockNumber] = append(rsp[p.BlockNumber], p)
	}

	return rsp
}

type TokenPriceResult struct {
	BlockNumber int64
	Token       string
	PriceUSD    float64
}

type TokenPriceResults []TokenPriceResult

type TokenPriceResultMap map[string]float64

func (t TokenPriceResultMap) Find(block int64, token string) (float64, bool) {
	val, found := t[fmt.Sprintf("%d-%s", block, token)]
	return val, found
}

func (ts TokenPriceResults) ToMap() TokenPriceResultMap {
	rsp := make(TokenPriceResultMap, len(ts))

	for _, t := range ts {
		rsp[fmt.Sprintf("%d-%s", t.BlockNumber, t.Token)] = t.PriceUSD
	}

	return rsp
}

//func (t *UniApolloClient) GetTokenPriceSync(ctx context.Context, params GetTokenPriceParams) (TokenPriceResults, error) {
//	rsp := make(TokenPriceResults, 0, len(params))
//
//	blockMap := params.GroupByBlock()
//
//	ethPrice, err := t.RequireETHPrice(ctx, x)
//	if err != nil {
//		return nil, err
//	}
//}

func (t *UniApolloClient) GetTokenPriceFromSwaps(ctx context.Context, params GetTokenPriceParams) (TokenPriceResults, error) {
	rsp := make(TokenPriceResults, 0, len(params))

	return rsp, nil
}

func (t *UniApolloClient) GetTokenPricesByDate(ctx context.Context, date string, tokens []string) (map[string]float64, error) {
	// uniswap 限制最大1000
	batches := helper.SplitSlice(tokens, 500)

	priceUSDMaps, errs := wg.Run[[]string, map[string]float64](ctx, batches, func(ctx context.Context, p []string) ([]map[string]float64, error) {
		priceETHMap, err := t.GetTokenPriceUSDByDate(ctx, date, p)
		if err != nil {
			return nil, err
		}

		return []map[string]float64{priceETHMap}, err
	})

	if len(errs) > 0 {
		return nil, errs[0]
	}

	rsp := make(map[string]float64, len(tokens))

	for _, priceMap := range priceUSDMaps {
		for s, f := range priceMap {
			rsp[s] = f
		}
	}

	return rsp, nil
}

func (t *UniApolloClient) GetTokenPriceUSDByDate(ctx context.Context, date string, tokens []string) (map[string]float64, error) {

	v2Map, err := t.getV2TokenPriceUSDByDate(ctx, date, tokens)
	if err != nil {
		return nil, err
	}

	v3Map, err := t.getV3TokenPriceUSDByDate(ctx, date, tokens)
	if err != nil {
		return nil, err
	}

	rsp := make(map[string]float64, len(tokens))

	for v2, f2 := range v2Map {
		rsp[v2] = f2
	}

	for v3, f3 := range v3Map {
		if f3 == 0 {
			continue
		}

		if rsp[v3] > 0 {
			continue
		}
		rsp[v3] = f3
	}

	return rsp, nil
}

func (t *UniApolloClient) getV3TokenPriceUSDByDate(ctx context.Context, date string, tokens []string) (map[string]float64, error) {

	dateTime, _ := time.Parse(time.DateOnly, date)
	dateTs := (dateTime.Unix() / 86400) * 86400

	tokens = helper.ToSet(tokens)
	tokensLen := len(tokens)

	params := map[string]interface{}{
		"date":   dateTs,
		"tokens": tokens,
		"first":  tokensLen,
	}

	sql := `
		 query tokens($date: Int!, $tokens: [Bytes!], $first: Int!) {
			  tokens(where: {id_in: $tokens}, first: $first) {
    				tokenDayData(where: {date: $date}) {
      					priceUSD
      					date
    				}
			        id
			  }
		 }
	`

	rsp := make(map[string]float64, tokensLen)

	var v3dst struct {
		Data struct {
			Tokens RawTokens
		}
	}

	v3Url := fmt.Sprintf(t.EntrypointFormat, "v3")
	err := t.execute(ctx, v3Url, sql, params, &v3dst)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	for _, x := range v3dst.Data.Tokens {
		if len(x.TokenDayData) > 0 {
			rsp[x.Id] = conv.Float64(x.TokenDayData[0].PriceUSD)
		}
	}

	return rsp, nil

}

func (t *UniApolloClient) getV2TokenPriceUSDByDate(ctx context.Context, date string, tokens []string) (map[string]float64, error) {

	dateTime, _ := time.Parse(time.DateOnly, date)
	dateTs := (dateTime.Unix() / 86400) * 86400

	tokens = helper.ToSet(tokens)
	tokensLen := len(tokens)

	params := map[string]interface{}{
		"date":   dateTs,
		"tokens": tokens,
		"first":  tokensLen,
	}

	sql := `
		 query tokenDayDatas($date: Int!, $tokens: [Bytes!], $first: Int!) {
			    tokenDayDatas(where: {date: $date, id_in: $tokens}) {
					id
					date
					priceUSD
				}
		 }
	`

	rsp := make(map[string]float64, tokensLen)

	var v3dst struct {
		Data struct {
			Tokens TokenDayDatas
		}
	}

	v3Url := fmt.Sprintf(t.EntrypointFormat, "v3")
	err := t.execute(ctx, v3Url, sql, params, &v3dst)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	for _, x := range v3dst.Data.Tokens {
		rsp[x.Id] = x.PriceUSD
	}

	return rsp, nil

}

func (t *UniApolloClient) GetTokenPrices(ctx context.Context, block int64, tokens []string) (map[string]float64, error) {

	// uniswap 限制最大1000
	batches := helper.SplitSlice(tokens, 500)

	priceETHMaps, errs := wg.Run[[]string, map[string]float64](ctx, batches, func(ctx context.Context, p []string) ([]map[string]float64, error) {
		priceETHMap, err := t.getTokenPriceETH(ctx, block, p)
		if err != nil {
			return nil, err
		}

		return []map[string]float64{priceETHMap}, err
	})

	if len(errs) > 0 {
		return nil, errs[0]
	}

	ethPrice, err := t.RequireETHPrice(ctx, block)
	if err != nil {
		return nil, err
	}

	rsp := make(map[string]float64, len(tokens))

	for _, ethMap := range priceETHMaps {
		for token, priceETH := range ethMap {
			priceUSD := new(big.Float).Mul(big.NewFloat(priceETH), big.NewFloat(ethPrice))
			realPrice, _ := priceUSD.Float64()

			rsp[token] = realPrice
		}
	}

	return rsp, nil
}

func (t *UniApolloClient) GetTokenPrice(ctx context.Context, params GetTokenPriceParams) (TokenPriceResults, int, error) {

	rsp := make(TokenPriceResults, 0, len(params))

	swg := sync.WaitGroup{}
	lc := sync.Mutex{}

	var errs []error

	blockMap := params.GroupByBlock()

	for b, y := range blockMap {

		tokens := y.Tokens()

		ethPrice, err := t.RequireETHPrice(ctx, b)
		if err != nil {
			return nil, 0, err
		}

		rsp = append(rsp, TokenPriceResult{
			BlockNumber: b,
			Token:       "eth",
			PriceUSD:    ethPrice,
		})

		// uniswap 限制最大1000
		batches := helper.SplitSlice(tokens, 500)

		for _, batch := range batches {

			swg.Add(1)
			go func(ctx context.Context, block int64, tokens []string) {
				defer helper.DeferFunc(func() {
					swg.Done()
				})

				priceETHMap, err := t.getTokenPriceETH(ctx, block, tokens)
				if err != nil {
					lc.Lock()
					errs = append(errs, err)
					lc.Unlock()
					return
				}

				lc.Lock()
				for token, priceETH := range priceETHMap {

					//// uni tokens 中拿到的价值为0的token 可能不准，需要从swap中做一次double check
					//if priceETH == 0 {
					//
					//}

					priceUSD := new(big.Float).Mul(big.NewFloat(priceETH), big.NewFloat(ethPrice))

					realPrice, _ := priceUSD.Float64()

					rsp = append(rsp, TokenPriceResult{
						BlockNumber: block,
						Token:       token,
						PriceUSD:    realPrice,
					})
				}
				lc.Unlock()

			}(ctx, b, batch)

		}

	}

	swg.Wait()

	if len(errs) > 0 {
		return rsp, len(blockMap), errs[0]
	}

	return rsp, len(blockMap), nil
}

func (t *UniApolloClient) RequireETHPrice(ctx context.Context, number int64) (float64, error) {

	price, err := t.getEthPriceFromV2(ctx, number)
	if err != nil {
		return 0, err
	}

	if price == 0 {
		price, err = t.getEthPriceFromV2(ctx, number)
		if err != nil {
			return 0, err
		}
	}

	if price == 0 {
		return 0, fmt.Errorf("no eth price found: %d", number)
	}

	return price, nil

}

func (t *UniApolloClient) getEthPriceFromV3(ctx context.Context, number int64) (float64, error) {

	if number < 12369620 {
		return 0, nil
	}

	params := map[string]interface{}{
		"number": number,
	}

	v3Sql := `
 		query bundles($number: Int!) {
			   bundles(first: 1, block: {number: $number}) {
    				id
    				ethPriceUSD
  				}
		}
    `

	var dst struct {
		Data struct {
			Bundles []struct {
				Id          string
				EthPriceUSD string
			}
		}
	}
	v3Url := fmt.Sprintf(t.EntrypointFormat, "v3")

	err := t.execute(ctx, v3Url, v3Sql, params, &dst)
	if err != nil {
		return 0, xerror.Wrap(err)
	}

	if len(dst.Data.Bundles) == 0 {
		return 0, nil
	}

	return conv.Float64(dst.Data.Bundles[0].EthPriceUSD), nil
}

func (t *UniApolloClient) getEthPriceFromV2(ctx context.Context, number int64) (float64, error) {

	if number < 10000833 {
		return 201, nil
	}

	params := map[string]interface{}{
		"number": number,
	}

	v2Sql := `
 		query bundles($number: Int!) {
			   bundles(first: 1, block: {number: $number}) {
    				id
    				ethPrice
  				}
		}
    `

	var dst struct {
		Data struct {
			Bundles []struct {
				Id       string
				EthPrice string
			}
		}
	}
	v2Url := fmt.Sprintf(t.EntrypointFormat, "v2")

	err := t.execute(ctx, v2Url, v2Sql, params, &dst)
	if err != nil {
		return 0, xerror.Wrap(err)
	}

	if len(dst.Data.Bundles) == 0 {
		return 0, nil
	}

	return conv.Float64(dst.Data.Bundles[0].EthPrice), nil
}

func (t *UniApolloClient) getTokenPriceETH(ctx context.Context, block int64, tokens []string) (map[string]float64, error) {

	tokens = helper.ToSet(tokens)
	tokensLen := len(tokens)

	params := map[string]interface{}{
		"number": block,
		"tokens": tokens,
		"first":  tokensLen,
	}

	sql := `
		 query tokens($number: Int!, $tokens: [Bytes!], $first: Int!) {
			  tokens(block: {number: $number}, where: {id_in: $tokens}, first: $first) {
					derivedETH
					id
			  }
		 }
	`

	// 合并，优先取非0
	rsp := make(map[string]float64, tokensLen)

	v2Map := make(map[string]RawToken, tokensLen)
	v3Map := make(map[string]RawToken, tokensLen)

	if block >= 10000833 {
		// v2
		var v2dst struct {
			Data struct {
				Tokens RawTokens
			}
		}

		v2Url := fmt.Sprintf(t.EntrypointFormat, "v2")

		err := t.execute(ctx, v2Url, sql, params, &v2dst)
		if err != nil {
			return nil, xerror.Wrap(err)
		}

		v2Map = v2dst.Data.Tokens.ToMap()
	}

	if block >= 12369620 {
		// v3
		var v3dst struct {
			Data struct {
				Tokens RawTokens
			}
		}

		v3Url := fmt.Sprintf(t.EntrypointFormat, "v3")
		err := t.execute(ctx, v3Url, sql, params, &v3dst)
		if err != nil {
			return nil, xerror.Wrap(err)
		}

		v3Map = v3dst.Data.Tokens.ToMap()
	}

	for k, y := range v2Map {
		rsp[k] = conv.Float64(y.DerivedETH)
	}
	for k, y := range v3Map {

		if _, found := v2Map[k]; found {
			if y.DerivedETH != "0" {
				rsp[k] = conv.Float64(y.DerivedETH)
			}
		} else {
			rsp[k] = conv.Float64(y.DerivedETH)
		}

	}

	return rsp, nil

}

func (t *UniApolloClient) ListV2Pairs(ctx context.Context, timestampStart, page, size int64) (APairs, error) {
	if size > 1000 {
		return nil, fmt.Errorf("argument must be between 0 and 1000, but is %d", size)
	}
	URL := fmt.Sprintf(t.EntrypointFormat, "v2")

	params := map[string]interface{}{
		"createdAtTimestampStart": timestampStart,
		"first":                   size,
		"skip":                    page * size,
	}

	sql := `
		 query pairs($createdAtTimestampStart: Int, $first: Int!, $skip: Int!) {
				pairs(first: $first, skip: $skip, where: {createdAtTimestamp_gt: $createdAtTimestampStart}) {
					  id
					  token0 {
						id
						symbol
						name
						decimals
					  }
					  token1 {
						id
						symbol
						name
						decimals
					  }
					createdAtTimestamp
				}
			  }
	`

	var dst struct {
		Data struct {
			Pairs APairs
		}
	}

	err := t.execute(ctx, URL, sql, params, &dst)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	for _, pair := range dst.Data.Pairs {
		pair.Version = "v2"
	}

	return dst.Data.Pairs, nil
}

func (t *UniApolloClient) ListV3Pairs(ctx context.Context, timestampStart, page, size int64) (APairs, error) {
	if size > 1000 {
		return nil, fmt.Errorf("argument must be between 0 and 1000, but is %d", size)
	}
	URL := fmt.Sprintf(t.EntrypointFormat, "v3")

	params := map[string]interface{}{
		"createdAtTimestampStart": timestampStart,
		"first":                   size,
		"skip":                    page * size,
	}

	sql := `
		 query pairs($createdAtTimestampStart: Int, $first: Int!, $skip: Int!) {
				pools(first: $first, skip: $skip, where: {createdAtTimestamp_gt: $createdAtTimestampStart}) {
					  id
					  token0 {
						id
						symbol
						name
						decimals
					  }
					  token1 {
						id
						symbol
						name
						decimals
					  }
					createdAtTimestamp
				}
			  }
	`

	var dst struct {
		Data struct {
			Pools APairs
		}
	}

	err := t.execute(ctx, URL, sql, params, &dst)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	for _, pair := range dst.Data.Pools {
		pair.Version = "v2"
	}

	return dst.Data.Pools, nil
}

// 0xEf1c6E67703c7BD7107eed8303Fbe6EC2554BF6B Uniswap: Universal Router
// 0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D Uniswap V2:  Router 2
// 0xe66B31678d6C16E9ebf358268a790B763C133750 0x: Coinbase Wallet Proxy
// 0x74de5d4FCbf63E00296fd95d33236B9794016631 未知
// 0xDef1C0ded9bec7F1a1670819833240f027b25EfF 0x: Exchange Proxy

// 都是稳定币的币对过滤掉
// 不包含稳定币的币对过滤掉

// 猜测 token 转出 并且 sender 为  0xEf1c6E67703c7BD7107eed8303Fbe6EC2554BF6B  to 如果不同 就是from
// 猜测 to 地址为 0xe66B31678d6C16E9ebf358268a790B763C133750，其转出的地址才是 to
// out  + sender = 0x: Exchange Proxy  => to 一定为 0x: Coinbase Wallet Proxy
// out  + sender = Uniswap V2:  Router 2  => to 就是 from

// 算出的最后结果 不管任何 都看看有没有上层转入 和 下层 转出

func (t *UniApolloClient) ListV2Swaps(ctx context.Context, pairId string, timestampStart, page, size int64, ethClient *ethclient.Client) (ASwaps, error) {
	if size > 1000 {
		return nil, fmt.Errorf("argument must be between 0 and 1000, but is %d", size)
	}

	URL := fmt.Sprintf(t.EntrypointFormat, "v2")

	params := map[string]interface{}{
		"timestampGt": timestampStart,
		"pairId":      pairId,
		"first":       size,
		"skip":        page * size,
	}
	sql := `
		query swaps($timestampGt: Int, $pairId: Bytes!, $first: Int!, $skip: Int!) {
				swaps(first: $first, skip: $skip, where:{ pair: $pairId, timestamp_gt: $timestampGt }) {
					 id
					 transaction {
						blockNumber
					 }
					 pair {
					  id
                      token0 { id, derivedETH, name }
                      token1 { id, derivedETH, name }
					 }
					 timestamp
					 sender
					 amount0In
					 amount0Out
					 amount1In
					 amount1Out
					 amountUSD
				 }
			}
	`

	var dst struct {
		Data struct {
			Swaps []struct {
				Id          string
				Transaction struct {
					BlockNumber string
				}
				Pair struct {
					Id     string
					Token0 struct {
						Id         string
						DerivedETH string
						Name       string
					}
					Token1 struct {
						Id         string
						DerivedETH string
						Name       string
					}
				}
				Timestamp  string
				Sender     string
				Amount0In  string
				Amount0Out string
				Amount1In  string
				Amount1Out string
				AmountUSD  string
			}
		}
	}

	err := t.execute(ctx, URL, sql, params, &dst)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	var rsp ASwaps
	for _, x := range dst.Data.Swaps {

		id := strings.Split(x.Id, "-")[0]

		// todo
		// v2 没有 from 需要根据 transaction hash 获取
		tx, _, err := ethClient.TransactionByHash(ctx, common.HexToHash(id))
		if err != nil {
			slf.WithError(err).Errorf("from empty, %v", x.Id)
			continue
		}

		msg, _ := tx.AsMessage(types.NewLondonSigner(tx.ChainId()), nil)

		from := msg.From()

		y := ASwap{
			Id:      x.Id,
			Version: "v2",
			Pair: &APair{
				Id: x.Pair.Id,
				Token0: &AToken{
					Id:   x.Pair.Token0.Id,
					Name: x.Pair.Token0.Name,
				},
				Token1: &AToken{
					Id:   x.Pair.Token1.Id,
					Name: x.Pair.Token1.Name,
				},
			},
			Timestamp:   conv.Int64(x.Timestamp),
			From:        from.Hex(),
			BlockNumber: x.Transaction.BlockNumber,
		}

		amount0In, _ := new(big.Float).SetString(x.Amount0In)
		amount0Out, _ := new(big.Float).SetString(x.Amount0Out)

		amount1In, _ := new(big.Float).SetString(x.Amount1In)
		amount1Out, _ := new(big.Float).SetString(x.Amount1Out)

		y.AmountUSD = x.AmountUSD
		y.Amount0 = amount0Out.Sub(amount0Out, amount0In).String()
		y.Amount1 = amount1Out.Sub(amount1Out, amount1In).String()

		rsp = append(rsp, &y)
	}

	return rsp, nil
}

func (t *UniApolloClient) ListV2PairsByBlockNumber(ctx context.Context, blockNumberFrom, blockNumberTo uint64) (APairs, error) {
	URL := fmt.Sprintf(t.EntrypointFormat, "v2")

	params := map[string]interface{}{
		"blockNumberGte": blockNumberFrom,
		"blockNumberLte": blockNumberTo,
		"first":          1000,
	}

	sql := `
		 query pairs($blockNumberGte: Int!, $blockNumberLte: Int!, $first: Int) {
				pairs(first: $first, where: {createdAtBlockNumber_gte: $blockNumberGte, createdAtBlockNumber_lte: $blockNumberLte}) {
					id
					token0 {
						id
						symbol
						name
						decimals
					}
					token1 {
						id
						symbol
						name
						decimals
					}
					createdAtBlockNumber
					createdAtTimestamp
				}
			  }
	`

	var dst struct {
		Data struct {
			Pairs APairs
		}
	}

	err := t.execute(ctx, URL, sql, params, &dst)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	for _, pair := range dst.Data.Pairs {
		pair.Swap()
		pair.Version = "uniswap_v2"
	}

	return dst.Data.Pairs, nil
}

func (t *UniApolloClient) ListV3PairsByBlockNumber(ctx context.Context, blockNumberFrom, blockNumberTo uint64) (APairs, error) {

	URL := fmt.Sprintf(t.EntrypointFormat, "v3")

	params := map[string]interface{}{
		"blockNumberGte": blockNumberFrom,
		"blockNumberLte": blockNumberTo,
		"first":          1000,
	}

	sql := `
		 query pairs($blockNumberGte: Int!, $blockNumberLte: Int!, $first: Int) {
				pools(first: $first, where: {createdAtBlockNumber_gte: $blockNumberGte, createdAtBlockNumber_lte: $blockNumberLte}) {
					id
					token0 {
						id
						name
						symbol
						decimals
					}
					token1 {
						id
						name
						symbol
						decimals
					}
					createdAtTimestamp
					createdAtBlockNumber
				}
			  }
	`

	var dst struct {
		Data struct {
			Pools APairs
		}
	}

	err := t.execute(ctx, URL, sql, params, &dst)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	for _, pair := range dst.Data.Pools {
		pair.Swap()
		pair.Version = "uniswap_v3"
	}

	return dst.Data.Pools, nil
}

func (t *UniApolloClient) ListV2SwapsByBlockNumber(ctx context.Context, blockNumberGte, blockNumberLte uint64) (ASwaps, error) {

	URL := fmt.Sprintf(t.EntrypointFormat, "v2")

	params := map[string]interface{}{
		"blockNumber_gte": blockNumberGte,
		"blockNumber_lte": blockNumberLte,
	}

	sql := `
		query swaps($blockNumber_lte: Int!, $blockNumber_gte: Int!) {
				swaps(where: {transaction_: {blockNumber_gte: $blockNumber_gte, blockNumber_lte: $blockNumber_lte}}) {
					 id
					 transaction {
						id
						blockNumber
					 }
					 pair {
					  id
                      token0 { id, derivedETH, name }
                      token1 { id, derivedETH, name }
					 }
					 timestamp
					 sender
					 amount0In
					 amount0Out
					 amount1In
					 amount1Out
					 amountUSD
					 logIndex
				 }
			}
	`

	var dst struct {
		Data struct {
			Swaps RawV2Swaps
		}
	}

	err := t.execute(ctx, URL, sql, params, &dst)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	swaps := dst.Data.Swaps.AsASwaps()

	return swaps, nil
}

// 时间跨度不要太大 最好1m以内
// 同一个 transaction 中每个路由会都会有一条swap记录
// 同一个 transaction中 可能出现 同时存在 v2 v3的 swap log  eg. 0x0afae86907511b9085232f7f2330749b0378e92431bc7eebe75c5ea8a21af736
func (t *UniApolloClient) ListV3SwapsByBlockNumber(ctx context.Context, blockNumberGte, blockNumberLte uint64) (ASwaps, error) {
	URL := fmt.Sprintf(t.EntrypointFormat, "v3")
	params := map[string]interface{}{
		"blockNumber_gte": blockNumberGte,
		"blockNumber_lte": blockNumberLte,
	}

	sql := `
		query swaps($blockNumber_lte: Int!, $blockNumber_gte: Int!) {
				swaps(where: {transaction_: {blockNumber_gte: $blockNumber_gte, blockNumber_lte: $blockNumber_lte}}) {
					id
					transaction {
						blockNumber
						id
					}
 					pool {
					   id
					   token0 { id, derivedETH, name }
					   token1 { id, derivedETH, name }
					}
					origin
					timestamp
					sender
					amount0
					amount1
					amountUSD
					logIndex
				 }
			}

	`

	var dst struct {
		Data struct {
			Swaps RawV3Swaps
		}
	}

	err := t.execute(ctx, URL, sql, params, &dst)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return dst.Data.Swaps.AsASwaps(), nil
}

func (t *UniApolloClient) ListV3SwapsByPairId(ctx context.Context, pairId string, timestampStart, page, size int64) (ASwaps, error) {
	if size > 1000 {
		return nil, fmt.Errorf("argument must be between 0 and 1000, but is %d", size)
	}
	URL := fmt.Sprintf(t.EntrypointFormat, "v3")

	params := map[string]interface{}{
		"timestampGt": timestampStart,
		"pairId":      pairId,
		"first":       size,
		"skip":        page * size,
	}

	sql := `
		query swaps($timestampGt: Int, $pairId: Bytes!, $first: Int!, $skip: Int!) {
				swaps(first: $first, skip: $skip, where:{ pool: $pairId, timestamp_gt: $timestampGt }) {
					id
					transaction {
						blockNumber
					}
 					pool {
					   id
					   token0 { id, derivedETH, name }
					   token1 { id, derivedETH, name }
					}
					origin
					timestamp
					sender
					amount0
					amount1
					amountUSD	
				 }
			}

	`

	var dst struct {
		Data struct {
			Swaps RawV3Swaps
		}
	}

	err := t.execute(ctx, URL, sql, params, &dst)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	var rsp ASwaps
	for _, x := range dst.Data.Swaps {

		y := ASwap{
			Id:      x.Id,
			Version: "v3",
			Pair: &APair{
				Id: x.Pool.Id,
				Token0: &AToken{
					Id:   x.Pool.Token0.Id,
					Name: x.Pool.Token0.Name,
				},
				Token1: &AToken{
					Id:   x.Pool.Token1.Id,
					Name: x.Pool.Token1.Name,
				},
			},
			Timestamp:   conv.Int64(x.Timestamp),
			From:        x.Origin,
			BlockNumber: x.Transaction.BlockNumber,
		}

		// v3 的支出是正数，收入是负数。反过来更合理一些
		amount0, _ := new(big.Float).SetString(x.Amount0)
		amount0 = amount0.Sub(big.NewFloat(0), amount0)

		amount1, _ := new(big.Float).SetString(x.Amount1)
		amount1 = amount1.Sub(big.NewFloat(0), amount1)

		y.AmountUSD = x.AmountUSD
		y.Amount0 = amount0.String()
		y.Amount1 = amount1.String()

		rsp = append(rsp, &y)
	}

	return rsp, nil
}

func (t *UniApolloClient) FindTokenMetadatas(ctx context.Context, addresses []string) (ATokens, error) {
	if len(addresses) <= 1000 {
		return t.findTokenMetadatas(ctx, addresses)
	}

	var rsp ATokens
	for _, batch := range helper.SplitSlice(addresses, 900) {
		tokens, err := t.findTokenMetadatas(ctx, batch)
		if err != nil {
			return nil, err
		}

		rsp = append(rsp, tokens...)
	}

	return rsp, nil
}

func (t *UniApolloClient) findTokenMetadatas(ctx context.Context, addresses []string) (ATokens, error) {
	addresses = helper.ToSet(addresses)

	sql := `
		query tokens($ids: [Bytes!], $first: Int) {
			tokens(where: {id_in: $ids}, first: $first){
                id
				name
				symbol
				decimals
			}
		}
	`

	var dst struct {
		Data struct {
			Tokens []*AToken
		} `json:"data"`
	}

	params := map[string]interface{}{
		"ids":   addresses,
		"first": len(addresses),
	}

	v2Host := fmt.Sprintf(t.EntrypointFormat, "v2")

	err := t.execute(ctx, v2Host, sql, params, &dst)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	if len(dst.Data.Tokens) == 0 {

		v3Host := fmt.Sprintf(t.EntrypointFormat, "v3")

		err = t.execute(ctx, v3Host, sql, params, &dst)
		if err != nil {
			return nil, xerror.Wrap(err)
		}
	}

	return dst.Data.Tokens, nil
}

func (t *UniApolloClient) RequireTokenMetadata(ctx context.Context, address string) (*AToken, error) {
	metadatas, err := t.FindTokenMetadatas(ctx, []string{address})
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	if len(metadatas) == 0 {
		return nil, xerror.Wrapf("no metadata found by address: %s", address)
	}

	return metadatas[0], nil
}

func (t *UniApolloClient) execute(ctx context.Context, url, sql string, params map[string]interface{}, dst any) error {

	reqBody := map[string]interface{}{
		//"operationName": "pairs",
		"variables": params,
		"query":     sql,
	}

	bodyBytes := bytes.NewBuffer([]byte(conv.M2J(reqBody)))
	req, err := http.NewRequest("POST", url, bodyBytes)
	if err != nil {
		return xerror.Wrap(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return xerror.Wrap(err)
	}

	bytesBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return xerror.Wrap(err)
	}

	if resp.StatusCode != http.StatusOK {
		return xerror.Wrapf(string(bytesBody))
	}

	var errResp errorResponse
	if err := json.Unmarshal(bytesBody, &errResp); err != nil {
		return err
	}

	if len(errResp.Errors) > 0 {
		return errors.New(errResp.Errors[0].Message)
	}

	//defer resp.Body.Close()

	if err := json.Unmarshal(bytesBody, dst); err != nil {
		return xerror.Wrap(err)
	}

	return nil
}

type errorResponse struct {
	Errors []struct {
		Locations []struct {
			Line   int `json:"line"`
			Column int `json:"column"`
		} `json:"locations"`
		Message string `json:"message"`
	} `json:"errors"`
}
