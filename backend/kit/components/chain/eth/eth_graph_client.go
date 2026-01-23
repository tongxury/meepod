package eth

import (
	"context"
	"encoding/json"
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/kit/components/sdk/httpcli"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.uber.org/ratelimit"
	"strings"
	"sync"
	"time"
)

func NewGraphClient(url string) *GraphClient {
	if !strings.HasSuffix(url, "/graphql") {
		url += "/graphql"
	}

	gc := &GraphClient{url: url}

	return gc
}

type GraphClient struct {
	url string
}

func (e *GraphClient) RequireTradeByHash(ctx context.Context, hash string, excludeCode bool) (*GTrade, error) {
	sql := `
		query trades($hash: Bytes32!) {
		  	transaction (hash: $hash) {
			  	hash
			  	index
			  	from {
					address
					transactionCount
					balance
					code
				}
			  	to {
					address
					transactionCount
					balance
					code
			  	}
				inputData
				value
				status
				createdContract {
					address
					code
				}
				block {
					number
					transactionCount  
					timestamp
				}
			}
		}
	`

	if excludeCode {
		sql = strings.ReplaceAll(sql, "code", "")
	}

	var dst struct {
		Data struct {
			Trade *GTrade `json:"transaction"`
		}
	}

	resultBytes, code, err := httpcli.Client().POST(ctx, e.url, map[string]interface{}{
		"variables": map[string]interface{}{"hash": hash},
		"query":     sql,
	})

	if code != 200 {
		return nil, fmt.Errorf("httpcli.POST: %d", code)
	}

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(resultBytes, &dst); err != nil {
		return nil, err
	}

	if dst.Data.Trade == nil {
		return nil, fmt.Errorf("no trade found by : %s", hash)
	}

	return dst.Data.Trade, nil
}

func (e *GraphClient) GetLogs(ctx context.Context, number int64, addresses []string, topics []string) (GLogs, error) {
	sql := `
		query logs($number: Long!, $addresses: [Address!], $topics: [[Bytes32!]!]) {
		  	logs(filter: { fromBlock: $number, toBlock: $number, addresses: $addresses, topics: $topics}) {
		     	index
				account {
          			address
        		}
        		topics
        		data
				transaction {
					hash
				}
		  	}
		}
	`

	var dst struct {
		Data struct {
			Logs GLogs
		}
	}

	resultBytes, code, err := httpcli.Client().POST(ctx, e.url, map[string]interface{}{
		"variables": map[string]interface{}{"number": hexutil.EncodeUint64(uint64(number)), "addresses": addresses, "topics": topics},
		"query":     sql,
	})

	if code != 200 {
		return nil, fmt.Errorf("httpcli.POST: %d", code)
	}

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(resultBytes, &dst); err != nil {
		return nil, err
	}

	if dst.Data.Logs == nil {
		return nil, fmt.Errorf("dst.Data.Logs == nil")
	}

	return dst.Data.Logs, nil
}

func (e *GraphClient) GetEthTradesByNumberSync(ctx context.Context, numberFrom, numberTo int64, options ...GetTradesOption) (GEthTrades, error) {

	var wg sync.WaitGroup
	var lc sync.Mutex
	var errs []error
	var rsp GEthTrades

	limiter := ratelimit.New(200)

	for i := numberFrom; i <= numberTo; i++ {

		limiter.Take()

		wg.Add(1)
		go func(ctx context.Context, number int64) {
			defer helper.DeferFunc(func() {
				wg.Done()
			})

			trades, err := e.GetEthTradesByNumber(ctx, number, number, options...)
			if err != nil {
				lc.Lock()
				errs = append(errs, err)
				lc.Unlock()
				return
			}
			lc.Lock()
			for _, trade := range trades {
				rsp = append(rsp, trade)
			}
			lc.Unlock()

		}(ctx, i)
	}
	wg.Wait()

	if len(errs) > 0 {
		return nil, errs[0]
	}

	return rsp, nil

}

func (e *GraphClient) GetEthTradesByNumber(ctx context.Context, numberFrom, numberTo int64, options ...GetTradesOption) (GEthTrades, error) {

	sql := blockFullSql
	if len(options) > 0 && options[0].Simple {
		sql = `
			query blocks($from: Long!, $to: Long!) {
			  blocks(from: $from, to: $to) {
				transactions {
				  from {
					address
				  }
				  to {
					address
				  }
				  value
				  status
				}
			  }
			}
		`
	}

	trades, err := e.getTradesBy(ctx, sql, map[string]interface{}{"from": numberFrom, "to": numberTo})
	if err != nil {
		return nil, err
	}

	var rsp GEthTrades
	for _, trade := range trades {
		if trade.Value() > 0 {
			rsp = append(rsp, trade.AsGEthTrade())
		}
	}

	return rsp, nil
}

func (e *GraphClient) GetBlocksByNumberSync(ctx context.Context, fromNumber, toNumber int64, excludeCode bool) (GBlocks, error) {

	var rsp GBlocks
	var errs []error

	wg := sync.WaitGroup{}
	lc := sync.Mutex{}

	for i := fromNumber; i <= toNumber; i++ {

		wg.Add(1)
		go func(ctx2 context.Context, number int64) {
			defer helper.DeferFunc(func() {
				wg.Done()
			})

			blocks, err := e.GetBlocksByNumber(ctx, number, number, excludeCode)
			if err != nil {
				lc.Lock()
				errs = append(errs, err)
				lc.Unlock()
				return
			}

			lc.Lock()
			rsp = append(rsp, blocks...)
			lc.Unlock()
		}(ctx, i)
	}

	wg.Wait()

	if len(errs) != 0 {
		return nil, errs[0]
	}

	return rsp, nil

}

func (e *GraphClient) GetBlockMetasByNumber(ctx context.Context, fromNumber, toNumber int64) (GBlocks, error) {
	sql := `
		query blocks($from: Long!, $to: Long!) {
		  blocks(from: $from, to: $to) {
			number
			transactionCount  
			timestamp
		  }
		}
	`

	var dst struct {
		Data struct {
			Blocks GBlocks
		}
	}

	resultBytes, code, err := httpcli.Client().POST(ctx, e.url, map[string]interface{}{
		"variables": map[string]interface{}{
			"from": fromNumber,
			"to":   toNumber,
		},
		"query": sql,
	})

	if code != 200 {
		return nil, fmt.Errorf("httpcli.POST: %d", code)
	}

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(resultBytes, &dst); err != nil {
		return nil, err
	}

	if dst.Data.Blocks == nil {
		return nil, fmt.Errorf("dst.Data.Block == nil")
	}

	return dst.Data.Blocks, nil
}

func (e *GraphClient) GetBlocksByNumber(ctx context.Context, fromNumber, toNumber int64, excludeCode bool) (GBlocks, error) {
	sql := `
		query blocks($from: Long!, $to: Long!) {
		  blocks(from: $from, to: $to) {
			number
			transactionCount  
			timestamp
			transactions {
				logs {
        			index
					account {
						address
						transactionCount
						balance
					}
					topics
      			}
				hash
				index
				from {
					address
					transactionCount
					balance
				}
				to {
					address
					transactionCount
					balance
					code
				}
				inputData
				value
				status
				createdContract {
					address
				}
				block {
					number
					transactionCount  
					timestamp
				}
			}
		  }
		}
	`

	if excludeCode {
		sql = strings.ReplaceAll(sql, "code", "")
	}

	var dst struct {
		Data struct {
			Blocks GBlocks
		}
	}

	resultBytes, code, err := httpcli.Client().POST(ctx, e.url, map[string]interface{}{
		"variables": map[string]interface{}{
			"from": fromNumber,
			"to":   toNumber,
		},
		"query": sql,
	})

	if code != 200 {
		return nil, fmt.Errorf("httpcli.POST: %d", code)
	}

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(resultBytes, &dst); err != nil {
		return nil, err
	}

	if dst.Data.Blocks == nil {
		return nil, fmt.Errorf("dst.Data.Block == nil")
	}

	return dst.Data.Blocks, nil
}

type GetTradesOption struct {
	WithoutCode bool
	Simple      bool
}

func (e *GraphClient) GetTradesByNumberInBatches(ctx context.Context, fromNumber, toNumber int64, batchSize int64, options ...GetTradesOption) (GTrades, error) {

	wg := sync.WaitGroup{}
	lc := sync.Mutex{}
	var errs []error

	if batchSize == 0 {
		batchSize = 1
	}

	limiter := ratelimit.New(2, ratelimit.Per(5*time.Second))

	var rsp GTrades
	for from := fromNumber; from <= toNumber; from += batchSize {
		to := from + batchSize - 1

		if to > toNumber {
			to = toNumber
		}

		limiter.Take()

		wg.Add(1)

		go func(ctx context.Context, from, to int64) {

			defer helper.DeferFunc(func() {
				wg.Done()
			})

			slf.Debugw("GetTradesByNumberInBatches ...", slf.Int64("from", from), slf.Int64("to", to))

			trades, err := e.GetTradesByNumber(ctx, from, to, options...)
			if err != nil {
				lc.Lock()
				errs = append(errs, err)
				lc.Unlock()
				return
			}

			lc.Lock()
			rsp = append(rsp, trades...)
			lc.Unlock()

		}(ctx, from, to)

	}

	wg.Wait()

	if len(errs) > 0 {
		return nil, errs[0]
	}

	return rsp, nil
}

func (e *GraphClient) GetTradesByNumber(ctx context.Context, fromNumber, toNumber int64, options ...GetTradesOption) (GTrades, error) {

	var withoutCode = false
	if len(options) > 0 {
		withoutCode = options[0].WithoutCode
	}
	/*
	  balance： 当前已同步的最高区块的 eth balance, 并非指定区块的余额
	*/
	sql := `
		query blocks($from: Long!, $to: Long!) {
		  blocks(from: $from, to: $to) {
			transactions {
			  hash
			  index
			  from {
				address
				transactionCount
				balance
				code
			  }
			  to {
				address
				transactionCount
				balance
				code
			  }
			  inputData
			  value
			  status
			  createdContract {
				address
				code
			  }
				block {
					number
					transactionCount  
					timestamp
				}
			}
		  }
		}
	`

	if withoutCode {
		sql = strings.ReplaceAll(sql, "code", "")
	}

	return e.getTradesBy(ctx, sql, map[string]interface{}{"from": fromNumber, "to": toNumber})
}

func (e *GraphClient) getTradesBy(ctx context.Context, sql string, params map[string]interface{}) (GTrades, error) {

	var dst struct {
		Data struct {
			Blocks GBlocks
		}
	}

	resultBytes, code, err := httpcli.Client().POST(ctx, e.url, map[string]interface{}{
		"variables": params,
		"query":     sql,
	})

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("httpcli.POST: %d, %s", code, string(resultBytes))
	}

	if err := json.Unmarshal(resultBytes, &dst); err != nil {
		return nil, err
	}

	var rsp GTrades
	for _, block := range dst.Data.Blocks {
		for _, trade := range block.Trades {
			if trade.IsSuccess() {
				rsp = append(rsp, trade)
			}
		}

	}

	return rsp, nil
}

var (
	blockFullSql string = `
		query blocks($from: Long!, $to: Long!) {
		  blocks(from: $from, to: $to) {
			number
			transactionCount  
			timestamp
			transactions {
				logs {
        			index
					account {
						address
						transactionCount
						balance
					}
					topics
      			}
				hash
				index
				from {
					address
					transactionCount
					balance
				}
				to {
					address
					transactionCount
					balance
					code
				}
				inputData
				value
				status
				createdContract {
					address
				}
				block {
					number
					transactionCount  
					timestamp
				}
			}
		  }
		}
	`
)
