package service

import (
	"context"
	"fmt"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/helper/timed"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/app/payment/db"
	coredb "gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/types"
	"github.com/go-pg/pg/v10"
)

type PaymentService struct {
}

func (t *PaymentService) ListPayments(ctx context.Context, userId, storeId string, page, size int64) (types.Payments, int64, error) {

	dbPayments, total, err := new(db.Payment).List(ctx, userId, storeId, page, size)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	payments, err := t.Assemble(ctx, dbPayments)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	return payments, total, nil
}

func (t *PaymentService) Assemble(ctx context.Context, payments db.Payments) (types.Payments, error) {

	_, storeIds, userIds := payments.Ids()
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

	var rsp types.Payments

	for _, x := range payments {

		user := types.FromDbUser(dbUsersMap[x.UserId])

		dnStore := dbStoresMap[x.StoreId]
		store := types.FromDbStore(dnStore, dbUsersMap[dnStore.OwnerId])

		y := types.Payment{
			Id:          x.Id,
			User:        user,
			Store:       store,
			BizId:       x.BizId,
			BizCategory: enum.BizCategory(x.BizCategory),
			Category:    enum.PaymentCategory(x.Category),
			Amount:      -x.Amount,
			CreatedAtTs: x.CreatedAt.Unix(),
			CreatedAt:   timed.SmartTime(x.CreatedAt.Unix()),
			Status:      enum.PaymentStatus(x.Status),
			Remark:      x.Extra.Remark,
		}

		rsp = append(rsp, &y)
	}

	return rsp, nil

}

func (t *PaymentService) ListPayMethods(ctx context.Context, storeId, userId string, amount float64) (types.PayMethods, error) {
	account, err := new(db.Account).RequireByUserIdAndStoreId(ctx, userId, storeId)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	var rsp types.PayMethods

	if account.Balance > amount {
		rsp = append(rsp, &types.PayMethod{
			Id:   enum.PayMethod_Account.Value,
			Name: enum.PayMethod_Account.Name,
		})
	}

	rsp = append(rsp,
		&types.PayMethod{
			Id:    enum.PayMethod_Alipay.Value,
			Name:  enum.PayMethod_Alipay.Name,
			Color: enum.PayMethod_Alipay.Color,
		},
		&types.PayMethod{
			Id:    enum.PayMethod_Wechat.Value,
			Name:  enum.PayMethod_Wechat.Name,
			Color: enum.PayMethod_Wechat.Color,
		},
	)
	//rsp = append(rsp, &types.PayMethod{
	//	Id:   "alipay",
	//	Name: "支付宝",
	//})

	return rsp, nil
}

func (t *PaymentService) getAmount(ctx context.Context, id, userId, category string) (float64, error) {

	switch category {
	case enum.BizCategory_Order.Value:
		order, err := new(coredb.Order).RequireByIdAndCreatorId(ctx, id, userId)
		if err != nil {
			return 0, xerror.Wrap(err)
		}
		return order.Amount, nil
	case enum.BizCategory_GroupShare.Value:
		order, err := new(coredb.OrderGroupShare).RequireByIdAndCreatorId(ctx, id, userId)
		if err != nil {
			return 0, xerror.Wrap(err)
		}
		return order.Amount, nil
	}

	return 0, fmt.Errorf("invalid id: %s, category: %s", id, category)
}

// redis 队列通知 todo
func (t *PaymentService) updatePayed(ctx context.Context, tx *pg.Tx, id, category string) (bool, error) {
	switch category {
	case enum.BizCategory_Order.Value:
		updated, err := new(coredb.Order).UpdateToPayed(ctx, tx, id, false)
		if err != nil {
			return false, xerror.Wrap(err)
		}

		return updated, nil
	case enum.BizCategory_GroupShare.Value:
		updated, err := new(coredb.OrderGroupShare).UpdateToPayed(ctx, tx, id)
		if err != nil {
			return false, xerror.Wrap(err)
		}
		return updated, nil
	}

	return false, fmt.Errorf("invalid id: %s, category: %s", id, category)
}

func (t *PaymentService) PayByQrCode(ctx context.Context, storeId, userId, orderId, category, clientIp string) (*types.PaymentTopup, error) {

	amount, err := t.getAmount(ctx, orderId, userId, category)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	up, err := new(TopupService).AddBuyingTopUp(ctx, storeId, userId, orderId, category, amount, clientIp)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return up, nil
}

// 需要保证此方法不可重入
func (t *PaymentService) PayByAccount(ctx context.Context, storeId, userId, orderId, bizCategory string) (bool, error) {

	amount, err := t.getAmount(ctx, orderId, userId, bizCategory)
	if err != nil {
		return false, xerror.Wrap(err)
	}

	// 检查余额
	account, err := new(db.Account).RequireByUserIdAndStoreId(ctx, userId, storeId)
	if err != nil {
		return false, xerror.Wrap(err)
	}

	if account.Balance < amount {
		return false, nil
	}

	err = comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		// 修改订单状态
		updated, err := t.updatePayed(ctx, tx, orderId, bizCategory)
		if err != nil {
			return xerror.Wrap(err)
		}

		if updated {
			dbPayment := db.Payment{
				UserId:      userId,
				StoreId:     storeId,
				BizId:       orderId,
				BizCategory: bizCategory,
				Category:    enum.PaymentCategory_BuyTicket.Value,
				Amount:      amount,
				Status:      enum.PaymentStatus_Payed.Value,
			}

			// 添加支付记录
			err = dbPayment.Insert(ctx, tx)
			if err != nil {
				return xerror.Wrap(err)
			}

		}

		return nil
	})

	return true, xerror.Wrap(err)

}

func (t *PaymentService) Rollback(ctx context.Context, orderIds []string, bizCategory string) error {

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		for _, id := range orderIds {

			_, err := new(db.Payment).Revert(ctx, tx, id, bizCategory)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return xerror.Wrap(err)

}
