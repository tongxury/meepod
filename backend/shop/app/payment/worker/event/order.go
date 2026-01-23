package event

import (
	"context"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/app/payment/service"
	keeperservice "gitee.com/meepo/backend/shop/app/payment/service/keeper"
	coredb "gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	redisV9 "github.com/redis/go-redis/v9"
)

func ConsumeSwitchEvent(ctx context.Context, messages []redisV9.XMessage) error {

	slf.Debugw("SwitchEvent ", slf.Reflect("messages", messages))

	for _, message := range messages {

		id := conv.String(message.Values["Id"])
		storeId := conv.String(message.Values["StoreId"])
		toStoreId := conv.String(message.Values["ToStoreId"])
		amount := conv.Float64(message.Values["Amount"])

		if toStoreId == "" {
			continue
		}

		err := new(keeperservice.CoStorePaymentService).PaySwitchOrder(ctx, storeId, toStoreId, enum.BizCategory_Order.Value, id, amount)
		if err != nil {
			slf.WithError(err).Errorw("PaySwitchOrder err")
			// err内部处理 免得阻塞
			continue
		}

		_, err = new(coredb.Order).UpdateToSynced(ctx, []string{id}, "sync_switch")
		if err != nil {
			slf.WithError(err).Errorw("UpdateToSynced err", slf.String("t", "sync_switch"))
			continue
		}
	}

	return nil
}

func ConsumeCancelEvent(ctx context.Context, messages []redisV9.XMessage) error {

	slf.Debugw("CancelEvent ", slf.Reflect("messages", messages))

	for _, message := range messages {

		id := conv.String(message.Values["Id"])

		if err := rollbackOrder(ctx, id); err != nil {
			slf.WithError(err).Errorw("rollback err", slf.String("id", id))
			continue
		}
	}

	return nil
}

func ConsumeRejectEvent(ctx context.Context, messages []redisV9.XMessage) error {

	slf.Debugw("RejectEvent ", slf.Reflect("messages", messages))

	for _, message := range messages {

		id := conv.String(message.Values["Id"])

		if err := rollbackOrder(ctx, id); err != nil {
			slf.WithError(err).Errorw("rollback err", slf.String("id", id))
			continue
		}
	}

	return nil
}

func rollbackOrder(ctx context.Context, id string) error {

	err := new(service.PaymentService).Rollback(ctx, []string{id}, enum.BizCategory_Order.Value)
	if err != nil {
		return xerror.Wrap(err)
	}

	_, err = new(coredb.Order).UpdateToSynced(ctx, []string{id}, "sync_rollback")
	if err != nil {
		return xerror.Wrap(err)
	}
	return nil
}
