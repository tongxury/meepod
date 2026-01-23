package adminserivce

import (
	"context"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/app/payment/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"github.com/go-pg/pg/v10"
)

type StorePaymentService struct {
}

func (t *StorePaymentService) Topup(ctx context.Context, storeId string, amount float64) error {

	payments := db.StorePayment{
		StoreId:  storeId,
		Category: enum.StorePaymentCategory_TopUp.Value,
		Amount:   amount,
		Status:   enum.StorePaymentStatus_Confirmed.Value,
	}

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {
		err := payments.Insert(ctx, tx)
		return err
	})

	return xerror.Wrap(err)
}
