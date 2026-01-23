package keeperservice

import (
	"context"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/kit/services/util/gind/errorx"
	"gitee.com/meepo/backend/shop/app/payment/db"
	"gitee.com/meepo/backend/shop/app/payment/service"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/types"
	"github.com/go-pg/pg/v10"
)

type AccountService struct {
	service.AccountService
}

func (t *AccountService) Decr(ctx context.Context, storeId, id string, amount float64, remark string) error {

	wallet, err := new(db.Account).RequireByIdAndStoreId(ctx, id, storeId)
	if err != nil {
		return xerror.Wrap(err)
	}

	if wallet.Balance < amount {
		return errorx.UserMessage("余额不足")
	}

	err = comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		dbPayment := &db.Payment{
			UserId:   wallet.UserId,
			StoreId:  storeId,
			Amount:   amount,
			Status:   enum.PaymentStatus_Payed.Value,
			Category: enum.PaymentCategory_DecrByKeeper.Value,
			Extra:    db.PaymentExtra{Remark: remark},
		}

		err := dbPayment.Insert(ctx, tx)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return xerror.Wrap(err)
	}

	return nil
}

func (t *AccountService) GetStoreSummary(ctx context.Context, storeId string) (*types.AccountSummary, error) {

	userCount, totalBalance, err := new(db.Account).GetStoreSummary(ctx, storeId)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return &types.AccountSummary{
		UserCount:    userCount,
		TotalBalance: totalBalance,
	}, nil
}

func (t *AccountService) ListAccounts(ctx context.Context, storeId string, page, size int64) (types.Accounts, int64, error) {

	dbAccounts, total, err := new(db.Account).List(ctx, storeId, page, size)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	rsp, err := t.AssembleAccounts(ctx, dbAccounts)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	return rsp, total, nil
}
