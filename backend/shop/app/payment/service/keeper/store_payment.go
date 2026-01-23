package keeperservice

import (
	"context"
	"gitee.com/meepo/backend/kit/components/sdk/helper/timed"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/app/payment/db"
	coredb "gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/types"
)

type StorePaymentService struct {
}

func (t *StorePaymentService) ListPayments(ctx context.Context, month, keeperId, storeId string, page, size int64) (types.StorePayments, int64, error) {

	dbPayments, total, err := new(db.StorePayment).List(ctx, db.ListStorePaymentsParams{
		StoreId: storeId, Month: month,
		Page: page, Size: size,
	})
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	payments, err := t.Assemble(ctx, dbPayments)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	return payments, total, nil
}

func (t *StorePaymentService) Assemble(ctx context.Context, payments db.StorePayments) (types.StorePayments, error) {

	_, storeIds := payments.Ids()
	dbStores, err := new(coredb.Store).ListByIds(ctx, storeIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	dbStoresMap := dbStores.AsMap()

	_, ownerIds := dbStores.Ids()

	dbUsers, err := new(coredb.User).ListByIds(ctx, ownerIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	dbUsersMap := dbUsers.AsMap()

	var rsp types.StorePayments

	for _, x := range payments {

		store := dbStoresMap[x.StoreId]
		storeOwner := dbUsersMap[store.OwnerId]

		y := types.StorePayment{
			Id:          x.Id,
			Store:       types.FromDbStore(store, storeOwner),
			Category:    enum.StorePaymentCategory(x.Category),
			BizCategory: enum.BizCategory(x.BizCategory),
			BizId:       x.BizId,
			Amount:      x.Amount,
			CreatedAtTs: x.CreatedAt.Unix(),
			CreatedAt:   timed.SmartTime(x.CreatedAt.Unix()),
			Status:      enum.StorePaymentStatus(x.Status),
		}

		rsp = append(rsp, &y)
	}

	return rsp, nil

}
