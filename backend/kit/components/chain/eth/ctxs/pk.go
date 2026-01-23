package ethCtxs

import (
	"context"
	"gitee.com/meepo/backend/kit/components/chain/eth/contracts"
	ethutils "gitee.com/meepo/backend/kit/components/chain/eth/utils"
	"gitee.com/meepo/backend/kit/components/sdk/bn"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
)

type PkeyCtx struct {
	Pk string
	//EPK     string
	Url     string
	Address common.Address
}

func NewPkeyCtxByPk(pk string, url string) (*PkeyCtx, error) {
	addr, err := ethutils.PrivateKeyToAddress(pk)
	if err != nil {
		return nil, err
	}

	return &PkeyCtx{
		Address: *addr,
		Pk:      pk,
		Url:     url,
	}, nil
}

//func NewPkeyCtx(epk string, url string) (*PkeyCtx, error) {
//	pk, err := encryptor.DESCBC(app.SaltKey).Decrypt(epk)
//	if err != nil {
//		return nil, err
//	}
//	addr, err := ethutils.PrivateKeyToAddress(pk)
//	if err != nil {
//		return nil, err
//	}
//
//	return &PkeyCtx{
//		Address: *addr,
//		Pk:      pk,
//		EPK:     epk,
//		Url:     url,
//	}, nil
//}

func (e *PkeyCtx) AddressHex() string {
	return strings.ToLower(e.Address.Hex())
}

func (e *PkeyCtx) WaitFor(ctx context.Context, tx *types.Transaction) (bool, error) {

	//e.Client().TransactionReceipt(ctx, tx.Hash())
	return false, nil
}

func (e *PkeyCtx) Transfer(ctx context.Context, toAddress string, amount *big.Int) (*types.Transaction, error) {

	client := e.Client()

	nonce, err := client.PendingNonceAt(ctx, e.Address)
	if err != nil {
		return nil, err
	}

	gasLimit := uint64(21000) // in units
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	tx := types.NewTransaction(nonce, common.HexToAddress(toAddress), amount, gasLimit, gasPrice, nil)

	chainID, err := client.NetworkID(ctx)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(e.Pk)
	if err != nil {
		return nil, err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, err
	}

	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func (e *PkeyCtx) Client() *ethclient.Client {
	client, _ := ethclient.Dial(e.Url)
	return client
}

func (e *PkeyCtx) ChainId(ctx context.Context) (*big.Int, error) {
	return e.Client().ChainID(ctx)
}

func (e *PkeyCtx) ERC20Balance(tokenAddress common.Address) (*bn.BN, error) {

	erc20, err := contracts.NewErc20(tokenAddress, e.Client())
	if err != nil {
		return nil, err
	}

	balance, err := erc20.BalanceOf(nil, e.Address)
	if err != nil {
		return nil, err
	}

	decimals, err := erc20.Decimals(nil)
	if err != nil {
		return nil, err
	}

	return bn.BigNumber(balance, uint(decimals)), nil
}

func (e *PkeyCtx) Balance(ctx context.Context) (*big.Int, error) {

	at, err := e.Client().BalanceAt(ctx, e.Address, nil)
	if err != nil {
		return nil, err
	}

	return at, nil
}

func (e *PkeyCtx) Auth(ctx context.Context) (*bind.TransactOpts, error) {

	privateKey, err := crypto.HexToECDSA(e.Pk)
	if err != nil {
		return nil, err
	}
	chainId, err := e.ChainId(ctx)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, err
	}

	return auth, nil
}
