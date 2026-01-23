package eth

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"gitee.com/meepo/backend/kit/components/chain/eth/contracts"
	ethutils "gitee.com/meepo/backend/kit/components/chain/eth/utils"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/redis/go-redis/v9"
	"go.uber.org/ratelimit"
	"math"
	"math/big"
	"strings"
	"sync"
	"time"
)

func NewEthClient(clients Clients) (*Client, error) {
	return &Client{cs: clients}, nil
}

type Client struct {
	cs Clients
}

func (e *Client) GetEthClient() *ethclient.Client {
	return e.cs.Get()
}

func (e *Client) GetBlockTsCache(ctx context.Context, blocks []string, redisClient *redis.Client) (map[string]int64, error) {

	blocks = helper.ToSet(blocks)
	addressesLen := len(blocks)

	rsp := make(map[string]int64, addressesLen)
	if len(blocks) == 0 {
		return rsp, nil
	}

	redisKey := "block_metadata.ts"
	cacheResult, err := redisClient.HMGet(ctx, redisKey, blocks...).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}

	cacheLen := len(cacheResult)

	shouldFoundAddresses := make([]string, 0, addressesLen)

	for i := 0; i < addressesLen; i++ {
		if cacheLen-1 < i {
			shouldFoundAddresses = append(shouldFoundAddresses, blocks[i])
			continue
		}

		if cacheResult[i] == nil {
			shouldFoundAddresses = append(shouldFoundAddresses, blocks[i])
			continue
		}

		rsp[blocks[i]] = conv.Int64(cacheResult[i])
	}

	if len(shouldFoundAddresses) == 0 {
		return rsp, nil
	}

	rawMap, err := e.GetBlocksTs(ctx, shouldFoundAddresses)
	if err != nil {
		return nil, err
	}

	_, err = redisClient.HSet(ctx, redisKey, rawMap).Result()
	if err != nil {
		slf.WithError(err).Errorw("HSet err")
	}
	redisClient.Expire(ctx, redisKey, 365*24*time.Hour)

	for k, r := range rawMap {
		rsp[k] = conv.Int64(r)
	}

	return rsp, nil
}

func (e *Client) GetBlocksTs(ctx context.Context, blocks []string) (map[string]interface{}, error) {
	rsp := make(map[string]interface{}, len(blocks))
	var errs []error

	wg := sync.WaitGroup{}
	lc := sync.Mutex{}

	for _, x := range blocks {

		wg.Add(1)
		x := x
		go func(ctx2 context.Context, number string) {
			defer helper.DeferFunc(func() {
				wg.Done()
			})

			block, err := e.cs.Get().BlockByNumber(ctx2, big.NewInt(conv.Int64(number)))

			if err != nil {
				lc.Lock()
				errs = append(errs, err)
				lc.Unlock()
				return
			}

			lc.Lock()
			rsp[number] = int64(block.Time())
			lc.Unlock()

		}(ctx, x)
	}

	wg.Wait()

	if len(errs) != 0 {
		return nil, errs[0]
	}

	return rsp, nil
}

func (e *Client) GetTransferLogsByToken(ctx context.Context, token string, fromNumber, toNumber int64) (TransferEventLogs, error) {

	var from *big.Int
	if fromNumber > 0 {
		from = big.NewInt(fromNumber)
	}

	var to *big.Int
	if toNumber > 0 {
		from = big.NewInt(toNumber)
	}

	logs, err := e.cs.Get().FilterLogs(ctx, ethereum.FilterQuery{
		FromBlock: from,
		ToBlock:   to,
		Addresses: []common.Address{common.HexToAddress(token)},
		Topics: [][]common.Hash{
			{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")},
		},
	})

	if err != nil {
		return nil, err
	}

	var finalLogs TransferEventLogs
	for _, log := range logs {
		if len(log.Topics) != 3 {
			continue
		}
		finalLogs = append(finalLogs, TransferEventLog(log))
	}

	return finalLogs, nil

}

func (e *Client) GetTransferLogsByNumber(ctx context.Context, fromNumber, toNumber int64, tokens []string) (TransferEventLogs, error) {
	logs, err := e.GetLogsByNumberByTopicSync(ctx, fromNumber, toNumber, "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef", tokens)
	if err != nil {
		return nil, err
	}

	var finalLogs TransferEventLogs
	for _, log := range logs {
		if len(log.Topics) != 3 {
			continue
		}
		finalLogs = append(finalLogs, log)
	}

	return finalLogs, nil

}

func (e *Client) GetLogsByNumberByTopicSync(ctx context.Context, fromNumber, toNumber int64, topic string, tokens []string) (EventLogs, error) {

	var rsp EventLogs
	var errs []error

	wg := sync.WaitGroup{}
	lc := sync.Mutex{}

	limiter := ratelimit.New(500)

	var tokenAddresses []common.Address
	for _, token := range tokens {
		tokenAddresses = append(tokenAddresses, common.HexToAddress(token))
	}

	for i := fromNumber; i <= toNumber; i++ {

		limiter.Take()

		wg.Add(1)
		go func(ctx context.Context, number int64) {
			defer helper.DeferFunc(func() {
				wg.Done()
			})

			logs, err := e.cs.Get().FilterLogs(ctx, ethereum.FilterQuery{
				FromBlock: big.NewInt(number),
				ToBlock:   big.NewInt(number),
				Addresses: tokenAddresses,
				Topics: [][]common.Hash{
					{common.HexToHash(topic)},
				},
			})

			if err != nil {
				lc.Lock()
				errs = append(errs, err)
				lc.Unlock()
				return
			}

			if len(logs) > 0 {
				lc.Lock()
				for _, log := range logs {
					rsp = append(rsp, EventLog(log))
				}

				lc.Unlock()
			}

		}(ctx, i)
	}

	wg.Wait()

	if len(errs) != 0 {
		return nil, errs[0]
	}

	return rsp, nil

}

func (e *Client) GetTokensDecimalsCache(ctx context.Context, addresses []string, redisClient *redis.Client) (map[string]int, int, error) {

	addresses = helper.ToSet(addresses)
	addressesLen := len(addresses)

	rsp := make(map[string]int, addressesLen)
	if len(addresses) == 0 {
		return rsp, 0, nil
	}

	redisKey := "token_address.decimals"
	cacheResult, err := redisClient.HMGet(ctx, redisKey, addresses...).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, 0, err
	}

	cacheLen := len(cacheResult)

	shouldFoundAddresses := make([]string, 0, addressesLen)

	for i := 0; i < addressesLen; i++ {
		if cacheLen-1 < i {
			shouldFoundAddresses = append(shouldFoundAddresses, addresses[i])
			continue
		}

		if cacheResult[i] == nil {
			shouldFoundAddresses = append(shouldFoundAddresses, addresses[i])
			continue
		}

		rsp[addresses[i]] = conv.Int(cacheResult[i])
	}

	if len(shouldFoundAddresses) == 0 {
		return rsp, 0, nil
	}

	rawMap, err := e.GetTokensDecimals(ctx, shouldFoundAddresses)
	if err != nil {
		return nil, 0, err
	}

	_, err = redisClient.HSet(ctx, redisKey, rawMap).Result()
	if err != nil {
		slf.WithError(err).Errorw("HSet err")
	}
	redisClient.Expire(ctx, redisKey, 365*24*time.Hour)

	for k, r := range rawMap {
		rsp[k] = conv.Int(r)
	}

	return rsp, len(shouldFoundAddresses), nil
}

// >0, 0, -1(非erc20)
func (e *Client) GetTokensDecimals(ctx context.Context, addresses []string) (map[string]interface{}, error) {

	addresses = helper.ToSet(addresses)

	rsp := make(map[string]interface{}, len(addresses))

	var errs []error
	wg := sync.WaitGroup{}
	lc := sync.Mutex{}

	limiter := ratelimit.New(500)

	for _, x := range addresses {

		limiter.Take()

		wg.Add(1)
		go func(address string) {
			defer helper.DeferFunc(func() {
				wg.Done()
			})

			val, err := e.GetTokenDecimals(ctx, address)
			if err != nil {
				lc.Lock()
				errs = append(errs, err)
				lc.Unlock()
				return
			}

			lc.Lock()
			rsp[address] = val
			lc.Unlock()

		}(x)
	}

	wg.Wait()

	if len(errs) > 0 {
		return nil, errs[0]
	}

	return rsp, nil
}

func (e *Client) IsNormalByCacheInBatches(ctx context.Context, addresses []string, redisClient *redis.Client) (map[string]bool, error) {

	batches := helper.SplitSlice(addresses, 100000)

	var rsp map[string]bool

	for _, batch := range batches {

		r, _, err := e.IsNormalByCache(ctx, batch, redisClient)
		if err != nil {
			return nil, err
		}

		rsp = helper.AppendMap(rsp, r)
	}

	return rsp, nil
}

func (e *Client) IsNormalByCache(ctx context.Context, addresses []string, redisClient *redis.Client) (map[string]bool, int, error) {

	addresses = helper.ToSet(addresses)
	addressesLen := len(addresses)

	rsp := make(map[string]bool, addressesLen)
	if len(addresses) == 0 {
		return rsp, 0, nil
	}

	redisKey := "address.isNormalV3"
	cacheResult, err := redisClient.HMGet(ctx, redisKey, addresses...).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, 0, err
	}

	cacheLen := len(cacheResult)

	shouldFoundAddresses := make([]string, 0, addressesLen)

	for i := 0; i < addressesLen; i++ {
		if cacheLen-1 < i {
			shouldFoundAddresses = append(shouldFoundAddresses, addresses[i])
			continue
		}

		if cacheResult[i] == nil {
			shouldFoundAddresses = append(shouldFoundAddresses, addresses[i])
			continue
		}

		rsp[addresses[i]] = cacheResult[i] == "1"
	}

	if len(shouldFoundAddresses) == 0 {
		return rsp, 0, nil
	}

	rawMap, err := e.IsNormal(ctx, shouldFoundAddresses)
	if err != nil {
		return nil, 0, err
	}

	_, err = redisClient.HSet(ctx, redisKey, rawMap).Result()
	if err != nil {
		slf.WithError(err).Errorw("HSet err")
	}

	redisClient.Expire(ctx, redisKey, 365*24*time.Hour)

	for k, r := range rawMap {
		rsp[k] = r == "1"
	}

	return rsp, len(shouldFoundAddresses), nil

}

func (e *Client) IsNormal(ctx context.Context, addresses []string) (map[string]interface{}, error) {

	addresses = helper.ToSet(addresses)

	rsp := make(map[string]interface{}, len(addresses))

	var errs []error
	wg := sync.WaitGroup{}
	lc := sync.Mutex{}

	limiter := ratelimit.New(300)

	for _, x := range addresses {

		if helper.InSliceIgnoreCase(x, IGNORE_ADDRESSES) {
			rsp[x] = false
			continue
		}

		limiter.Take()

		wg.Add(1)
		go func(address string) {
			defer helper.DeferFunc(func() {
				wg.Done()
			})

			val, err := e.cs.Get().CodeAt(ctx, common.HexToAddress(address), nil)
			if err != nil {
				lc.Lock()
				errs = append(errs, err)
				lc.Unlock()
				return
			}

			lc.Lock()
			rsp[address] = helper.Choose(len(val) == 0, "1", "0")
			lc.Unlock()

		}(x)
	}

	wg.Wait()

	if len(errs) > 0 {
		return nil, errs[0]
	}

	return rsp, nil

}

type GetTokenBalancesParam struct {
	TokenAddress string
	Account      string
	Decimals     int
	BlockNumber  int64
}

func (p *GetTokenBalancesParam) Key() string {
	return fmt.Sprintf("%d-%s-%s", p.BlockNumber, p.Account, p.TokenAddress)
}

type GetTokenBalancesParams []GetTokenBalancesParam

type TokenBalancesResult struct {
	TokenAddress string
	Account      string
	Amount       *big.Float
	BlockNumber  int64
}

type TokenBalancesResults []*TokenBalancesResult

type TokenBalancesResultMap map[string]big.Float

func (m TokenBalancesResultMap) Find(block int64, account, tokenAddress string) (big.Float, bool) {
	val, found := m[fmt.Sprintf("%d-%s-%s", block, account, tokenAddress)]
	return val, found
}

func (rs TokenBalancesResults) AsMap() TokenBalancesResultMap {

	rsp := make(TokenBalancesResultMap, len(rs))

	for _, r := range rs {
		rsp[fmt.Sprintf("%d-%s-%s", r.BlockNumber, r.Account, r.TokenAddress)] = *r.Amount
	}
	return rsp
}

type Options struct {
	IgnoreErrRecord bool
	QPSLimit        int
	RpcIndex        int
}

//func (e *Client) GetTokenBalancesLoopCache(ctx context.Context, params GetTokenBalancesParams, options ...Options) (TokenBalancesResults, error) {
//
//	cacheKeySet := mapset.NewSet[string]()
//	for _, x := range params {
//		cacheKeySet.Add(x.Key())
//	}
//
//	cacheKeys := cacheKeySet.ToSlice()
//
//	redisKey := "tokenBalance.cache"
//	cacheBalances, err := comp.SDK().Redis().HMGet(ctx, redisKey, cacheKeys...).Result()
//	if err != nil {
//		return nil, err
//	}
//
//	cacheBalanceMap := make(map[string]float64, len(cacheKeys))
//	for i, balance := range cacheBalances {
//		if balance != nil {
//			cacheBalanceMap[cacheKeys[i]] = conv.Float64(balance)
//		}
//	}
//
//	toCallParams := make(GetTokenBalancesParams, 0, len(cacheKeys))
//	for _, param := range params {
//		if val, found := cacheBalanceMap[param.Key()] ; found {
//
//
//		}
//
//	}
//
//}
//
//
//
//
//
//}

// 一直报错会导致一直循环 todo
func (e *Client) GetTokenBalancesLoop(ctx context.Context, params GetTokenBalancesParams, qpsLimit int) (TokenBalancesResults, error) {

	if len(params) == 0 {
		return nil, nil
	}

	rsp := make(TokenBalancesResults, 0, len(params))

	paramsMap := make(map[string]GetTokenBalancesParam, len(params))
	for _, x := range params {
		key := fmt.Sprintf("%s-%s-%d", x.Account, x.TokenAddress, x.BlockNumber)
		paramsMap[key] = x
	}
	totalParamsLen := len(paramsMap)

	i := 0

	for {

		tmpParams := helper.MapValues(paramsMap)
		if len(tmpParams) > qpsLimit {
			tmpParams = tmpParams[0:qpsLimit]
		}

		slf.Debugw("GetTokenBalancesLoop",
			slf.Int("params", len(params)),
			slf.Int("total params", totalParamsLen),
			slf.Int("current params", len(tmpParams)),
			slf.Int("rsp", len(rsp)),
			slf.Int("loop", i),
		)

		balances, errs := e.GetTokenBalances(ctx, tmpParams, Options{QPSLimit: qpsLimit})
		if len(errs) != 0 {
			slf.WithError(errs[0]).Errorw("GetTokenBalances err")
		}

		for _, x := range balances {
			delete(paramsMap, fmt.Sprintf("%s-%s-%d", x.Account, x.TokenAddress, x.BlockNumber))
			rsp = append(rsp, x)
		}

		if len(rsp) == totalParamsLen {
			break
		}

		i += 1
	}

	return rsp, nil
}

func (e *Client) GetTokenBalances(ctx context.Context, params GetTokenBalancesParams, options ...Options) (TokenBalancesResults, []error) {

	wg := sync.WaitGroup{}
	lc := sync.Mutex{}
	var errs []error

	rsp := make(TokenBalancesResults, 0, len(params))

	var qpsLimit = 10

	if len(options) > 0 {
		if options[0].QPSLimit > 0 {
			qpsLimit = options[0].QPSLimit
		}
	}

	limiter := ratelimit.New(qpsLimit)

	for _, x := range params {

		limiter.Take()

		wg.Add(1)
		go func(token, account string, decimals int, blockNumber int64) {
			defer helper.DeferFunc(func() {
				wg.Done()
			})

			val, err := e.GetTokenBalance(ctx, blockNumber, token, account)

			if err != nil {

				if helper.ContainsAnyIgnoreCase(err.Error(),
					"no contract code",
					"execution reverted",
				) {
					val = big.NewInt(-1)
					err = nil
				} else {
					lc.Lock()
					errs = append(errs, err)
					lc.Unlock()

					return
				}
			}

			realValue := new(big.Float).Quo(new(big.Float).SetInt(val), big.NewFloat(math.Pow10(decimals)))

			lc.Lock()

			rsp = append(rsp, &TokenBalancesResult{
				TokenAddress: token,
				Account:      account,
				Amount:       realValue,
				BlockNumber:  blockNumber,
			})

			lc.Unlock()

		}(x.TokenAddress, x.Account, x.Decimals, x.BlockNumber)
	}

	wg.Wait()

	if len(errs) > 0 {
		return rsp, errs
	}

	//if len(params) != len(rsp) {
	//	return nil, fmt.Errorf("len(params) != len(rsp)")
	//}

	return rsp, nil
}

func (e *Client) GetTokenBalance(ctx context.Context, blockNumber int64, tokenAddress, account string) (*big.Int, error) {

	if tokenAddress == "eth" {
		balance, err := e.cs.Get().BalanceAt(ctx, common.HexToAddress(account), big.NewInt(blockNumber))
		if err != nil {
			return nil, err
		}

		return balance, nil

	} else {
		erc20, err := contracts.NewErc20(common.HexToAddress(tokenAddress), e.cs.Get())
		if err != nil {
			return nil, err
		}

		balance, err := erc20.BalanceOf(&bind.CallOpts{BlockNumber: big.NewInt(blockNumber)}, common.HexToAddress(account))
		if err != nil {
			return nil, err
		}

		return balance, nil
	}

}

func (e *Client) GetBalancesByAccounts(ctx context.Context, accounts []string) (map[string]*big.Int, error) {

	wg := sync.WaitGroup{}
	lc := sync.Mutex{}

	rsp := map[string]*big.Int{}

	for _, x := range accounts {
		wg.Add(1)
		go func(account string) {
			defer helper.DeferFunc(func() {
				wg.Done()
			})

			val, err := e.cs.Get().BalanceAt(ctx, common.HexToAddress(account), nil)
			if err != nil {
				return
			}

			lc.Lock()
			rsp[account] = val
			lc.Unlock()

		}(x)
	}

	wg.Wait()

	return rsp, nil
}

func (e *Client) GetTokensMetadata(ctx context.Context, addresses []string) (TokenMetadatas, error) {

	var rsp TokenMetadatas

	for _, x := range addresses {
		name, err := e.GetTokenName(ctx, x)
		if err != nil {
			return nil, err
		}

		symbol, err := e.GetTokenSymbol(ctx, x)
		if err != nil {
			return nil, err
		}

		decimals, err := e.GetTokenDecimals(ctx, x)
		if err != nil {
			return nil, err
		}

		rsp = append(rsp, &TokenMetadata{
			Address:  x,
			Symbol:   symbol,
			Name:     name,
			Decimals: decimals,
		})
	}

	return rsp, nil
}

func (e *Client) GetTokenDecimals(ctx context.Context, address string) (int, error) {

	// invalid opcode: INVALID error
	if helper.EqualsIgnoreCase(address, "0xb5a5f22694352c15b00323844ad545abb2b11028") {
		return 18, nil
	}
	if helper.EqualsIgnoreCase(address, "0xc086a2c476d0868ad741088c03252dad293a4b01") {
		return 0, nil
	}
	if helper.EqualsIgnoreCase(address, "0xc351d570c4b1324363d7ce6fbd53e120ead9fae3") {
		return 0, nil
	}
	if helper.EqualsIgnoreCase(address, "0xaad05381e621226e2951b3256d7aca6912cbda6a") {
		return 0, nil
	}

	contractAddress := common.HexToAddress(address)

	erc20, err := contracts.NewErc20(contractAddress, e.cs.Get())
	if err != nil {
		return 0, err
	}
	d, err := erc20.Decimals(nil)

	if err != nil {
		if strings.Contains(err.Error(), "abi: attempting to unmarshall an empty string while arguments are expected") {
			return 0, nil
		}

		if strings.Contains(err.Error(), "invalid jump destination") { // 非erc20
			return 0, nil
		}
		if strings.Contains(err.Error(), "no contract code") { // 非合约
			return -1, nil
		}
		if strings.Contains(err.Error(), "execution reverted") { // 非erc20
			return -1, nil
		}

		slf.WithError(err).Errorw("GetTokenDecimals ", slf.String("address", address))
		return 0, err
	}
	return int(d), nil
}

func (e *Client) GetSender(ctx context.Context, hash string) (string, error) {
	tx, _, err := e.cs.Get().TransactionByHash(ctx, common.HexToHash(hash))
	if err != nil {
		return "", err
	}

	msg, _ := tx.AsMessage(types.NewLondonSigner(tx.ChainId()), nil)
	from := msg.From()

	return strings.ToLower(from.Hex()), nil
}

func (e *Client) GetTokenName(ctx context.Context, address string) (string, error) {

	contractAddress := common.HexToAddress(address)

	erc20, err := contracts.NewErc20(contractAddress, e.cs.Get())
	// 有的ERC20 没有这个方法 eg.0x5ddbee6ad14852d5f78b6eeb6b040391821ff45c
	if err != nil && !strings.Contains(err.Error(), "execution reverted") {
		return "", err
	}

	name, err := erc20.Name(nil)
	// 有的ERC20 没有这个方法 eg.0x0000000000a39bb272e79075ade125fd351887ac
	if err != nil {
		if strings.Contains(err.Error(), "abi: attempting to unmarshall an empty string while arguments are expected") {
			return "", nil
		}
		if strings.Contains(err.Error(), "invalid jump destination") { // 非erc20
			return "", nil
		}
		if strings.Contains(err.Error(), "no contract code") { // 非合约
			return "", nil
		}
		if strings.Contains(err.Error(), "execution reverted") { // 非erc20
			return "", nil
		}
		if strings.Contains(err.Error(), "abi: cannot marshal in to go slice: offset") {
			return "", nil
		}

		return "", err
	}
	return name, nil
}

func (e *Client) GetTokenSymbol(ctx context.Context, address string) (string, error) {

	contractAddress := common.HexToAddress(address)

	erc20, err := contracts.NewErc20(contractAddress, e.cs.Get())
	if err != nil {
		return "", err
	}
	symbol, err := erc20.Symbol(nil)
	// 有的ERC20 没有这个方法 eg.0x0000000000a39bb272e79075ade125fd351887ac
	if err != nil {
		if strings.Contains(err.Error(), "abi: attempting to unmarshall an empty string while arguments are expected") {
			return "", nil
		}
		if strings.Contains(err.Error(), "invalid jump destination") { // 非erc20
			return "", nil
		}
		if strings.Contains(err.Error(), "no contract code") { // 非合约
			return "", nil
		}
		if strings.Contains(err.Error(), "execution reverted") { // 非erc20
			return "", nil
		}
		// eg. 0x28c8d01FF633eA9Cd8fc6a451D7457889E698de6
		if strings.Contains(err.Error(), "abi: cannot marshal in to go slice: offset") { // 非erc20
			return "", nil
		}

		return "", err
	}
	return symbol, nil
}

// function eg. name()
func (e *Client) CallContract(ctx context.Context, contractAddress common.Address, function string) ([]byte, error) {
	s := ethutils.GetFuncSelector(function)
	selector, err := hex.DecodeString(strings.ReplaceAll(s, "0x", ""))
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	result, err := e.cs.Get().CallContract(ctx, ethereum.CallMsg{
		To:   &contractAddress,
		Data: selector,
	}, nil)

	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return result, nil
}
