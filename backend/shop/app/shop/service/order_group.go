package service

import (
	"context"
	"fmt"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/kit/components/sdk/helper/timed"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/kit/services/util/gind/errorx"
	"gitee.com/meepo/backend/kit/services/util/oss"
	"gitee.com/meepo/backend/shop/app/payment/service"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/types"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-pg/pg/v10"
	redisV9 "github.com/redis/go-redis/v9"
	"math"
	"strings"
)

type OrderGroupService struct{}

//func (t *OrderGroupService) CreateGroupAndFirstGroupOrder(ctx context.Context, userId, storeId, planId string, totalVolume, volume, floor int64, remark string, rewardRate float64) (string, error) {
//
//	groupId, err := new(OrderGroupService).AddGroup(ctx, userId, storeId, planId, totalVolume, floor, remark, rewardRate)
//	if err != nil {
//		return "", errorx.ServerError(err)
//	}
//
//	return t.JoinGroup(ctx, userId, storeId, groupId, volume)
//}

func (t *OrderGroupService) JoinGroup(ctx context.Context, userId, storeId, groupId string, volume int64) (string, error) {

	orderGroup, err := new(db.OrderGroup).RequireById(ctx, groupId)
	if err != nil {
		return "", xerror.Wrap(err)
	}

	// group 状态
	if orderGroup.Status != enum.OrderGroupStatus_Submitted.Value {
		return "", errorx.UserMessage("合买单状态异常")
	}

	// 当前剩余
	if orderGroup.VolumeLeft() < volume {
		return "", errorx.ParamErrorf("not enough volume: left %d", orderGroup.VolumeLeft())
	}

	//if orderGroup.Floor > volume {
	//	return "", errorx.ParamErrorf("not enough volume: floor %d", orderGroup.Floor)
	//}

	plan, err := new(db.Plan).RequireById(ctx, orderGroup.PlanId)
	if err != nil {
		return "", errorx.ParamError(err)
	}

	amount := math.Ceil((plan.RealAmount() / conv.Float64(orderGroup.Volume)) * conv.Float64(volume))

	dbOrder := &db.OrderGroupShare{
		Volume:  volume,
		Amount:  amount,
		StoreId: storeId,
		GroupId: groupId,
		UserId:  userId,
		Status:  enum.OrderGroupShareStatus_Submitted.Value,
	}

	err = dbOrder.Insert(ctx)
	if err != nil {
		return "", err
	}

	payed, err := new(service.PaymentService).PayByAccount(ctx, storeId, userId, dbOrder.Id, enum.BizCategory_GroupShare.Value)
	if err != nil || !payed {
		// todo
		rerr := new(db.OrderGroupShare).Rollback(ctx, dbOrder.Id)
		if rerr != nil {
			slf.WithError(rerr).Errorw("[MAIN] Rollback err")
		}
		return "", err
	}

	return dbOrder.Id, nil
}

func (t *OrderGroupService) GetGroup(ctx context.Context, userId, id string) (*types.OrderGroup, error) {

	dbGroup, err := new(db.OrderGroup).RequireById(ctx, id)
	if err != nil {
		return nil, errorx.ServerError(err)
	}

	groups, err := t.Assemble(ctx, userId, db.OrderGroups{dbGroup})
	if err != nil {
		return nil, errorx.ServerError(err)
	}

	return groups[0], nil
}

func (t *OrderGroupService) ListJoinableGroups(ctx context.Context, userId, storeId string, page, size int64) (types.OrderGroups, int64, error) {

	dbGroups, total, err := new(db.OrderGroup).List(ctx, db.ListOrderGroupsParams{
		StoreId: storeId, MStatus: enum.JoinableOrderGroupStatus, Page: page, Size: size,
	})

	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	rsp, err := t.Assemble(ctx, userId, dbGroups)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	for _, x := range rsp {
		if x.User.Id == userId {
			x.Tags = append(x.Tags, types.SelfTag)
		}
	}

	return rsp, total, nil
}

func (t *OrderGroupService) ListMyOrderGroups(ctx context.Context, userId, storeId string, page, size int64) (types.OrderGroups, int64, error) {

	groupIds, err := new(db.OrderGroupShare).ListJoinedGroupIds(ctx, storeId, userId)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	dbGroups, total, err := new(db.OrderGroup).ListMyGroups(ctx, userId, storeId, nil, groupIds, page, size)

	if err != nil {
		return nil, 0, errorx.ServerError(err)
	}

	rsp, err := t.Assemble(ctx, userId, dbGroups)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	return rsp, total, nil
}

func (t *OrderGroupService) ListCreatedOrderGroups(ctx context.Context, userId, storeId string, page, size int64) (types.OrderGroups, int64, error) {

	dbGroups, total, err := new(db.OrderGroup).List(ctx, db.ListOrderGroupsParams{
		UserId: userId, StoreId: storeId, Page: page, Size: size,
	})

	if err != nil {
		return nil, 0, errorx.ServerError(err)
	}

	rsp, err := t.Assemble(ctx, userId, dbGroups)
	if err != nil {
		return nil, 0, errorx.ServerError(err)
	}

	return rsp, total, nil
}

func (t *OrderGroupService) RequireByStoreIdAndId(ctx context.Context, storeId, id string) (*types.OrderGroup, error) {

	order, err := new(db.OrderGroup).RequireByIdAndStoreId(ctx, id, storeId)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	orders, err := t.Assemble(ctx, "", db.OrderGroups{order})
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return orders[0], nil
}

func (t *OrderGroupService) Assemble(ctx context.Context, userId string, groups db.OrderGroups) (types.OrderGroups, error) {

	var userIds []string
	_, planIds, groupCreatorIds, storeIds, _ := groups.Ids()
	userIds = append(userIds, groupCreatorIds...)

	stores, err := new(db.Store).ListByIds(ctx, storeIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	storesMap := stores.AsMap()
	_, storeOwnerIds := stores.Ids()
	userIds = append(userIds, storeOwnerIds...)

	plans, err := new(db.Plan).ListByIds(ctx, planIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	plansMap := plans.AsMap()
	_, _, planCreatorIds, issueIds := plans.Ids()
	userIds = append(userIds, planCreatorIds...)

	issues, err := new(db.Issue).ListByIds(ctx, issueIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	issuesTap := issues.AsMap()

	_, itemIds, _, _ := plans.Ids()
	items, err := new(db.Item).ListByIds(ctx, itemIds, false)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	itemsMap := items.AsMap()

	users, err := new(db.User).ListByIds(ctx, userIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	usersMap := users.AsMap()

	var rsp types.OrderGroups

	for _, x := range groups {

		dbPlan := plansMap[x.PlanId]
		dbPlanCreator := usersMap[dbPlan.UserId]

		dbIssue := issuesTap[dbPlan.IssueId()]

		dbItem := itemsMap[dbPlan.ItemId]
		dbGroupCreator := usersMap[x.UserId]

		dbStore := storesMap[x.StoreId]
		dbStoreOwner := usersMap[dbStore.OwnerId]

		y := types.OrderGroup{
			Id:            x.Id,
			Plan:          types.FromDbPlan(dbPlan, dbItem, dbPlanCreator, dbIssue),
			User:          types.FromDbUser(dbGroupCreator),
			Store:         types.FromDbStore(dbStore, dbStoreOwner),
			Amount:        dbPlan.RealAmount(),
			Volume:        x.Volume,
			VolumeOrdered: x.VolumeOrdered,
			Floor:         x.Floor,
			RewardRate:    x.RewardRate,
			Remark:        x.Remark,
			CreatedAt:     timed.SmartTime(x.CreatedAt.Unix()),
			CreatedAtTs:   x.CreatedAt.Unix(),
			Status:        enum.OrderGroupStatus(x.Status),
			Tags:          nil,
			Joinable:      helper.InSlice(x.Status, enum.JoinableOrderGroupStatus),
			JoinerCount:   int64(len(x.Shares)),
			Rejectable:    helper.InSlice(x.Status, enum.RejectableOrderGroupStatus),
			Acceptable:    helper.InSlice(x.Status, enum.AcceptableOrderGroupStatus),
			Ticketable:    helper.InSlice(x.Status, enum.TicketableOrderGroupStatus),
			Switchable:    helper.InSlice(x.Status, enum.TicketableOrderGroupStatus) && x.ToStoreId == "",
			Prized:        x.Status == enum.OrderStatus_Prized.Value,
		}

		userIdsTmp := mapset.NewSet[string]()
		for _, u := range x.Shares {
			userIdsTmp.Add(u)
		}
		y.JoinerCount = int64(len(userIdsTmp.ToSlice()))

		if x.ToStoreId != "" {
			dbToStore := storesMap[x.ToStoreId]
			dbToStoreCreator := usersMap[dbToStore.OwnerId]

			y.ToStore = types.FromDbStore(dbToStore, dbToStoreCreator)
		}

		if x.UserId == userId {
			y.Tags = append(y.Tags, types.SelfTag)
		}

		y.TicketImages = oss.Resources(strings.Split(x.Extra.Tickets, ","))

		rewardInfo := x.Extra.Summary
		if x.Status == enum.OrderGroupStatus_Prized.Value {
			if rewardInfo != "" {
				y.Status.Name += fmt.Sprintf("【%s】", rewardInfo)
				y.Status.Color = "green"
			} else {
				y.Status.Name += fmt.Sprintf("【%s】", "未中奖")
			}
		}

		rsp = append(rsp, &y)
	}

	return rsp, nil
}

func (t *OrderGroupService) AddGroup(ctx context.Context, caller, storeId string, planId string, volume, floor int64, remark string, rewardRate float64) (string, error) {

	plan, err := new(db.Plan).RequireById(ctx, planId)
	if err != nil {
		return "", xerror.Wrap(err)
	}

	// plan状态 pending
	if plan.Status != enum.PlanStatus_Saved.Value {
		return "", errorx.ParamErrorf("only pending plan can be ordered: %s", plan.Id)
	}

	// 合买单创建者必须是plan本人
	if plan.UserId != caller {
		return "", errorx.ParamErrorf("order group creator must be plan owner: expected:%s, actual:%s",
			plan.UserId, caller)
	}

	//amount := (plan.RealAmount() / conv.Float64(plan.Volume)) * conv.Float64(params.Volume)

	dbOrderGroup := db.OrderGroup{
		UserId:        caller,
		PlanId:        planId,
		ItemId:        plan.ItemId,
		IssueIndex:    plan.Issue,
		StoreId:       storeId,
		Volume:        volume,
		VolumeOrdered: 0,
		Floor:         floor,
		Remark:        remark,
		RewardRate:    rewardRate,
		Status:        enum.OrderGroupStatus_Submitted.Value,
	}

	err = comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		insert, err := dbOrderGroup.Insert(ctx, tx)
		if err != nil {
			return err
		}

		if insert {
			_, err := new(db.Plan).UpdateStatus(ctx, tx, planId, enum.PlanStatus_Binding.Value)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return dbOrderGroup.Id, nil
}

func (t *OrderGroupService) SendEvent(ctx context.Context, event string, order *db.OrderGroup) {

	err := comp.SDK().Redis().XAdd(ctx, &redisV9.XAddArgs{
		Stream: event,
		Values: map[string]interface{}{
			"Id":          order.Id,
			"UserId":      order.UserId,
			"ItemId":      order.ItemId,
			"PlanId":      order.PlanId,
			"StoreId":     order.StoreId,
			"ToStoreId":   order.ToStoreId,
			"BizCategory": enum.BizCategory_OrderGroup.Value,
		},
	}).Err()
	if err != nil {
		slf.WithError(err).Errorw("XAdd err")
	}
}
