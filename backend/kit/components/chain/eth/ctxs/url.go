package ethCtxs

import "github.com/ethereum/go-ethereum/ethclient"

type UrlCtx struct {
	Url string
}

func (u *UrlCtx) GetClient() *ethclient.Client {
	client, _ := ethclient.Dial(u.Url)
	return client
}
