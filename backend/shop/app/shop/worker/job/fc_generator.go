package job

import (
	"context"
	"gitee.com/meepo/backend/kit/components/sdk/helper/mathd"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/issue/adapter"
	"gitee.com/meepo/backend/shop/core/issue/types"
	"time"
)

// GenerateNext 生成当前时间所属issue, 并补齐其之前的所有缺失的issue(服务不可用)
func GenerateIssue(itemId string, randomDelaySeconds int) {

	time.Sleep(time.Duration(mathd.RandNumber(0, randomDelaySeconds)) * time.Second)

	ctx := context.Background()

	latestIssue, err := new(db.Issue).FindLatest(ctx, itemId)
	if err != nil {
		slf.WithError(err).Errorw("FindLatest err")
		return
	}

	var lastIssue *types.Issue
	if latestIssue != nil {

		if latestIssue.PrizedAt.After(time.Now()) {
			return
		}

		startedAt := latestIssue.PrizedAt.Add(time.Second)

		lastIssue = &types.Issue{
			Index:   latestIssue.Index,
			StartAt: startedAt,
			CloseAt: latestIssue.CloseAt,
			PrizeAt: latestIssue.PrizedAt,
			Prized:  latestIssue.Status == enum.IssueStatus_Prized.Value,
		}
	}

	newIssue, err := adapter.Generate(itemId, lastIssue)
	if newIssue == nil {
		return
	}

	newDbIssue := &db.Issue{
		Id:        itemId + "-" + newIssue.Index,
		ItemId:    itemId,
		Index:     newIssue.Index,
		PrizedAt:  newIssue.PrizeAt,
		StartedAt: newIssue.StartAt,
		CloseAt:   newIssue.CloseAt,
		Status:    enum.IssueStatus_Ongoing.Value,
	}

	err = newDbIssue.Upsert(ctx)
	if err != nil {
		slf.WithError(err).Errorw("newIssue.Upsert err")
		return
	}

}
