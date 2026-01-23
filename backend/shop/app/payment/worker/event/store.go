package event

import (
	"context"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/shop/app/payment/service"
	redisV9 "github.com/redis/go-redis/v9"
)

func ConsumeStoreCreatedEvent(ctx context.Context, messages []redisV9.XMessage) error {

	slf.Debugw("ConsumeStoreCreatedEvent ", slf.Reflect("messages", messages))

	for _, message := range messages {
		storeId := conv.String(message.Values["storeId"])

		_, err := new(service.StoreService).AddStoreV2(ctx, storeId)
		if err != nil {
			slf.WithError(err).Errorw("AddStoreV2 err", slf.String("storeId", storeId))
			return err
		}
	}

	return nil
}
