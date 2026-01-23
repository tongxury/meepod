package changenow

import (
	"context"
	"encoding/json"
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/httpcli"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
)

func NewClient(url, apiKey string) (*ChClient, error) {
	// https://api.changenow.io
	// 12cc519749942fa36b62e2e205a4f99791c0fd1e89b865b76c6bf379646ccab7
	return &ChClient{url: url, apiKey: apiKey}, nil
}

type ChClient struct {
	url    string
	apiKey string
}

func (c *ChClient) CreateOrder(ctx context.Context, fromCurrency, toCurrency, fromNetwork, toNetwork, amount, toAddress string) (*ChOrder, error) {

	req := &reqBody{
		FromCurrency: fromCurrency,
		ToCurrency:   toCurrency,
		FromNetwork:  fromNetwork,
		ToNetwork:    toNetwork,
		FromAmount:   amount,
		Address:      toAddress,
	}

	rspBytes, _, err := httpcli.Client().POST(ctx, c.url+"/v2/exchange", req, httpcli.Options{
		Headers: map[string]string{
			"x-changenow-api-key": c.apiKey,
		},
	})

	if err != nil {
		return nil, xerror.Wrap(err)
	}

	var rsp ChOrder
	_ = json.Unmarshal(rspBytes, &rsp)

	return &rsp, nil
}

func (c *ChClient) EstimateFee(ctx context.Context, fromCurrency, toCurrency, fromNetwork, toNetwork, fromAmount string) (string, error) {
	//https: //api.changenow.io/v2/exchange/network-fee?fromCurrency=usdt&toCurrency=usdt&fromNetwork=eth&toNetwork=eth&fromAmount=100&convertedCurrency=usd&convertedNetwork=usd
	//https://api.changenow.io/v2/exchange/estimated-amount?fromCurrency=btc&toCurrency=usdt&fromAmount=0.1&toAmount=&fromNetwork=btc&toNetwork=eth&flow=fixed-rate&type=&useRateId=

	url := fmt.Sprintf("%s/v2/exchange/estimated-amount?fromCurrency=%s&toCurrency=%s&fromNetwork=%s&toNetwork=%s&fromAmount=%s",
		c.url, fromCurrency, toCurrency, fromNetwork, toNetwork, fromAmount)

	rspBytes, _, err := httpcli.Client().GET(ctx, url, httpcli.Options{
		Headers: map[string]string{
			"x-changenow-api-key": c.apiKey,
		},
	})

	if err != nil {
		return "", xerror.Wrap(err)
	}

	var rsp EstimateFeeResp
	_ = json.Unmarshal(rspBytes, &rsp)

	return conv.String(rsp.ToAmount), nil
}

func (c *ChClient) GetOrderStatus(ctx context.Context, orderId string) (*OrderStatus, error) {

	rspBytes, _, err := httpcli.Client().GET(ctx, c.url+"/v2/exchange/by-id?id="+orderId, httpcli.Options{
		Headers: map[string]string{
			"x-changenow-api-key": c.apiKey,
		},
	})

	if err != nil {
		return nil, xerror.Wrap(err)
	}

	var rsp OrderStatus
	_ = json.Unmarshal(rspBytes, &rsp)

	return &rsp, nil
}
