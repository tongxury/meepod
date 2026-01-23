package eth

import (
	"context"
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"github.com/ethereum/go-ethereum/ethclient"
)

func NewEthClients(urls []string) (Clients, error) {

	var ethClients Clients

	for _, url := range urls {

		if url == "" {
			return nil, fmt.Errorf("invalid url :%s", url)
		}

		client, err := ethclient.Dial(url)
		if err != nil {
			return nil, err
		}

		_, err = client.ChainID(context.Background())
		if err != nil {
			return nil, err
		}

		ethClients = append(ethClients, client)
	}

	return ethClients, nil

}

func NewClient(url string) (*ethclient.Client, error) {

	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	_, err = client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	return client, err
}

type Clients []*ethclient.Client

func (es Clients) Get() *ethclient.Client {

	l := len(es)
	return es[helper.GetRandom(0, l-1)]
}
