package main

import (
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/comp/comps"
	"gitee.com/meepo/backend/kit/components/sdk/runner"
	"gitee.com/meepo/backend/shop/app/shop/worker/job"
	"gitee.com/meepo/backend/shop/core/enum"
	"github.com/robfig/cron/v3"
)

func main() {
	args := comp.Flags().Logger().Redis().Postgres().Server().
		CustomStr("prize.cron", "@every 10s", "").
		CustomStr("normal.cron", "@every 5s", "").
		CustomStr("proxy.pool.url", "http://localhost:5010", "").
		Parse()

	comp.SDK().Preparing().Logger(args.Log).Postgres(args.RepoPostgres).Redis(args.RepoRedis)

	run := runner.New()

	cr := cron.New(cron.WithSeconds())

	normalCron := comp.Flags().GetStr("normal.cron")
	prizeCron := comp.Flags().GetStr("prize.cron")

	// 推广
	_, _ = cr.AddJob(normalCron, comps.NewJobWrapper(func() { job.CalculateProxyRewards(60) }))
	_, _ = cr.AddJob(normalCron, comps.NewJobWrapper(func() { job.ConfirmProxyRewards(60) }))
	_, _ = cr.AddJob(normalCron, comps.NewJobWrapper(func() { job.PayProxyRewards(60) }))

	{
		//  ==================== 普通彩种 ====================
		// 生成issue
		_, _ = cr.AddJob(normalCron, comps.NewJobWrapper(func() { job.GenerateIssue(enum.ItemId_dlt, 60) }))
		_, _ = cr.AddJob(normalCron, comps.NewJobWrapper(func() { job.GenerateIssue(enum.ItemId_ssq, 60) }))
		_, _ = cr.AddJob(normalCron, comps.NewJobWrapper(func() { job.GenerateIssue(enum.ItemId_f3d, 60) }))
		_, _ = cr.AddJob(normalCron, comps.NewJobWrapper(func() { job.GenerateIssue(enum.ItemId_x7c, 60) }))
		_, _ = cr.AddJob(normalCron, comps.NewJobWrapper(func() { job.GenerateIssue(enum.ItemId_pl3, 60) }))
		_, _ = cr.AddJob(normalCron, comps.NewJobWrapper(func() { job.GenerateIssue(enum.ItemId_pl5, 60) }))

		// 获取开奖结果
		_, _ = cr.AddJob(normalCron, comps.NewJobWrapper(func() { job.FindIssueTarget(enum.ItemId_dlt) }))
		_, _ = cr.AddJob(normalCron, comps.NewJobWrapper(func() { job.FindIssueTarget(enum.ItemId_ssq) }))
		_, _ = cr.AddJob(normalCron, comps.NewJobWrapper(func() { job.FindIssueTarget(enum.ItemId_f3d) }))
		_, _ = cr.AddJob(normalCron, comps.NewJobWrapper(func() { job.FindIssueTarget(enum.ItemId_x7c) }))
		_, _ = cr.AddJob(normalCron, comps.NewJobWrapper(func() { job.FindIssueTarget(enum.ItemId_pl3) }))
		_, _ = cr.AddJob(normalCron, comps.NewJobWrapper(func() { job.FindIssueTarget(enum.ItemId_pl5) }))

		// 对比开奖结果
		_, _ = cr.AddJob(prizeCron, comps.NewJobWrapper(func() { job.PrizeUnPrizedOrderV2(enum.ItemId_dlt, 5) }))
		_, _ = cr.AddJob(prizeCron, comps.NewJobWrapper(func() { job.PrizeUnPrizedOrderV2(enum.ItemId_ssq, 5) }))
		_, _ = cr.AddJob(prizeCron, comps.NewJobWrapper(func() { job.PrizeUnPrizedOrderV2(enum.ItemId_f3d, 5) }))
		_, _ = cr.AddJob(prizeCron, comps.NewJobWrapper(func() { job.PrizeUnPrizedOrderV2(enum.ItemId_x7c, 5) }))
		_, _ = cr.AddJob(prizeCron, comps.NewJobWrapper(func() { job.PrizeUnPrizedOrderV2(enum.ItemId_pl3, 5) }))
		_, _ = cr.AddJob(prizeCron, comps.NewJobWrapper(func() { job.PrizeUnPrizedOrderV2(enum.ItemId_pl5, 5) }))
	}

	{
		//  ==================== 足彩 ====================
		// 数据准备
		_, _ = cr.AddJob(normalCron, comps.NewJobWrapper(job.FetchZjcMatches))
		_, _ = cr.AddJob(normalCron, comps.NewJobWrapper(job.FetchZ14Matches))

		// 生成issue
		_, _ = cr.AddJob(normalCron, comps.NewJobWrapper(job.GenerateZ14))

		// 获取比赛结果
		_, _ = cr.AddJob(normalCron, comps.NewJobWrapper(job.FetchZjcMatchesResult))
		_, _ = cr.AddJob(normalCron, comps.NewJobWrapper(job.FetchZ14MatchesResult))

		// 对比开奖结果
		_, _ = cr.AddJob(prizeCron, comps.NewJobWrapper(func() { job.PrizeUnPrizedZcOrderV2(enum.ItemId_rx9, 5) }))
		_, _ = cr.AddJob(prizeCron, comps.NewJobWrapper(func() { job.PrizeUnPrizedZcOrderV2(enum.ItemId_sfc, 5) }))
	}

	run.CronService(cr)

	run.Run()
}
