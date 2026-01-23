package keeperservice

import (
	"context"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/kit/components/sdk/helper/mathd"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/kit/services/util/gind/errorx"
	paymentdb "gitee.com/meepo/backend/shop/app/payment/db"
	"gitee.com/meepo/backend/shop/app/shop/service"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/types"
	"github.com/go-pg/pg/v10"
	"time"
)

type OrderService struct {
	service.OrderService
}

//func (t *OrderService) Pay(ctx context.Context, storeId, orderId string) error {
//
//	order, err := new(db.Order).RequireByIdAndStoreId(ctx, orderId, storeId)
//	if err != nil {
//		return xerror.Wrap(err)
//	}
//
//	err = comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {
//
//		// 添加支付记录
//		err = new(db.Payment).InsertPlanPayment(ctx, tx, order.UserId, storeId, orderId, 0, "+", "byKeeper")
//		if err != nil {
//			return xerror.Wrap(err)
//		}
//
//		payed, err := new(db.Order).UpdateToPayed(ctx, tx, orderId, true)
//		if err != nil {
//			return xerror.Wrap(err)
//		}
//
//		if !payed {
//			return xerror.Wrapf("order cannot be payed: %s", orderId)
//		}
//
//		return nil
//	})
//
//	return xerror.Wrap(err)
//}

func (t *OrderService) Switch(ctx context.Context, storeId, orderId, toStoreId string) error {

	order, err := new(db.Order).RequireByIdAndStoreId(ctx, orderId, storeId)
	if err != nil {
		return xerror.Wrap(err)
	}

	if storeId == toStoreId {
		return errorx.ParamErrorf("不能转给自己")
	}

	wallet, err := new(paymentdb.StoreWallet).FindByStoreId(ctx, toStoreId)
	if err != nil {
		return xerror.Wrap(err)
	}

	if wallet == nil || wallet.Balance < mathd.Min(enum.CoStoreRewardMax, order.Amount*enum.CoStoreRewardRate) {
		return errorx.UserMessage("合作方服务费余额不足")
	}

	err = comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		switched, err := new(db.Order).UpdateToStoreId(ctx, tx, orderId, toStoreId)
		if err != nil {
			return xerror.Wrap(err)
		}

		if !switched {
			return xerror.Wrapf("cannot switched: %s", orderId)
		}

		return nil
	})

	return xerror.Wrap(err)
}

func (t *OrderService) Accept(ctx context.Context, storeId, orderId string) error {

	_, err := new(db.Order).RequireByIdAndStoreIdOrToStoreId(ctx, orderId, storeId)
	if err != nil {
		return xerror.Wrap(err)
	}

	err = comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		rejected, err := new(db.Order).UpdateToAccepted(ctx, tx, orderId)
		if err != nil {
			return xerror.Wrap(err)
		}

		if !rejected {
			return xerror.Wrapf("order cannot be accepted: %s", orderId)
		}

		return nil
	})

	return xerror.Wrap(err)
}

func (t *OrderService) Reject(ctx context.Context, storeId, orderId, reasonId string) error {

	dbOrder, err := new(db.Order).RequireByIdAndStoreIdOrToStoreId(ctx, orderId, storeId)
	if err != nil {
		return xerror.Wrap(err)
	}

	err = comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		rejected, err := new(db.Order).UpdateToRejected(ctx, tx, orderId, reasonId)
		if err != nil {
			return xerror.Wrap(err)
		}

		if !rejected {
			return xerror.Wrapf("order cannot reject: %s", orderId)
		}

		//err = new(db.Payment).RevertPayedXX(ctx, tx, dbOrder.UserId, orderId)
		//if err != nil {
		//	return err
		//}

		return nil
	})

	if err != nil {
		return xerror.Wrap(err)
	}

	t.SendEvent(ctx, "order.reject.event", dbOrder)

	return nil
}

func (t *OrderService) Ticket(ctx context.Context, storeId, orderId string, images []string) error {

	dbOrder, err := new(db.Order).RequireByIdAndStoreIdOrToStoreId(ctx, orderId, storeId)
	if err != nil {
		return xerror.Wrap(err)
	}

	err = comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		ticketed, err := new(db.Order).UpdateToTicketed(ctx, tx, orderId, images)
		if err != nil {
			return xerror.Wrap(err)
		}

		if !ticketed {
			return xerror.Wrapf("order cannot ticket: %s", orderId)
		}

		return nil
	})

	if dbOrder.ToStoreId != "" {
		t.SendEvent(ctx, "order.switch.event", dbOrder)
	}

	return xerror.Wrap(err)
}

func (t *OrderService) ListOrders(ctx context.Context, keeperStoreId, itemId string, status []string, from, to *time.Time, page, size int64) (types.Orders, int64, error) {

	dbOrders, total, err := new(db.Order).List(ctx, db.ListOrdersParams{
		StoreIdOrToStoreId: keeperStoreId, ItemId: itemId, MStatus: status,
		CreatedAtFrom: from, CreatedAtTo: to, Page: page, Size: size,
	})
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	orders, err := t.AssembleForKeeper(ctx, keeperStoreId, dbOrders)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	return orders, total, nil
}

func (t *OrderService) RequireByStoreIdAndId(ctx context.Context, storeId, id string) (*types.Order, error) {

	order, err := new(db.Order).RequireByIdAndStoreIdOrToStoreId(ctx, id, storeId)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	orders, err := t.AssembleForKeeper(ctx, storeId, db.Orders{order})
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return orders[0], nil
}

func (t *OrderService) AssembleForKeeper(ctx context.Context, keeperStoreId string, dbOrders db.Orders) (types.Orders, error) {
	orders, err := t.Assemble(ctx, dbOrders)
	if err != nil {
		return nil, err
	}

	_, _, _, userIds, _, _ := dbOrders.Ids()

	countsMap, err := new(db.Order).GetOrderCounts(ctx, userIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	storeUsers, _, err := new(db.StoreUser).List(ctx, db.ListStoreUsersParams{
		StoreId: keeperStoreId, UserIds: userIds,
	})
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	storeUsersMap := storeUsers.AsUserIdMap()

	for _, x := range orders {

		x.Switchable = enum.ShowConStore && helper.InSlice(x.Status.Value, enum.TicketableStatus) && x.ToStore == nil

		if x.ToStore != nil {
			if x.ToStore.Id == keeperStoreId {
				x.Tags = append(x.Tags, types.SwitchInTag)
			} else {
				x.Tags = append(x.Tags, types.SwitchOutTag)
			}

			// 转出后不能再出票
			if keeperStoreId != x.ToStore.Id {
				x.Ticketable = false
			}
		}

		// 新用户标签
		if countsMap[x.User.Id] < 2 {
			x.User.Tags = append(x.User.Tags, types.NewTag)
		}

		if val, found := storeUsersMap[x.User.Id]; found {
			if val.Extra.Remark != "" {
				x.User.Nickname = val.Extra.Remark
			}
		}

	}

	return orders, nil
}

func (t *OrderService) GetOrder(ctx context.Context, keeperId, keeperStoreId, orderId string) (*types.Order, error) {

	dbOrder, err := new(db.Order).RequireByIdAndStoreIdOrToStoreId(ctx, orderId, keeperStoreId)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	orders, err := t.AssembleForKeeper(ctx, keeperStoreId, db.Orders{dbOrder})
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	order := orders[0]

	return order, nil
}
