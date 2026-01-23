package adminserivce

import (
	"context"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/app/shop/service"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/types"
)

type OrderService struct {
	service.OrderService
}

func (t *OrderService) ListOrders(ctx context.Context, id, storeId, status string, page, size int64) (types.Orders, int64, error) {

	var stats []string
	if status != "" {
		stats = []string{}
	}

	dbOrders, total, err := new(db.Order).List(ctx, db.ListOrdersParams{
		Id: id, StoreId: storeId, MStatus: stats, Page: page, Size: size,
	})

	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	orders, err := t.Assemble(ctx, dbOrders)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	return orders, total, nil
}
