package swap

import (
	"gitee.com/meepo/backend/kit/components/chain/eth"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"math/big"
	"strings"
)

func (rs RawV3Swaps) AsASwaps() (rsp ASwaps) {
	// todo A => B => C 转换成 A => C
	// 目前 只保留 包含非稳定币的Swap 过滤掉 两个稳定币的Swap
	for _, r := range rs {

		if helper.InSliceIgnoreCase(r.Pool.Token0.Id, eth.STABLE_COINS) &&
			helper.InSliceIgnoreCase(r.Pool.Token1.Id, eth.STABLE_COINS) {

			continue
		}

		rsp = append(rsp, r.AsASwap())
	}

	return
}

type RawV3Swaps []RawV3Swap

type ByLogIndex RawV3Swaps

func (a ByLogIndex) Len() int           { return len(a) }
func (a ByLogIndex) Less(i, j int) bool { return a[i].LogIndex < a[j].LogIndex }
func (a ByLogIndex) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (r *RawV3Swap) AsASwap() *ASwap {

	id := strings.Split(r.Id, "#")[0] + "#" + r.LogIndex

	y := ASwap{
		Id:      id,
		Version: "uniswap_v3",
		Pair: &APair{
			Id: r.Pool.Id,
			Token0: &AToken{
				Id:   r.Pool.Token0.Id,
				Name: r.Pool.Token0.Name,
			},
			Token1: &AToken{
				Id:   r.Pool.Token1.Id,
				Name: r.Pool.Token1.Name,
			},
		},
		Timestamp:   conv.Int64(r.Timestamp),
		From:        r.Origin,
		BlockNumber: r.Transaction.BlockNumber,
		LogIndex:    conv.Int(r.LogIndex),
		Hash:        r.Transaction.Id,
	}

	// v3 的支出是正数，收入是负数。反过来更合理一些
	amount0, _ := new(big.Float).SetString(r.Amount0)
	amount0 = amount0.Sub(big.NewFloat(0), amount0)

	amount1, _ := new(big.Float).SetString(r.Amount1)
	amount1 = amount1.Sub(big.NewFloat(0), amount1)

	y.AmountUSD = r.AmountUSD
	y.Amount0 = amount0.String()
	y.Amount1 = amount1.String()

	y.Pair.Swap()
	if y.Pair.Swapped {
		tmp := y.Amount0
		y.Amount0 = y.Amount1
		y.Amount1 = tmp
	}

	return &y
}

type RawToken struct {
	Id           string
	DerivedETH   string
	Name         string
	Symbol       string
	Decimals     string
	TokenDayData []DatePrice
}

type DatePrice struct {
	PriceUSD string
	Date     int64
}

type TokenDayData struct {
	Id       string
	Date     string
	PriceUSD float64
}

type TokenDayDatas []TokenDayData

type RawTokens []RawToken

func (ts RawTokens) ToMap() map[string]RawToken {
	rsp := make(map[string]RawToken, len(ts))

	for _, t := range ts {
		rsp[t.Id] = t
	}
	return rsp
}

func (ts RawTokens) AsAToken() ATokens {

	var rsp ATokens
	for _, t := range ts {
		rsp = append(rsp, &AToken{
			Id:       t.Id,
			Name:     t.Name,
			Symbol:   t.Symbol,
			Decimals: t.Decimals,
			PriceETH: conv.Float64(t.DerivedETH),
		})
	}

	return rsp
}

type RawV3Swap struct {
	Id          string
	Transaction struct {
		Id          string
		BlockNumber string
	}
	Pool struct {
		Id     string
		Token0 RawToken
		Token1 RawToken
	}
	Origin    string
	Timestamp string
	Sender    string
	Amount0   string
	Amount1   string
	AmountUSD string
	LogIndex  string
}

type RawV2Swap struct {
	Id          string
	Transaction struct {
		Id          string
		BlockNumber string
	}
	Pair struct {
		Id     string
		Token0 RawToken
		Token1 RawToken
	}
	Timestamp  string
	Sender     string
	Amount0In  string
	Amount0Out string
	Amount1In  string
	Amount1Out string
	AmountUSD  string
	LogIndex   string
}

type RawV2Swaps []RawV2Swap

func (rs RawV2Swaps) AsASwaps() (rsp []*ASwap) {
	for _, r := range rs {
		if helper.InSliceIgnoreCase(r.Pair.Token0.Id, eth.STABLE_COINS) &&
			helper.InSliceIgnoreCase(r.Pair.Token1.Id, eth.STABLE_COINS) {
			continue
		}
		rsp = append(rsp, r.AsASwap())
	}
	return
}

func (r *RawV2Swap) AsASwap() *ASwap {
	id := strings.Split(r.Id, "-")[0] + "#" + r.LogIndex

	from := ""

	y := ASwap{
		Id:      id,
		Version: "uniswap_v2",
		Pair: &APair{
			Id: r.Pair.Id,
			Token0: &AToken{
				Id:   r.Pair.Token0.Id,
				Name: r.Pair.Token0.Name,
			},
			Token1: &AToken{
				Id:   r.Pair.Token1.Id,
				Name: r.Pair.Token1.Name,
			},
		},
		Timestamp:   conv.Int64(r.Timestamp),
		From:        from,
		BlockNumber: r.Transaction.BlockNumber,
		LogIndex:    conv.Int(r.LogIndex),
		Hash:        r.Transaction.Id,
	}

	amount0In, _ := new(big.Float).SetString(r.Amount0In)
	amount0Out, _ := new(big.Float).SetString(r.Amount0Out)

	amount1In, _ := new(big.Float).SetString(r.Amount1In)
	amount1Out, _ := new(big.Float).SetString(r.Amount1Out)

	y.AmountUSD = r.AmountUSD
	y.Amount0 = new(big.Float).Sub(amount0Out, amount0In).String()
	y.Amount1 = new(big.Float).Sub(amount1Out, amount1In).String()

	y.Pair.Swap()
	if y.Pair.Swapped {
		tmp := y.Amount0
		y.Amount0 = y.Amount1
		y.Amount1 = tmp
	}

	return &y
}
