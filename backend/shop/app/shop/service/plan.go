package service

import (
	"context"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/kit/services/util/gind/errorx"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/issue/adapter"
	"gitee.com/meepo/backend/shop/core/types"
)

type PlanService struct {
}

func (t *PlanService) Delete(ctx context.Context, userId, planId string) error {

	deleted, err := new(db.Plan).Delete(ctx, planId, userId)
	if err != nil {
		return errorx.ServerError(err)
	}

	if !deleted {
		return errorx.ServiceErrorf("cannot delete: %s", planId)
	}

	return nil
}

func (t *PlanService) ListSavedPlans(ctx context.Context, userId, storeId string, page, size int64) (types.Plans, int64, error) {

	plans, total, err := new(db.Plan).List(ctx, db.ListPlansParams{
		UserId: userId, StoreId: storeId, MStatus: []string{enum.PlanStatus_Saved.Value}, Page: page, Size: size,
	})
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	rsp, err := t.Assemble(ctx, plans)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	return rsp, total, nil
}

//func (t *PlanService) RequirePlanById(ctx context.Context, planId, storeId string) (*model.Plan, error) {
//
//	dbPlan, err := new(db.Plan).RequireById(ctx, planId, storeId)
//	if err != nil {
//		return nil, xerror.Wrap(err)
//	}
//
//	plans, err := t.ToModel(ctx, db.Plans{dbPlan})
//	if err != nil {
//		return nil, xerror.Wrap(err)
//	}
//
//	return plans[0], nil
//}

func (t *PlanService) AddPlan(ctx context.Context, caller, storeId string, params types.PlanForm) (string, error) {

	// 期数
	current, err := new(IssueService).FindCurrentIssue(ctx, params.ItemId)
	if err != nil {
		return "", xerror.Wrap(err)
	}

	if current.Status.Value != enum.IssueStatus_Ongoing.Value {
		return "", errorx.UserMessage("此彩种当前不可购买")
	}

	_, _, amount, err := adapter.Parse(params.ItemId, params.Content)
	if err != nil {
		return "", xerror.Wrap(err)
	}

	dbPlan := db.Plan{
		ItemId:   params.ItemId,
		Issue:    current.Index,
		Content:  params.Content,
		Multiple: params.Multiple,
		Amount:   amount,
		//Volume:   params.Volume,
		StoreId: storeId,
		//Type:     params.Type,
		Status: enum.PlanStatus_Saved.Value,
		UserId: caller,
	}

	_, err = dbPlan.Insert(ctx)
	if err != nil {
		return "", err
	}

	return dbPlan.Id, nil
}

func (t *PlanService) Assemble(ctx context.Context, plans db.Plans) (types.Plans, error) {
	var rsp types.Plans

	_, itemIds, userIds, issueIds := plans.Ids()

	issues, err := new(db.Issue).ListByIds(ctx, issueIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	issuesTap := issues.AsMap()

	items, err := new(db.Item).ListByIds(ctx, itemIds, false)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	itemMap := items.AsMap()

	users, err := new(db.User).ListByIds(ctx, userIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	usersMap := users.AsMap()

	for _, x := range plans {

		issue := issuesTap[x.IssueId()]

		y := types.FromDbPlan(x, itemMap[x.ItemId], usersMap[x.UserId], issue)
		rsp = append(rsp, y)
	}

	return rsp, nil
}
