package chCtxs

import (
	"context"
	"gitee.com/meepo/backend/kit/components/chain/changenow"
)

func NewTrx2EthUSDTCtx() *ChCtx {
	return &ChCtx{
		Url: "https://api.changenow.io",
		//ApiKey:       "12cc519749942fa36b62e2e205a4f99791c0fd1e89b865b76c6bf379646ccab7",
		ApiKey:       "2a48db5d876d6f0ba42fe490394125d952bb2d4a3ce6ca512a3d3e246ab92e07",
		FromCurrency: "usdt",
		ToCurrency:   "usdt",
		FromNetwork:  "trx",
		ToNetwork:    "eth",
	}
}

func NewTrx2EtcCtx(currency string) *ChCtx {

	return &ChCtx{
		Url: "https://api.changenow.io",
		//ApiKey:       "12cc519749942fa36b62e2e205a4f99791c0fd1e89b865b76c6bf379646ccab7",
		ApiKey:       "2a48db5d876d6f0ba42fe490394125d952bb2d4a3ce6ca512a3d3e246ab92e07",
		FromCurrency: currency,
		ToCurrency:   currency,
		FromNetwork:  "trx",
		ToNetwork:    "eth",
	}
}

type ChCtx struct {
	Url          string
	ApiKey       string
	FromCurrency string
	ToCurrency   string
	FromNetwork  string
	ToNetwork    string
}

func (c *ChCtx) client() *changenow.ChClient {
	cli, _ := changenow.NewClient(c.Url, c.ApiKey)
	return cli
}

func (c *ChCtx) EstimateFee(ctx context.Context, sendAmount string) (string, error) {
	gotAmount, err := c.client().EstimateFee(ctx, c.FromCurrency, c.ToCurrency, c.FromNetwork, c.ToNetwork, sendAmount)
	if err != nil {
		return "", err
	}
	return gotAmount, nil
}

func (c *ChCtx) CreateOrder(ctx context.Context, amount, toAddress string) (*changenow.ChOrder, error) {
	return c.client().CreateOrder(ctx, c.FromCurrency, c.ToCurrency, c.FromNetwork, c.ToNetwork, amount, toAddress)
}

func (c *ChCtx) GetOrderStatus(ctx context.Context, orderId string) (*changenow.OrderStatus, error) {
	return c.client().GetOrderStatus(ctx, orderId)
}
