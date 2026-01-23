package swap

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"io"
	"net/http"
)

func NewSushiApolloClient() *SushiApolloClient {
	return &SushiApolloClient{
		//https://thegraph.com/hosted-service/subgraph/sushiswap/exchange
		EntrypointFormat: "https://api.thegraph.com/subgraphs/name/sushiswap/exchange",
	}
}

type SushiApolloClient struct {
	EntrypointFormat string
}

func (t *SushiApolloClient) ListPairsByBlockNumber(ctx context.Context, blockNumberFrom, blockNumberTo uint64) (APairs, error) {

	params := map[string]interface{}{
		"blockNumberGte": blockNumberFrom,
		"blockNumberLte": blockNumberTo,
		"first":          1000,
	}

	sql := `
		 query pairs($blockNumberGte: Int!, $blockNumberLte: Int!, $first: Int) {
				pairs(first:  $first, where: {block_gte: $blockNumberGte, block_lte: $blockNumberLte}) {
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
					block
					timestamp
				}
			  }
	`

	var dst struct {
		Data struct {
			Pairs SushiPairs
		}
	}

	err := t.execute(ctx, t.EntrypointFormat, sql, params, &dst)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	pairs := dst.Data.Pairs.AsAPairs()

	for _, pair := range pairs {
		pair.Swap()
		pair.Version = "sushiswap"
	}

	return pairs, nil
}

func (t *SushiApolloClient) ListSwapsByBlockNumber(ctx context.Context, blockNumberGte, blockNumberLte uint64) (ASwaps, error) {

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
                      token0 { id, name, decimals, symbol }
                      token1 { id, name, decimals, symbol }
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
			Swaps SushiSwaps
		}
	}

	err := t.execute(ctx, t.EntrypointFormat, sql, params, &dst)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	swaps := RawV2Swaps(dst.Data.Swaps).AsASwaps()

	for _, swap := range swaps {
		swap.Version = "sushiswap"
	}

	return swaps, nil
}

func (t *SushiApolloClient) execute(ctx context.Context, url, sql string, params map[string]interface{}, dst any) error {

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
