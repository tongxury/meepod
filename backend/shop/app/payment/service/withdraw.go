package service

import (
	"context"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/kit/components/sdk/helper/timed"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/kit/services/util/gind/errorx"
	"gitee.com/meepo/backend/shop/app/payment/db"
	coredb "gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/types"
	"github.com/go-pg/pg/v10"
)

type WithdrawService struct {
}

func (t *WithdrawService) Cancel(ctx context.Context, userId, storeId, orderId string) error {

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		canceled, amount, err := new(db.Withdraw).UpdateToCanceled(ctx, tx, orderId, userId)
		if err != nil {
			return err
		}

		if canceled {
			_, err := tx.Model((*db.Account)(nil)).Context(ctx).
				Where("user_id = ?", userId).
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

func (t *WithdrawService) ListWithdraws(ctx context.Context, userId, storeId string, page, size int64) (types.PaymentWithdraws, int64, error) {

	orders, total, err := new(db.Withdraw).List(ctx, userId, storeId, nil, page, size)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	rsp, err := t.Assemble(ctx, orders)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	return rsp, total, nil
}

func (t *WithdrawService) AddWithdraw(ctx context.Context, storeId, userId string, amount float64) (*types.PaymentWithdraw, error) {

	//orderCategory := enum.OrderCategory_Withdraw.Value

	exists, err := new(db.Withdraw).ExistsByStatus(ctx, userId, storeId, enum.WithdrawStatus_Submitted.Value)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	if exists {
		return nil, errorx.UserMessage("订单已存在")
	}

	dbWithdraw := &db.Withdraw{
		UserId:  userId,
		StoreId: storeId,
		Amount:  amount,
		Status:  enum.WithdrawStatus_Submitted.Value,
	}

	err = comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {
		inserted, err := dbWithdraw.Insert(ctx, tx)
		if err != nil {
			return err
		}

		// 冻结待提现的金额
		if inserted {
			_, err := tx.Model((*db.Account)(nil)).Context(ctx).
				Where("user_id = ?", userId).
				Where("store_id = ?", storeId).
				Set("balance = balance - ?", amount).Update()

			if err != nil {
				return xerror.Wrap(err)
			}
		}

		return nil
	})

	Withdraws, err := t.Assemble(ctx, db.Withdraws{dbWithdraw})
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return Withdraws[0], nil
}

func (t *WithdrawService) FindById(ctx context.Context, id string) (*types.PaymentWithdraw, error) {

	dbOrder, err := new(db.Withdraw).FindById(ctx, id)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	orders, err := t.Assemble(ctx, db.Withdraws{dbOrder})
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return orders[0], nil
}

func (t *WithdrawService) Assemble(ctx context.Context, orders db.Withdraws) (types.PaymentWithdraws, error) {

	_, storeIds, userIds := orders.Ids()

	dbStores, err := new(coredb.Store).ListByIds(ctx, storeIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	dbStoresMap := dbStores.AsMap()

	_, ownerIds := dbStores.Ids()
	userIds = append(userIds, ownerIds...)
	dbUsers, err := new(coredb.User).ListByIds(ctx, userIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	dbUsersMap := dbUsers.AsMap()

	var rsp types.PaymentWithdraws

	for _, x := range orders {

		user := types.FromDbUser(dbUsersMap[x.UserId])

		dnStore := dbStoresMap[x.StoreId]
		store := types.FromDbStore(dnStore, dbUsersMap[dnStore.OwnerId])

		y := types.PaymentWithdraw{
			Id:          x.Id,
			User:        user,
			Store:       store,
			Amount:      x.Amount,
			CreatedAt:   timed.SmartTime(x.CreatedAt.Unix()),
			CreatedAtTs: x.CreatedAt.Unix(),
			Status:      enum.WithdrawStatus(x.Status),
			Remark:      x.Extra.Remark,
			Image:       x.Extra.Image,
			Acceptable:  helper.InSlice(x.Status, enum.AcceptableWithdrawStatus),
			Rejectable:  helper.InSlice(x.Status, enum.RejectableWithdrawStatus),
			Cancelable:  helper.InSlice(x.Status, enum.CancelableWithdrawStatus),
		}

		rsp = append(rsp, &y)
	}

	return rsp, nil
}
