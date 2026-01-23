package job

import (
	"context"
	"fmt"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/helper/mathd"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/issue/adapter"
	"github.com/go-pg/pg/v10"
	"time"
)

func FindIssueTarget(itemId string) {

	ctx := context.Background()

	issue, err := new(db.Issue).FindLastUnPrized(ctx, itemId)
	if err != nil {
		slf.WithError(err).Errorw("FindLastUnPrized err")
		return
	}

	if issue == nil {
		return
	}

	// 还未开奖
	if issue.PrizedAt.After(time.Now()) {
		return
	}

	result, err := adapter.FetchTarget(itemId, issue.Index)
	if err != nil {
		slf.WithError(err).Errorw("FetchLatestResult err")
		return
	}

	if result == nil || result.Result == "" {
		return
	}

	err = comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {
		_, err := new(db.Issue).UpdatePrizeResult(ctx, tx, itemId+"-"+result.Issue,
			result.Result, nil, result.PrizeAt)
		return err
	})

	if err != nil {
		slf.WithError(err).Errorw("UpdatePrizeResult err")
	}

}

func PrizeUnPrizedOrderV2(itemId string, randomDelaySeconds int) {

	time.Sleep(time.Duration(mathd.RandNumber(0, randomDelaySeconds)) * time.Second)

	ctx := context.Background()

	// 未开奖的Issues
	issues, _, err := new(db.Issue).List(ctx, db.ListIssuesParams{
		ItemIds: []string{itemId}, MExcludeStatus: []string{enum.IssueStatus_Prized.Value}})
	if err != nil {
		slf.WithError(err).Errorw("List Issues err")
		return
	}

	_, issueIndexes := issues.Ids()

	// 普通order
	prizeOrders(ctx, issueIndexes, itemId)
	// 合买order
	prizeOrderGroups(ctx, issueIndexes, itemId)
}

func prizeOrders(ctx context.Context, excludeIssueIndexes []string, itemId string) {

	// 普通order
	orders, _, err := new(db.Order).List(ctx, db.ListOrdersParams{
		ItemId:              itemId,
		ExcludeIssueIndexes: excludeIssueIndexes,
		MStatus:             []string{enum.OrderStatus_Ticketed.Value},
		Size:                100,
	})
	if err != nil {
		slf.WithError(err).Errorw("List Orders err")
		return
	}

	if len(orders) == 0 {
		return
	}

	if err := doPrizeOrders(ctx, orders); err != nil {
		slf.WithError(err).Errorw("prizeOrders err")
	}

}

func doPrizeOrders(ctx context.Context, orders db.Orders) error {
	_, planIds, _, _, _, issueIds := orders.Ids()

	plans, _, err := new(db.Plan).List(ctx, db.ListPlansParams{
		Ids: planIds,
	})
	if err != nil {
		return xerror.Wrap(err)
	}
	plansMap := plans.AsMap()

	issues, _, err := new(db.Issue).List(ctx, db.ListIssuesParams{Ids: issueIds})
	if err != nil {
		return xerror.Wrap(err)
	}
	issuesMap := issues.AsMap()

	var successOrderIds []string
	var rewards db.Rewards

	for _, x := range orders {

		plan := plansMap[x.PlanId]
		issue := issuesMap[x.IssueId()]

		// 查到开奖结果
		records, err := adapter.Prize(plan.ItemId, plan.Content, issue.Result)
		if err != nil {
			slf.WithError(err).Errorw("Prize err", slf.Reflect("plan", plan))
			continue
		}

		if len(records) > 0 {
			dbReward := db.Reward{
				BizId:       x.Id,
				BizCategory: enum.BizCategory_Order.Value,
				UserId:      x.UserId,
				StoreId:     x.StoreId,
				Status:      enum.RewardStatus_Confirmed.Value,
			}

			totalCount, totalAmount := records.Total()
			totalAmount = totalAmount * float64(plan.Multiple)
			dbReward.Amount = totalAmount

			dbReward.Extra = db.RewardExtra{
				Summary:     fmt.Sprintf("%d注%.2f元(%d倍)", totalCount, totalAmount, plan.Multiple),
				TotalCount:  int64(totalCount),
				TotalAmount: totalAmount,
				Multiple:    plan.Multiple,
			}

			rewards = append(rewards, &dbReward)

		}

		successOrderIds = append(successOrderIds, x.Id)

	}

	err = comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {
		_, err := new(db.Order).UpdateToPrized(ctx, tx, successOrderIds)
		if err != nil {
			return err
		}

		_, err = rewards.InsertBatch(ctx, tx)
		if err != nil {
			return err
		}

		for _, x := range rewards {
			_, err := new(db.Order).UpdateExtra(ctx, tx, x.BizId, "summary", x.Extra.Summary)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return nil

}

func prizeOrderGroups(ctx context.Context, excludeIssueIndexes []string, itemId string) {
	// 合买order
	orders, _, err := new(db.OrderGroup).List(ctx, db.ListOrderGroupsParams{
		ItemId:              itemId,
		ExcludeIssueIndexes: excludeIssueIndexes,
		MStatus:             []string{enum.OrderStatus_Ticketed.Value},
		Size:                100,
	})
	if err != nil {
		slf.WithError(err).Errorw("List Orders err")
		return
	}

	if len(orders) == 0 {
		return
	}

	if err := doPrizeOrderGroups(ctx, orders); err != nil {
		slf.WithError(err).Errorw("doPrizeOrderGroups err")
	}

}

func doPrizeOrderGroups(ctx context.Context, orders db.OrderGroups) error {
	_, planIds, _, _, issueIds := orders.Ids()

	plans, _, err := new(db.Plan).List(ctx, db.ListPlansParams{
		Ids: planIds,
	})
	if err != nil {
		return xerror.Wrap(err)
	}
	plansMap := plans.AsMap()

	issues, _, err := new(db.Issue).List(ctx, db.ListIssuesParams{Ids: issueIds})
	if err != nil {
		return xerror.Wrap(err)
	}
	issuesMap := issues.AsMap()

	var successOrderIds []string
	var rewards db.Rewards
	orderRewardSummaries := make(map[string]string, len(orders))

	for _, x := range orders {

		plan := plansMap[x.PlanId]
		issue := issuesMap[x.IssueId()]

		records, err := adapter.Prize(plan.ItemId, plan.Content, issue.Result)
		if err != nil {
			slf.WithError(err).Errorw("Prize err", slf.Reflect("plan", plan))
			continue
		}

		if len(records) > 0 {

			shares, err := new(db.OrderGroupShare).FindByGroupIdAndStatus(ctx, x.Id, []string{enum.OrderGroupShareStatus_Payed.Value})
			if err != nil {
				slf.WithError(err).Errorw("FindByGroupIdAndStatus err")
				continue
			}

			totalCount, totalAmount := records.Total()
			totalAmount = totalAmount * float64(plan.Multiple)

			for _, share := range shares {

				totalVolume := x.Volume
				volume := share.Volume

				amount := mathd.ToFixed4((conv.Float64(volume) / conv.Float64(totalVolume)) * totalAmount)

				dbReward := db.Reward{
					BizId:       share.Id,
					BizCategory: enum.BizCategory_GroupShare.Value,
					UserId:      share.UserId,
					StoreId:     x.StoreId,
					Amount:      amount,
					Status:      enum.RewardStatus_Confirmed.Value,
					Extra: db.RewardExtra{
						Summary:      fmt.Sprintf("%d注%.2f元(%d倍,%d/%d)", totalCount, amount, plan.Multiple, volume, totalVolume),
						TotalCount:   int64(totalCount),
						TotalAmount:  totalAmount,
						Multiple:     plan.Multiple,
						OrderGroupId: x.Id,
						TotalVolume:  totalVolume,
						Volume:       volume,
					},
				}

				rewards = append(rewards, &dbReward)

			}

			orderRewardSummaries[x.Id] = fmt.Sprintf("%d注%.2f元(%d倍)", totalCount, totalAmount, plan.Multiple)
		}

		successOrderIds = append(successOrderIds, x.Id)

	}

	err = comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {
		_, err := new(db.OrderGroup).UpdateToPrized(ctx, tx, successOrderIds)
		if err != nil {
			return err
		}

		_, err = rewards.InsertBatch(ctx, tx)
		if err != nil {
			return err
		}

		for orderId, summary := range orderRewardSummaries {
			_, err := new(db.OrderGroup).UpdateExtra(ctx, tx, orderId, "summary", summary)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return nil

}
