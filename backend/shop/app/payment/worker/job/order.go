package job

import (
	"context"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/shop/app/payment/service"
	keeperservice "gitee.com/meepo/backend/shop/app/payment/service/keeper"
	coredb "gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
)

func PaySwitchAmount() {

	ctx := context.Background()

	orders, _, err := new(coredb.Order).List(ctx, coredb.ListOrdersParams{
		MStatus:           enum.CoStorePayableStatus,
		ToStoreIdNotEmpty: true,
		SyncSwitch:        1,
		Page:              1, Size: 100,
	})
	if err != nil {
		slf.WithError(err).Errorw("List err")
		return
	}

	if len(orders) == 0 {
		return
	}

	for _, order := range orders {

		err = new(keeperservice.CoStorePaymentService).PaySwitchOrder(ctx, order.StoreId, order.ToStoreId, enum.BizCategory_Order.Value, order.Id, order.Amount)

		if err != nil {
			slf.WithError(err).Errorw("PaySwitchOrder err")
			return
		}

	}

	ids, _, _, _, _, _ := orders.Ids()

	_, err = new(coredb.Order).UpdateToSynced(ctx, ids, "sync_switch")
	if err != nil {
		slf.WithError(err).Errorw("UpdateToSynced err", slf.String("t", "sync_switch"))
		return
	}
}

func RollbackOrderPayment() {
	ctx := context.Background()

	orders, _, err := new(coredb.Order).List(ctx, coredb.ListOrdersParams{
		MStatus:      []string{enum.OrderStatus_Canceled.Value, enum.OrderStatus_Rejected.Value, enum.OrderStatus_Timeout.Value},
		SyncRollback: 1,
		Page:         1, Size: 100,
	})
	if err != nil {
		slf.WithError(err).Errorw("ListByStatus err")
		return
	}

	if len(orders) == 0 {
		return
	}

	ids, _, _, _, _, _ := orders.Ids()

	err = new(service.PaymentService).Rollback(ctx, ids, enum.BizCategory_Order.Value)
	if err != nil {
		slf.WithError(err).Errorw("Rollback err")
		return
	}

	_, err = new(coredb.Order).UpdateToSynced(ctx, ids, "sync_rollback")
	if err != nil {
		slf.WithError(err).Errorw("UpdateToSynced err")
		return
	}

}
