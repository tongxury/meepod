package keeperservice

import (
	"context"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/kit/services/util/gind/errorx"
	"gitee.com/meepo/backend/shop/app/shop/service"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/types"
	"github.com/go-pg/pg/v10"
)

type OrderGroupService struct {
	service.OrderGroupService
}

func (t *OrderGroupService) Accept(ctx context.Context, storeId, orderId string) error {

	_, err := new(db.OrderGroup).RequireByIdAndStoreId(ctx, orderId, storeId)
	if err != nil {
		return xerror.Wrap(err)
	}

	err = comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		rejected, err := new(db.OrderGroup).UpdateToAccepted(ctx, tx, orderId)
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

func (t *OrderGroupService) Reject(ctx context.Context, storeId, orderId, reasonId string) error {

	dbOrder, err := new(db.OrderGroup).RequireByIdAndStoreId(ctx, orderId, storeId)
	if err != nil {
		return xerror.Wrap(err)
	}

	err = comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		rejected, err := new(db.OrderGroup).UpdateToRejected(ctx, tx, orderId, reasonId)
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

	t.SendEvent(ctx, "order_group.reject.event", dbOrder)

	return xerror.Wrap(err)
}

func (t *OrderGroupService) Switch(ctx context.Context, storeId, orderId, toStoreId string) error {

	dbOrder, err := new(db.OrderGroup).RequireByIdAndStoreId(ctx, orderId, storeId)
	if err != nil {
		return xerror.Wrap(err)
	}

	if storeId == toStoreId {
		return errorx.ParamErrorf("switch to self")
	}

	err = comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		switched, err := new(db.OrderGroup).UpdateToStoreId(ctx, tx, orderId, toStoreId)
		if err != nil {
			return xerror.Wrap(err)
		}

		if !switched {
			return xerror.Wrapf("cannot switched: %s", orderId)
		}

		return nil
	})

	t.SendEvent(ctx, "order_group.switch.event", dbOrder)

	return xerror.Wrap(err)
}

func (t *OrderGroupService) Ticket(ctx context.Context, storeId, orderId string, images []string) error {

	_, err := new(db.OrderGroup).RequireByIdAndStoreIdOrToStoreId(ctx, orderId, storeId)
	if err != nil {
		return xerror.Wrap(err)
	}

	err = comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		ticketed, err := new(db.OrderGroup).UpdateToTicketed(ctx, tx, orderId, images)
		if err != nil {
			return xerror.Wrap(err)
		}

		if !ticketed {
			return xerror.Wrapf("order cannot ticket: %s", orderId)
		}

		return nil
	})

	return xerror.Wrap(err)
}

func (t *OrderGroupService) GetGroup(ctx context.Context, keeperId, storeId, orderId string) (*types.OrderGroup, error) {

	dbGroup, err := new(db.OrderGroup).RequireByIdAndStoreIdOrToStoreId(ctx, orderId, storeId)
	if err != nil {
		return nil, errorx.ServerError(err)
	}

	groups, err := t.AssembleForKeeper(ctx, storeId, db.OrderGroups{dbGroup})
	if err != nil {
		return nil, errorx.ServerError(err)
	}

	return groups[0], nil
}

func (t *OrderGroupService) ListOrderGroups(ctx context.Context, storeId string, page, size int64) (types.OrderGroups, int64, error) {

	dbGroups, total, err := new(db.OrderGroup).ListKeeperOrders(ctx, "", storeId, nil, nil, page, size)

	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	rsp, err := t.AssembleForKeeper(ctx, storeId, dbGroups)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	return rsp, total, nil

}

func (t *OrderGroupService) RequireByStoreIdAndId(ctx context.Context, storeId, id string) (*types.OrderGroup, error) {

	order, err := new(db.OrderGroup).RequireByIdAndStoreIdOrToStoreId(ctx, id, storeId)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	orders, err := t.AssembleForKeeper(ctx, storeId, db.OrderGroups{order})
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return orders[0], nil
}

func (t *OrderGroupService) AssembleForKeeper(ctx context.Context, keeperStoreId string, dbOrders db.OrderGroups) (types.OrderGroups, error) {
	orders, err := t.Assemble(ctx, "", dbOrders)
	if err != nil {
		return nil, err
	}

	for _, x := range orders {
		if x.ToStore != nil {
			if x.ToStore.Id == keeperStoreId {
				x.Tags = append(x.Tags, types.SwitchInTag)
			} else {
				x.Tags = append(x.Tags, types.SwitchOutTag)
			}

			// 转出后不鞥再出票
			if keeperStoreId != x.ToStore.Id {
				x.Ticketable = false
			}
		}
	}

	return orders, nil
}
