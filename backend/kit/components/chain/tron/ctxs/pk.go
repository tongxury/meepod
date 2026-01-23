package tronCtxs

import (
	"context"
	"gitee.com/meepo/backend/kit/components/chain/tron/tronclient"
	tronutils "gitee.com/meepo/backend/kit/components/chain/tron/utils"
	"gitee.com/meepo/backend/kit/components/sdk/bn"
	"math/big"
)

type PkeyCtx struct {
	Pk      string
	EPK     string
	Url     string
	Address string
}

//func NewPkeyCtx(epk string, url string) (*PkeyCtx, error) {
//	pk, err := encryptor.DESCBC(app.SaltKey).Decrypt(epk)
//	if err != nil {
//		return nil, err
//	}
//
//	address, err := tronutils.PrivateKeyToAddress(pk)
//	if err != nil {
//		return nil, err
//	}
//
//	return &PkeyCtx{
//		Address: address,
//		Pk:      pk,
//		EPK:     epk,
//		Url:     url,
//	}, nil
//}

func NewPkCtx(pk string, url string) (*PkeyCtx, error) {
	address, err := tronutils.PrivateKeyToAddress(pk)
	if err != nil {
		return nil, err
	}

	return &PkeyCtx{
		Address: address,
		Pk:      pk,
		Url:     url,
	}, nil
}

func (t *PkeyCtx) Client() *tronclient.TronClient {
	// todo
	c, _ := tronclient.NewTronClient(t.Url, "dd5d773e-4392-4bcb-95f1-811542e16dd7")
	return c
}

func (t *PkeyCtx) Transfer(ctx context.Context, from, to string, amount *big.Int) (string, error) {
	return t.Client().Transfer(ctx, from, to, amount.Int64(), t.Pk)
}

func (t *PkeyCtx) TransferTRC20(ctx context.Context, from, to, contract string, amount *big.Int) (string, error) {
	return t.Client().TransferTRC20(ctx, from, to, contract, amount, t.Pk)
}

//func (t *PkeyCtx) TransferUSDT(ctx context.Context, from, to string, val string) (string, error) {
//
//	num := bn.FromStr(val, 6)
//
//	return t.Client().TransferTRC20(ctx, from, to, "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t", num.RawVal(), t.Pk)
//}

func (t *PkeyCtx) TRC20Balance(ctx context.Context, contractAddress string) (*bn.BN, error) {
	balance, err := t.Client().GetTRC20Balance(ctx, t.Address, contractAddress)
	if err != nil {
		return nil, err
	}

	decimals, err := t.Client().GrpcClient().TRC20GetDecimals(contractAddress)
	if err != nil {
		return nil, err
	}

	return bn.FromRaw(balance, uint(decimals.Uint64())), nil

}

//func (t *PkeyCtx) USDTBalance(ctx context.Context) (*bn.BN, error) {
//	bg, err := t.Client().GetTRC20Balance(ctx, t.Address, "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t")
//	if err != nil {
//		return nil, err
//	}
//	return bn.BigNumber(bg, 6), nil
//}

func (t *PkeyCtx) Balance(ctx context.Context) (*bn.BN, error) {
	b, err := t.Client().GetBalance(ctx, t.Address)
	if err != nil {
		if err.Error() == "account not found" {
			return bn.BigNumber(big.NewInt(0), 6), nil
		}
		return nil, err
	}
	return bn.BigNumber(big.NewInt(b), 6), nil
}
