package service

import (
	"context"
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/types"
	"time"
)

type ItemService struct {
}

func (t *ItemService) ListMetaByIds(ctx context.Context, ids []string, allIfEmptyIds bool) (types.Items, error) {
	dbItems, err := new(db.Item).ListByIds(ctx, ids, allIfEmptyIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	items, err := t.Assemble(ctx, dbItems)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return items, nil
}

func (t *ItemService) Assemble(ctx context.Context, items db.Items) (types.Items, error) {

	var rsp types.Items
	for _, x := range items {

		y := types.FromDbItem(x)
		rsp = append(rsp, y)
	}

	return rsp, nil
}

func (t *ItemService) ListAllItems(ctx context.Context) (types.ItemStates, error) {

	var err error

	items, err := t.ListMetaByIds(ctx, nil, true)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	var states types.ItemStates

	for _, x := range items {
		y := types.ItemState{
			Item: *x,
		}

		issue, err := new(IssueService).FindCurrentIssue(ctx, x.Id)
		if err != nil {
			slf.WithError(err).Errorw("FindCurrentIssue err")
		}

		y.LatestIssue = issue

		y.Disabled = true
		y.Status = enum.ItemStatus_Unable

		states = append(states, &y)
	}

	return states, nil
}

func (t *ItemService) ListStoreItems(ctx context.Context, storeId string) (types.ItemStates, error) {

	store, err := new(StoreService).FindStoreById(ctx, storeId)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	itemIds := store.SelectedItemIds

	items, err := t.ListMetaByIds(ctx, itemIds, false)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	var states types.ItemStates

	for _, x := range items {
		y := types.ItemState{
			Item: *x,
		}

		issue, err := new(IssueService).FindCurrentIssue(ctx, x.Id)
		if err != nil {
			slf.WithError(err).Errorw("FindCurrentIssue err")
		}

		y.LatestIssue = issue

		y.Status, y.Extra, y.Disabled = t.mapIssueStatusToItemStatus(issue)

		if store.Status.Value != enum.StoreStatus_Confirmed.Value {
			y.Status = enum.ItemStatus_Invalid
			y.Disabled = true
		}

		states = append(states, &y)
	}

	return states, nil
}

func (t *ItemService) mapIssueStatusToItemStatus(issue *types.Issue) (enum.Enum, types.Extra, bool) {

	if issue == nil {
		return enum.ItemStatus_Closed, types.Extra{}, true
	}

	now := time.Now()

	countdown := types.Extra{Type: "countdown", Value: issue.CloseTimeLeft}
	issueIndex := types.Extra{Type: "text", Value: issue.Index}

	// 截止
	if issue.CloseAtTime.Before(now) {
		return enum.ItemStatus_Closed, issueIndex, true
	}

	// 已拉取到开奖结果 但未生成下一期
	if issue.Status.Value == enum.IssueStatus_Prized.Value {
		return enum.ItemStatus_Unstart, issueIndex, true
	}

	switch issue.Item.Id {
	case enum.ItemId_zjc:
		if len(issue.Extra) == 0 {
			return enum.ItemStatus_NoData, types.Extra{}, true
		}
		return enum.ItemStatus_Ongoing, types.Extra{Type: "text", Value: fmt.Sprintf("%d场", len(issue.Extra))}, false

	case enum.ItemId_rx9, enum.ItemId_sfc:

		if len(issue.Extra) == 0 {
			return enum.ItemStatus_NoData, issueIndex, true
		}
		return enum.ItemStatus_Ongoing, countdown, false
	}

	return enum.ItemStatus_Ongoing, countdown, false

}
