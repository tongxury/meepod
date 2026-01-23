package main

import (
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/comp/comps"
	"gitee.com/meepo/backend/kit/components/sdk/redisstream"
	"gitee.com/meepo/backend/kit/components/sdk/runner"
	"gitee.com/meepo/backend/shop/app/payment/worker/event"
	"gitee.com/meepo/backend/shop/app/payment/worker/job"
	"github.com/robfig/cron/v3"
)

func main() {
	args := comp.Flags().Logger().Redis().Postgres().Server().Alipay().Xinsh().
		Parse()

	comp.SDK().Preparing().Logger(args.Log).Postgres(args.RepoPostgres).Redis(args.RepoRedis).
		Alipay(args.AlipayOptions).Xinsh(args.XinshConfig)

	workerName := comp.Flags().ServerOptions.Name

	run := runner.New()
	run.
		RedisStreamService(redisstream.NewRedisStream(comp.SDK().Redis(),
			"order.switch.event.group", workerName, "order.switch.event", event.ConsumeSwitchEvent,
		)).
		RedisStreamService(redisstream.NewRedisStream(comp.SDK().Redis(),
			"order.cancel.event.group", workerName, "order.cancel.event", event.ConsumeCancelEvent,
		)).
		RedisStreamService(redisstream.NewRedisStream(comp.SDK().Redis(),
			"order.reject.event.group", workerName, "order.reject.event", event.ConsumeRejectEvent,
		)).
		RedisStreamService(redisstream.NewRedisStream(comp.SDK().Redis(),
			"order_group.reject.event.group", workerName, "order_group.reject.event", event.ConsumeOrderGroupRejectEvent,
		)).
		RedisStreamService(redisstream.NewRedisStream(comp.SDK().Redis(),
			"order_group.switch.event.group", workerName, "order_group.switch.event", event.ConsumeOrderGroupSwitchEvent,
		)).
		RedisStreamService(redisstream.NewRedisStream(comp.SDK().Redis(),
			"reward.pay.event.group", workerName, "reward.pay.event", event.ConsumePayRewardEvent,
		)).
		RedisStreamService(redisstream.NewRedisStream(comp.SDK().Redis(),
			"store.created.event.group", workerName, "store.created.event", event.ConsumeStoreCreatedEvent,
		))

	cr := cron.New(cron.WithSeconds())
	_, _ = cr.AddJob("@every 5m", comps.NewJobWrapper(job.XinshStoreApplyState))
	_, _ = cr.AddJob("@every 5s", comps.NewJobWrapper(job.XinshPayResult))

	_, _ = cr.AddJob("@every 5s", comps.NewJobWrapper(job.CleanTopups))
	// 店铺合作
	_, _ = cr.AddJob("@every 1h", comps.NewJobWrapper(job.PaySwitchAmount))
	// 订单
	_, _ = cr.AddJob("@every 1h", comps.NewJobWrapper(job.RollbackOrderPayment))
	_, _ = cr.AddJob("@every 1h", comps.NewJobWrapper(job.RollbackOrderGroupPayment))
	// 分账
	//_, _ = cr.AddJob("@every 2s", comps.NewJobWrapper(job.Settle))

	// 中奖
	_, _ = cr.AddJob("@every 1h", comps.NewJobWrapper(job.PayRewardAmount))

	run.CronService(cr)

	run.Run()
}
