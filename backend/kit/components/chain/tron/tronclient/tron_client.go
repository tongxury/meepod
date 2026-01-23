package tronclient

import (
	"context"
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"math/big"
)

func NewTronClient(url, apiKey string) (*TronClient, error) {

	c := client.NewGrpcClient(url)
	_ = c.SetAPIKey(apiKey)
	if err := c.Start(grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		return nil, err
	}
	return &TronClient{c: c}, nil
}

type TronClient struct {
	c *client.GrpcClient
}

func (t *TronClient) GrpcClient() *client.GrpcClient {
	return t.c
}

func (t *TronClient) GetBalance(ctx context.Context, address string) (int64, error) {

	account, err := t.c.GetAccount(address)
	if err != nil {
		return 0, err
	}

	return account.GetBalance(), nil
}

func (t *TronClient) signAndBroadcast(ctx context.Context, rawTx *core.Transaction, privateKey string) (*api.TransactionExtention, error) {
	pkBytes, err := common.FromHex(privateKey)
	if err != nil {
		return nil, err
	}

	signedTx, err := t.c.Client.AddSign(ctx, &core.TransactionSign{
		Transaction: rawTx,
		PrivateKey:  pkBytes,
	})
	if err != nil {
		return nil, err
	}
	if signedTx.GetTransaction() == nil {
		return nil, fmt.Errorf("%v", signedTx.Result.Message)
	}

	resp, err := t.c.Client.BroadcastTransaction(ctx, signedTx.GetTransaction())
	if err != nil {
		return nil, err
	}

	if resp.Code != api.Return_SUCCESS {
		return nil, fmt.Errorf("BroadcastTransaction result: %d", resp.Code)
	}

	return signedTx, nil
}

func (t *TronClient) Transfer(ctx context.Context, from, to string, amount int64, privateKey string) (string, error) {

	rawTx, err := t.c.Transfer(from, to, amount)
	if err != nil {
		return "", err
	}

	signedTx, err := t.signAndBroadcast(ctx, rawTx.GetTransaction(), privateKey)
	if err != nil {
		return "", err
	}

	return common.Bytes2Hex(signedTx.GetTxid()), nil

}

func (t *TronClient) GetTRC20Balance(ctx context.Context, address, contractAddress string) (*big.Int, error) {
	return t.c.TRC20ContractBalance(address, contractAddress)
}

func (t *TronClient) TransferTRC20(ctx context.Context, from, to, contract string, amount *big.Int, privateKey string) (string, error) {

	rawTx, err := t.c.TRC20Send(from, to, contract, amount, 1000000000)
	if err != nil {
		return "", err
	}

	signedTx, err := t.signAndBroadcast(ctx, rawTx.GetTransaction(), privateKey)
	if err != nil {
		return "", err
	}
	return common.Bytes2Hex(signedTx.GetTxid()), nil
}
