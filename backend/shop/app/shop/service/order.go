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
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/types"
	"github.com/go-pg/pg/v10"
	redisV9 "github.com/redis/go-redis/v9"
	"strings"
)

type OrderService struct {
}

func (t *OrderService) GetOwnOrder(ctx context.Context, userId, storeId, orderId string) (*types.Order, error) {

	dbOrder, err := new(db.Order).RequireByIdAndCreatorId(ctx, orderId, userId)
	if err != nil {
		return nil, errorx.ServerError(err)
	}

	//if dbOrder.StoreId != storeId {
	//	return nil, errorx.ServiceErrorf("store permission err, %s, %s", storeId, orderId)
	//}

	orders, err := t.Assemble(ctx, db.Orders{dbOrder})
	if err != nil {
		return nil, errorx.ServerError(err)
	}
	order := orders[0]

	return order, nil
}

func (t *OrderService) ListSubmittedOrders(ctx context.Context, userId, storeId string, page, size int64) (types.Orders, int64, error) {

	orders, total, err := new(db.Order).List(ctx,
		db.ListOrdersParams{
			UserId: userId, StoreId: storeId, MStatus: enum.SubmittedOrderStatus, Page: page, Size: size,
		})
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	rsp, err := t.Assemble(ctx, orders)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	return rsp, total, nil
}

func (t *OrderService) ListFollowableOrders(ctx context.Context, userId, storeId string, page, size int64) (types.Orders, int64, error) {

	orders, total, err := new(db.Order).List(ctx, db.ListOrdersParams{
		StoreId: storeId, MStatus: enum.FollowableStatus, Page: page, Size: size,
	})
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	rsp, err := t.Assemble(ctx, orders)
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

func (t *OrderService) ListMyOrders(ctx context.Context, userId, storeId string, page, size int64) (types.Orders, int64, error) {

	orders, total, err := new(db.Order).List(ctx, db.ListOrdersParams{
		UserId: userId, StoreId: storeId, Page: page, Size: size,
	})
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	rsp, err := t.Assemble(ctx, orders)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	return rsp, total, nil
}

func (t *OrderService) ListToPayOrders(ctx context.Context, userId, storeId string, page, size int64) (types.Orders, int64, error) {

	orders, total, err := new(db.Order).List(ctx, db.ListOrdersParams{
		UserId: userId, StoreId: storeId, MStatus: []string{enum.OrderStatus_Submitted.Value}, Page: page, Size: size,
	})
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	rsp, err := t.Assemble(ctx, orders)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	return rsp, total, nil
}

func (t *OrderService) ListCanceledOrders(ctx context.Context, userId, storeId string, page, size int64) (types.Orders, int64, error) {

	orders, total, err := new(db.Order).List(ctx, db.ListOrdersParams{
		UserId: userId, StoreId: storeId, MStatus: []string{enum.OrderStatus_Canceled.Value}, Page: page, Size: size,
	})
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	rsp, err := t.Assemble(ctx, orders)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	return rsp, total, nil
}

func (t *OrderService) Follow(ctx context.Context, caller, storeId, followOrderId string) (string, error) {

	followOrder, err := new(db.Order).RequireByIdAndStoreId(ctx, followOrderId, storeId)
	if err != nil {
		return "", xerror.Wrap(err)
	}

	if !followOrder.CanBeFollowed() {
		return "", errorx.ParamErrorf("order cannot be followed: %v", conv.S2J(followOrder))
	}

	// 复制 plan
	oldPlan, err := new(db.Plan).RequireById(ctx, followOrder.PlanId)
	if err != nil {
		return "", xerror.Wrap(err)
	}

	index, err := new(IssueService).FindCurrentIssueIndex(ctx, followOrder.ItemId)
	if err != nil {
		return "", xerror.Wrap(err)
	}

	newPlan := oldPlan
	newPlan.Issue = index

	_, err = newPlan.Insert(ctx)
	if err != nil {
		return "", xerror.Wrap(err)
	}

	dbOrder := db.Order{
		UserId:        caller,
		FollowOrderId: followOrderId,
		PlanId:        newPlan.Id,
		ItemId:        followOrder.ItemId,
		StoreId:       storeId,
		Amount:        followOrder.Amount,
		Status:        enum.OrderStatus_Submitted.Value,
	}

	_, inserted, err := dbOrder.InsertFollowOrder(ctx)
	if err != nil {
		return "", errorx.ServerError(xerror.Wrap(err))
	}

	if !inserted {
		return "", errorx.ParamErrorf("duplicate: %v", conv.S2J(followOrder))
	}

	return dbOrder.Id, nil
}

func (t *OrderService) AddOrder(ctx context.Context, caller, storeId, planId string, needUpload bool) (string, error) {

	plan, err := new(db.Plan).RequireById(ctx, planId)
	if err != nil {
		return "", errorx.ParamError(err)
	}

	if plan.Status != enum.PlanStatus_Saved.Value {
		return "", errorx.ParamErrorf("invalid status: %s", plan.Status)
	}

	if plan.UserId != caller {
		return "", errorx.ParamErrorf("permission err : %v", conv.S2J(plan))
	}

	dbOrder := &db.Order{
		UserId:     caller,
		StoreId:    storeId,
		PlanId:     planId,
		ItemId:     plan.ItemId,
		IssueIndex: plan.Issue,
		Status:     enum.OrderStatus_Submitted.Value,
		Amount:     plan.RealAmount(),
		Extra: db.OrderExtra{
			NeedUpload: helper.Choose(needUpload, 1, 0),
		},
	}

	_, err = dbOrder.InsertOrder(ctx)
	if err != nil {
		return "", errorx.ServerError(xerror.Wrap(err))
	}

	return dbOrder.Id, nil
}

func (t *OrderService) CancelOrder(ctx context.Context, userId, orderId string) error {

	dbOrder, err := new(db.Order).RequireByIdAndCreatorId(ctx, orderId, userId)
	if err != nil {
		return errorx.ServerError(err)
	}

	err = comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		canceled, err := new(db.Order).UpdateToCanceled(ctx, tx, orderId, userId)
		if err != nil {
			return xerror.Wrap(err)
		}

		if !canceled {
			return xerror.Wrapf("order cannot cancel: %s", orderId)
		}

		return nil
	})

	if err != nil {
		return xerror.Wrap(err)
	}

	t.SendEvent(ctx, "order.cancel.event", dbOrder)

	return nil
}

func (t *OrderService) SendEvent(ctx context.Context, event string, order *db.Order) {

	err := comp.SDK().Redis().XAdd(ctx, &redisV9.XAddArgs{
		Stream: event,
		Values: map[string]interface{}{
			"Id":        order.Id,
			"UserId":    order.UserId,
			"ItemId":    order.ItemId,
			"PlanId":    order.PlanId,
			"StoreId":   order.StoreId,
			"ToStoreId": order.ToStoreId,
			"Amount":    order.Amount,
		},
	}).Err()
	if err != nil {
		slf.WithError(err).Errorw("XAdd err")
	}
}

// RequireById 用于更新后获取最新数据 todo
func (t *OrderService) RequireByStoreIdAndId(ctx context.Context, storeId, id string) (*types.Order, error) {

	order, err := new(db.Order).RequireByIdAndStoreId(ctx, id, storeId)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	orders, err := t.Assemble(ctx, db.Orders{order})
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return orders[0], nil
}

func (t *OrderService) Assemble(ctx context.Context, orders db.Orders) (types.Orders, error) {

	var userIds []string
	_, planIds, storeIds, orderCreatorIds, _, _ := orders.Ids()
	userIds = append(userIds, orderCreatorIds...)

	plans, err := new(db.Plan).ListByIds(ctx, planIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	plansMap := plans.AsMap()
	_, itemIds, planCreatorIds, issueIds := plans.Ids()
	userIds = append(userIds, planCreatorIds...)

	issues, err := new(db.Issue).ListByIds(ctx, issueIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	issuesTap := issues.AsMap()

	stores, err := new(db.Store).ListByIds(ctx, storeIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	storesMap := stores.AsMap()
	_, storeOwnerIds := stores.Ids()
	userIds = append(userIds, storeOwnerIds...)

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

	var rsp types.Orders

	for _, x := range orders {

		dbPlan := plansMap[x.PlanId]

		dbIssue := issuesTap[dbPlan.IssueId()]

		dbItem := itemsMap[dbPlan.ItemId]
		dbStore := storesMap[x.StoreId]

		dbPlanCreator := usersMap[dbPlan.UserId]
		dbStoreCreator := usersMap[dbStore.OwnerId]

		y := types.Order{
			Id:            x.Id,
			Plan:          types.FromDbPlan(dbPlan, dbItem, dbPlanCreator, dbIssue),
			Store:         types.FromDbStore(dbStore, dbStoreCreator),
			User:          types.FromDbUser(usersMap[x.UserId]),
			Amount:        x.Amount,
			FollowOrderId: x.FollowOrderId,
			CreatedAt:     timed.SmartTime(x.CreatedAt.Unix()),
			CreatedAtTs:   x.CreatedAt.Unix(),
			Status:        enum.OrderStatus(x.Status),
			Cancelable:    x.CanCancelByStatus(),
			Rejectable:    x.CanRejectByStatus(),
			Acceptable:    helper.InSlice(x.Status, enum.AcceptableStatus),
			Ticketable:    helper.InSlice(x.Status, enum.TicketableStatus),
			//Switchable:    helper.InSlice(x.Status, enum.TicketableStatus) && x.ToStoreId == "",
			Followable:    helper.InSlice(x.Status, enum.FollowableStatus),
			Payable:       helper.InSlice(x.Status, enum.PayableStatus),
			KeeperPayable: false,
			Prized:        x.Status == enum.OrderStatus_Prized.Value,
			NeedUpload:    x.Extra.NeedUpload == 1,
		}

		if x.ToStoreId != "" {

			dbToStore := storesMap[x.ToStoreId]
			dbToStoreCreator := usersMap[dbToStore.OwnerId]

			y.ToStore = types.FromDbStore(dbToStore, dbToStoreCreator)
		}

		if x.FollowOrderId != "" {
			y.Tags = append(y.Tags, types.FollowTag)
		}

		rewardInfo := x.Extra.Summary
		if x.Status == enum.OrderStatus_Prized.Value {
			if rewardInfo != "" {
				y.Status.Name += fmt.Sprintf("【%s】", rewardInfo)
				y.Status.Color = "green"
			} else {
				y.Status.Name += fmt.Sprintf("【%s】", "未中奖")
			}
		}

		if x.Extra.Tickets != "" {
			y.TicketImages = oss.Resources(strings.Split(x.Extra.Tickets, ","))
		}

		rsp = append(rsp, &y)
	}

	return rsp, nil
}
