package service

import (
	"context"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/kit/services/util/gind/errorx"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/types"
	"time"
)

type IssueService struct{}

func (t *IssueService) FindCurrentIssueIndex(ctx context.Context, itemId string) (string, error) {
	now := time.Now()

	dbIssue, err := new(db.Issue).FindByTimestamp(ctx, itemId, now)
	if err != nil {
		return "", xerror.Wrap(err)
	}

	if dbIssue == nil {
		return "", nil
	}

	return dbIssue.Index, nil
}

func (t *IssueService) FindLatestIssueIndex(ctx context.Context, itemId string) (string, error) {

	dbIssue, err := new(db.Issue).FindLatest(ctx, itemId)
	if err != nil {
		return "", xerror.Wrap(err)
	}

	if dbIssue == nil {
		return "", nil
	}

	return dbIssue.Index, nil
}

func (t *IssueService) FindCurrentIssue(ctx context.Context, itemId string) (*types.Issue, error) {

	now := time.Now()

	dbIssue, err := new(db.Issue).FindByTimestamp(ctx, itemId, now)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	if dbIssue == nil {
		return nil, nil
	}

	issues, err := t.ToModels(ctx, db.Issues{dbIssue})
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	issue := issues[0]

	if err := t.repair(ctx, issue); err != nil {
		return nil, xerror.Wrap(err)
	}

	return issue, nil
}

func (t *IssueService) repair(ctx context.Context, issue *types.Issue) error {
	// 补充场次数据
	switch issue.Item.Id {
	case enum.ItemId_rx9, enum.ItemId_sfc:

		matches, err := new(MatchService).ListByCategoryAndIssue(ctx, enum.MatchCategory_Z14.Value, issue.Index)
		if err != nil {
			return xerror.Wrap(err)
		}

		issue.Extra = conv.AnySlice(matches)

	case enum.ItemId_zjc:

		issue.Index = ""

		// 开赛前4小时 停止销售
		ts := time.Now().Add(4 * time.Hour)
		matches, err := new(MatchService).ListMatchesByCategoryAndStartTs(ctx, enum.MatchCategory_Zjc.Value, ts)
		if err != nil {
			return xerror.Wrap(err)
		}
		issue.Extra = conv.AnySlice(matches)

	default:
	}

	return nil
}

//func (t *IssueService) LatestIssue(ctx context.Context, itemId string) (*model.Issue, error) {
//
//	dbIssue, err := new(db.Issue).FindLatest(ctx, itemId)
//	if err != nil {
//		return nil, xerror.Wrap(err)
//	}
//
//	if dbIssue == nil {
//		return nil, nil
//	}
//
//	issues, err := t.ToModels(ctx, db.Issues{dbIssue})
//	if err != nil {
//		return nil, xerror.Wrap(err)
//	}
//
//	issue := issues[0]
//
//	//now := time.Now()
//	// issue状态
//	//if dbIssue.PrizedAt.Before(now) {
//	//	issue.MStatus = enum.IssueStatus_Prized
//	//	//issue.MStatus.Desc = "未开始"
//	//} else if dbIssue.CloseAt.Before(now) {
//	//	issue.MStatus = enum.IssueStatus_Closed
//	//	//issue.MStatus.Desc = "未开始"
//	//}
//
//	if err := t.repair(ctx, issue); err != nil {
//		return nil, xerror.Wrap(err)
//	}
//
//	return issue, nil
//}

//func (t *IssueService) findLatest(ctx context.Context, itemId string) (*model.Issue, error) {
//	currentIssue, err := new(db.Issue).FindLatest(ctx, itemId)
//	if err != nil {
//		return nil, xerror.Wrap(err)
//	}
//
//	if currentIssue == nil {
//		return nil, nil
//	}
//
//	issues, err := t.ToModels(ctx, db.Issues{currentIssue})
//	if err != nil {
//		return nil, xerror.Wrap(err)
//	}
//
//	return issues[0], nil
//}

func (t *IssueService) ToModels(ctx context.Context, issues db.Issues) (types.Issues, error) {

	itemIds, _ := issues.Ids()

	items, err := new(db.Item).ListByIds(ctx, itemIds, false)
	if err != nil {
		return nil, errorx.ServerError(err)
	}
	itemsMap := items.AsMap()

	var rsp types.Issues

	for _, x := range issues {
		y := types.FromDbIssue(x, itemsMap[x.ItemId])
		rsp = append(rsp, y)
	}

	return rsp, nil

}
