package keeperservice

import (
	"context"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/app/payment/db"
	"gitee.com/meepo/backend/shop/app/payment/service"
	"gitee.com/meepo/backend/shop/core/types"
	"github.com/go-pg/pg/v10"
)

type WithdrawService struct {
	service.WithdrawService
}

func (t *WithdrawService) ListWithdraws(ctx context.Context, keeperId, storeId string, page, size int64) (types.PaymentWithdraws, int64, error) {

	orders, total, err := new(db.Withdraw).List(ctx, "", storeId, nil, page, size)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	rsp, err := t.Assemble(ctx, orders)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	return rsp, total, nil
}

func (t *WithdrawService) Accept(ctx context.Context, keeperId, storeId, id, image string) error {

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		_, err := new(db.Withdraw).UpdateToAccepted(ctx, tx, id, image)
		if err != nil {
			return err
		}

		return nil
	})

	return xerror.Wrap(err)
}

func (t *WithdrawService) Reject(ctx context.Context, keeperId, storeId, id, remark string) error {

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		var tmp db.Withdraws
		err := tx.Model(&tmp).Context(ctx).Where("id = ?", id).Select()
		if err != nil {
			return err
		}

		if len(tmp) == 0 {
			return nil
		}

		// todo
		canceled, amount, err := new(db.Withdraw).UpdateToRejected(ctx, tx, id, remark)
		if err != nil {
			return err
		}

		if canceled {
			_, err := tx.Model((*db.Account)(nil)).Context(ctx).
				Where("user_id = ?", tmp[0].UserId).
				Where("store_id = ?", storeId).
				Set("balance = balance + ?", amount).Update()

			if err != nil {
				return xerror.Wrap(err)
			}

		}

		return nil
	})

	return xerror.Wrap(err)
}
