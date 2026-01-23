package job

import (
	"context"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/shop/app/payment/service"
	coredb "gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
)

func PayRewardAmount() {

	ctx := context.Background()

	rewards, _, err := new(coredb.Reward).List(ctx, coredb.ListRewardsParams{
		MStatus: []string{enum.RewardStatus_Confirmed.Value},
		SyncPay: 1,
		Page:    1, Size: 100,
	})
	if err != nil {
		slf.WithError(err).Errorw("List err")
		return
	}

	if len(rewards) == 0 {
		return
	}

	for _, x := range rewards {

		_, err = new(service.TopupService).AddRewardTopUp(ctx, x.UserId, x.StoreId, x.BizId, x.BizCategory, x.Amount)

		if err != nil {
			slf.WithError(err).Errorw("AddRewardTopUp err")
			return
		}

	}

	ids, _, _ := rewards.Ids()

	_, err = new(coredb.Reward).UpdateToSynced(ctx, ids, "sync_pay")
	if err != nil {
		slf.WithError(err).Errorw("UpdateToSynced err", slf.String("t", "sync_pay"))
		return
	}

}
