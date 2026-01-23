package job

import (
	"context"
	"gitee.com/meepo/backend/kit/components/sdk/helper/mathd"
	"gitee.com/meepo/backend/kit/components/sdk/helper/timed"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/shop/app/payment/service"
	coredb "gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"time"
)

func PayProxyRewards(randomDelaySeconds int) {

	time.Sleep(time.Duration(mathd.RandNumber(0, randomDelaySeconds)) * time.Second)

	ctx := context.Background()

	rewards, _, err := new(coredb.ProxyReward).List(ctx, coredb.ListProxyRewardsParams{
		Status: enum.ProxyRewardStatus_Payed.Value, SyncPayed: 1, Page: 1, Size: 100,
	})
	if err != nil {
		slf.WithError(err).Errorw("List err")
		return
	}

	for _, reward := range rewards {

		_, err := new(service.TopupService).AddProxyRewardTopUp(ctx, reward.StoreId, reward.ProxyUserId, reward.Id, reward.RewardAmount)
		if err != nil {
			slf.WithError(err).Errorw("AddProxyRewardTopUp err")
			return
		}

	}

	_, _, ids := rewards.Ids()

	_, err = new(coredb.ProxyReward).UpdateToSynced(ctx, ids, "sync_payed")
	if err != nil {
		slf.WithError(err).Errorw("UpdateToSynced err")
		return
	}
}

func ConfirmProxyRewards(randomDelaySeconds int) {

	time.Sleep(time.Duration(mathd.RandNumber(0, randomDelaySeconds)) * time.Second)

	ctx := context.Background()
	y, m, _ := time.Now().In(timed.LocAsiaShanghai).Date()
	if m == 1 {
		y -= 1
		m = 12
	}

	lastMonth := time.Date(y, m, 0, 0, 0, 0, 0, timed.LocAsiaShanghai).Format("2006-01")

	err := new(coredb.ProxyReward).UpdateConfirmed(ctx, lastMonth)
	if err != nil {
		slf.WithError(err).Errorw("UpdateConfirmed err")
		return
	}
}

func CalculateProxyRewards(randomDelaySeconds int) {

	time.Sleep(time.Duration(mathd.RandNumber(0, randomDelaySeconds)) * time.Second)

	ctx := context.Background()
	// 本月所有的已出票的订单

	month := time.Now().In(timed.LocAsiaShanghai).Format("2006-01")

	rewards, err := new(coredb.ProxyReward).AggregateMonthRewards(ctx, month)
	if err != nil {
		slf.WithError(err).Errorw("AggregateMonthRewards err")
		return
	}

	// 补充佣金
	proxyIds, _, _ := rewards.Ids()

	proxies, _, err := new(coredb.Proxy).List(ctx, coredb.ListProxyParams{
		Ids: proxyIds,
	})
	if err != nil {
		return
	}

	proxyMap := proxies.AsMap()

	for _, x := range rewards {

		x.Month = month
		x.Status = enum.ProxyRewardStatus_Pending.Value

		proxy := proxyMap[x.ProxyId]
		x.ProxyUserId = proxy.UserId
		x.StoreId = proxy.StoreId
		x.RewardRate = proxy.RewardRate
		x.RewardAmount = mathd.ToFixed4(x.OrderAmount * proxy.RewardRate)
	}

	for _, reward := range rewards {

		_, err := reward.Upsert(ctx)
		if err != nil {
			slf.WithError(err).Errorw("Upsert err")
		}
	}

}
