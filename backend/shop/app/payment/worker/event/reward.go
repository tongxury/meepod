package event

import (
	"context"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/shop/app/payment/service"
	coredb "gitee.com/meepo/backend/shop/core/db"
	redisV9 "github.com/redis/go-redis/v9"
)

func ConsumePayRewardEvent(ctx context.Context, messages []redisV9.XMessage) error {

	slf.Debugw("RewardPayEvent ", slf.Reflect("messages", messages))

	for _, message := range messages {

		id := conv.String(message.Values["Id"])
		storeId := conv.String(message.Values["StoreId"])
		userId := conv.String(message.Values["UserId"])
		bizId := conv.String(message.Values["BizId"])
		bizCategory := conv.String(message.Values["BizCategory"])
		amount := conv.Float64(message.Values["Amount"])

		_, err := new(service.TopupService).AddRewardTopUp(ctx, userId, storeId, bizId, bizCategory, amount)
		if err != nil {
			slf.WithError(err).Errorw("AddRewardTopUp err")
			// err内部处理 免得阻塞
			continue
		}

		_, err = new(coredb.Reward).UpdateToSynced(ctx, []string{id}, "sync_pay")
		if err != nil {
			slf.WithError(err).Errorw("UpdateToSynced err", slf.String("t", "sync_pay"))
			continue
		}
	}

	return nil
}
