package db

import (
	"context"
	"fmt"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/kit/services/util/gind/errorx"
	"gitee.com/meepo/backend/kit/services/util/pgd"
	"gitee.com/meepo/backend/shop/core/enum"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-pg/pg/v10"
	"strings"
	"time"
)

type Order struct {
	tableName struct{} `pg:"t_orders,alias:o"`

	Id            string
	UserId        string
	PlanId        string
	ItemId        string
	IssueIndex    string
	StoreId       string
	Amount        float64
	FollowOrderId string
	CreatedAt     time.Time
	Status        string
	Extra         OrderExtra
	ToStoreId     string
}

type OrderExtra struct {
	NeedUpload int    `json:"needUpload"`
	Summary    string `json:"summary"`
	Tickets    string `json:"tickets"`
}

func (t *Order) IssueId() string {
	return t.ItemId + "-" + t.IssueIndex
}

//
//func (t *Order) NeedUpload() bool {
//	if t.Extra == "" || t.Extra == "{}" {
//		return false
//	}
//
//	var mp map[string]any
//
//	err := conv.J2M(t.Extra, &mp)
//	if err != nil {
//		slf.WithError(err).Errorw("J2M err")
//		return false
//	}
//
//	return conv.Int(mp["needUpload"]) == 1
//}
//
//func (t *Order) FindRewardInfo() string {
//
//	if t.Extra == "" || t.Extra == "{}" {
//		return ""
//	}
//
//	var mp map[string]any
//
//	err := conv.J2M(t.Extra, &mp)
//	if err != nil {
//		slf.WithError(err).Errorw("J2M err")
//		return ""
//	}
//
//	return conv.String(mp["reward"])
//}

type Orders []*Order

func (t *Order) CanBeFollowed() bool {
	if !helper.InSlice(t.Status, enum.FollowableStatus) {
		return false
	}

	//if t.GroupId != "" {
	//	return false
	//}
	return true
}
func (t *Order) CanCancelByStatus() bool {
	return helper.InSlice(t.Status, enum.CancelableStatus)
}
func (t *Order) CanRejectByStatus() bool {
	return helper.InSlice(t.Status, enum.RejectableStatus)
}

func (ts Orders) Ids() ([]string, []string, []string, []string, []string, []string) {
	tmp1 := mapset.NewSet[string]()
	tmp2 := mapset.NewSet[string]()
	tmp3 := mapset.NewSet[string]()
	tmp4 := mapset.NewSet[string]()
	tmp6 := mapset.NewSet[string]()
	tmp7 := mapset.NewSet[string]()
	for _, t := range ts {
		tmp1.Add(t.Id)
		tmp2.Add(t.PlanId)
		tmp3.Add(t.StoreId)
		if t.ToStoreId != "" {
			tmp3.Add(t.ToStoreId)
		}
		tmp4.Add(t.UserId)
		tmp6.Add(t.ItemId)
		tmp7.Add(t.ItemId + "-" + t.IssueIndex)
	}

	return tmp1.ToSlice(), tmp2.ToSlice(), tmp3.ToSlice(), tmp4.ToSlice(), tmp6.ToSlice(), tmp7.ToSlice()
}

func (t *Order) UpdateToSynced(ctx context.Context, ids []string, field string) (bool, error) {
	if len(ids) == 0 {
		return false, nil
	}

	update, err := comp.SDK().Postgres().Model(t).Context(ctx).
		Set(field+" = ?", 2).
		Where(field+" = 1").
		WhereIn("id in (?)", ids).
		Update()
	if err != nil {
		return false, err
	}

	return update.RowsAffected() > 0, nil
}

type OrderAgg struct {
	UserId      string
	OrderCount  int64
	OrderAmount float64
}

type OrderAggs []*OrderAgg

func (ts OrderAggs) AsMap() map[string]*OrderAgg {

	rsp := make(map[string]*OrderAgg, len(ts))

	for _, t := range ts {
		rsp[t.UserId] = t
	}

	return rsp

}

func (t *Order) AggregateByUserIds(ctx context.Context, storeId string, userIds []string) (OrderAggs, error) {

	if len(userIds) == 0 {
		return nil, nil
	}

	var tmp OrderAggs

	err := comp.SDK().Postgres().WithContext(ctx).Model(t).
		ColumnExpr("user_id").
		ColumnExpr("count(1) as order_count").
		ColumnExpr("sum(amount) as order_amount").
		WhereIn("status in (?)", enum.ProxyRewardableStatus).
		WhereIn("user_id in (?)", userIds).
		Where("store_id = ?", storeId).
		GroupExpr("user_id").Select(&tmp)
	if err != nil {
		return nil, err
	}
	return tmp, nil
}

func (t *Order) CountByStatus(ctx context.Context, storeId string, status []string) (int64, error) {
	total, err := comp.SDK().Postgres().WithContext(ctx).Model(t).
		WhereIn("status in (?)", status).
		Where("store_id = ?", storeId).
		Count()

	if err != nil {
		return 0, err
	}

	return int64(total), nil
}

func (t *Order) RequireByIdAndCreatorId(ctx context.Context, id, userId string) (*Order, error) {

	var tmp Orders
	err := comp.SDK().Postgres().WithContext(ctx).Model(&tmp).
		Where("id = ?", id).
		Where("user_id = ?", userId).
		Select()

	if err != nil {
		return nil, err
	}

	if len(tmp) == 0 {
		return nil, fmt.Errorf("no order found by: %s, %s", id, userId)
	}

	return tmp[0], nil
}

func (t *Order) RequireByIdAndStoreIdOrToStoreId(ctx context.Context, id, storeId string) (*Order, error) {

	var tmp Orders
	err := comp.SDK().Postgres().WithContext(ctx).Model(&tmp).
		Where("id = ?", id).
		Where("store_id = ? or to_store_id = ?", storeId, storeId).
		Select()

	if err != nil {
		return nil, err
	}

	if len(tmp) == 0 {
		return nil, fmt.Errorf("no order found by: %s, %s", id, storeId)
	}

	return tmp[0], nil
}

func (t *Order) RequireByIdAndStoreId(ctx context.Context, id, storeId string) (*Order, error) {

	var tmp Orders
	err := comp.SDK().Postgres().WithContext(ctx).Model(&tmp).
		Where("id = ?", id).
		Where("store_id = ?", storeId).
		Select()

	if err != nil {
		return nil, err
	}

	if len(tmp) == 0 {
		return nil, fmt.Errorf("no order found by: %s, %s", id, storeId)
	}

	return tmp[0], nil
}

func (t *Order) FindByPlanId(ctx context.Context, planId string) (Orders, error) {

	var rsp Orders
	err := comp.SDK().Postgres().WithContext(ctx).Model(&rsp).
		Where("plan_id = ?", planId).
		Select()

	if err != nil {
		return nil, err
	}

	return rsp, nil
}

type ListOrdersParams struct {
	Id                  string
	StoreIdOrToStoreId  string
	ItemId              string
	ExcludeIssueIndexes []string
	UserId              string
	ExcludeUserId       string
	StoreId             string
	ToStoreIdNotEmpty   bool
	CreatedAtFrom       *time.Time
	CreatedAtTo         *time.Time
	MStatus             []string
	SyncSwitch          int
	SyncRollback        int
	Page                int64
	Size                int64
}

func (t *Order) List(ctx context.Context, params ListOrdersParams) (Orders, int64, error) {

	var rsp Orders
	q := comp.SDK().Postgres().Model(&rsp).Context(ctx)

	if params.Id != "" {
		q = q.Where("id = ?", params.Id)
	}

	if params.ItemId != "" {
		q = q.Where("item_id = ?", params.ItemId)
	}
	if len(params.ExcludeIssueIndexes) > 0 {
		q = q.WhereIn("issue_index not in (?)", params.ExcludeIssueIndexes)
	}

	if params.StoreIdOrToStoreId != "" {
		q = q.Where("store_id = ? or to_store_id = ?", params.StoreIdOrToStoreId, params.StoreIdOrToStoreId)
	}

	if params.UserId != "" {
		q = q.Where("user_id = ?", params.UserId)
	}
	if params.ExcludeUserId != "" {
		q = q.Where("user_id != ?", params.ExcludeUserId)
	}

	if params.StoreId != "" {
		q = q.Where("store_id = ?", params.StoreId)
	}

	if len(params.MStatus) > 0 {
		q = q.WhereIn("status in (?)", params.MStatus)
	}

	if params.CreatedAtFrom != nil {
		q = q.Where("created_at >= ?", params.CreatedAtFrom)
	}
	if params.CreatedAtTo != nil {
		q = q.Where("created_at <= ?", params.CreatedAtFrom)
	}

	if params.ToStoreIdNotEmpty {
		q = q.Where("to_store_id != ''")
	}

	if params.SyncSwitch > 0 {
		q = q.Where("sync_switch = ?", params.SyncSwitch)
	}
	if params.SyncRollback > 0 {
		q = q.Where("sync_rollback = ?", params.SyncRollback)
	}

	if params.Size > 0 {
		page := params.Page
		if page <= 0 {
			page = 1
		}

		q = q.Limit(int(params.Size)).Offset(int((page - 1) * params.Size))
	}

	q = q.OrderExpr("created_at desc")
	count, err := q.SelectAndCount()

	if err != nil {
		return nil, 0, err
	}

	return rsp, int64(count), nil
}

func (t *Order) InsertFollowOrder(ctx context.Context) (*Order, bool, error) {
	inserted, err := comp.SDK().Postgres().Model(t).Context(ctx).
		//Where("user_id = ?", t.UserId).
		//Where("follow_order_id = ?", t.FollowOrderId).
		//OnConflict("(user_id, follow_order_id) do nothing").
		Insert()

	if err != nil {
		return nil, false, err
	}

	return t, inserted.RowsAffected() > 0, nil
}

func (t *Order) GetOrderCounts(ctx context.Context, userIds []string) (map[string]int64, error) {

	rsp := make(map[string]int64, len(userIds))

	if len(userIds) == 0 {
		return rsp, nil
	}

	var counts []struct {
		UserId string
		Count  int64
	}

	err := comp.SDK().Postgres().Model(t).Context(ctx).
		ColumnExpr("user_id").
		ColumnExpr("count(1) as count").
		WhereIn("user_id in (?)", userIds).
		GroupExpr("user_id").
		Select(&counts)
	if err != nil {
		return nil, err
	}

	for _, x := range counts {
		rsp[x.UserId] = x.Count
	}

	return rsp, nil
}

func (t *Order) InsertOrder(ctx context.Context) (*Order, error) {

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		inserted, err := tx.Model(t).Context(ctx).
			//OnConflict("(user_id, plan_id) where (group_id IS NULL) do nothing").
			Insert()

		if err != nil {
			return err
		}

		if inserted.RowsAffected() == 0 {
			return errorx.ServiceErrorf("duplicate order: userId %s, planId %s", t.UserId, t.PlanId)
		}

		_, err = tx.Model((*Plan)(nil)).Where("id = ?", t.PlanId).
			Set("status = ?", enum.PlanStatus_Binding.Value).
			Update()

		if err != nil {
			return err
		}

		return nil
	})

	return t, err
}

func (t *Order) UpdateToCanceled(ctx context.Context, tx *pg.Tx, id, userId string) (bool, error) {

	order := Order{Id: id}

	updated, err := tx.Model(&order).Context(ctx).
		Where("id = ?", id).
		Where("user_id = ?", userId).
		WhereIn("status in (?)", enum.CancelableStatus).
		Set("status = ?", enum.OrderStatus_Canceled.Value).
		Update()

	if err != nil {
		return false, xerror.Wrap(err)
	}

	return updated.RowsAffected() > 0, nil
}

func (t *Order) UpdateToAccepted(ctx context.Context, tx *pg.Tx, id string) (bool, error) {
	order := Order{Id: id}

	updated, err := tx.Model(&order).Context(ctx).
		Set("status = ?", enum.OrderStatus_Accepted.Value).
		Where("id = ?", id).
		WhereIn("status in (?)", enum.AcceptableStatus).
		Update()

	if err != nil {
		return false, xerror.Wrap(err)
	}

	return updated.RowsAffected() > 0, nil

}

func (t *Order) UpdateToStoreId(ctx context.Context, tx *pg.Tx, id, toStoreId string) (bool, error) {
	order := Order{Id: id}

	updated, err := tx.Model(&order).Context(ctx).
		Set("to_store_id = ?", toStoreId).
		Where("id = ?", id).
		Where("to_store_id = ''").
		WhereIn("status in (?)", enum.TicketableStatus).
		Update()

	if err != nil {
		return false, xerror.Wrap(err)
	}

	return updated.RowsAffected() > 0, nil

}

func (t *Order) UpdateToRejected(ctx context.Context, tx *pg.Tx, id, reasonId string) (bool, error) {
	order := Order{Id: id}

	updated, err := tx.Model(&order).Context(ctx).
		Set("status = ?", enum.OrderStatus_Rejected.Value).
		Set("extra =  jsonb_set(extra, ?, ?)", "{rejectReasonId}", reasonId).
		Where("id = ?", id).
		WhereIn("status in (?)", enum.RejectableStatus).
		Update()

	if err != nil {
		return false, xerror.Wrap(err)
	}

	return updated.RowsAffected() > 0, nil

}

func (t *Order) UpdateToTicketed(ctx context.Context, tx *pg.Tx, id string, images []string) (bool, error) {
	order := Order{Id: id}

	q := tx.Model(&order).Context(ctx).
		Set("status = ?", enum.OrderStatus_Ticketed.Value)

	if len(images) > 0 {
		q = q.Set("extra =  jsonb_set(extra, ?, ?)", "{tickets}", fmt.Sprintf("\"%s\"", strings.Join(images, ",")))
	}

	updated, err := q.Where("id = ?", id).
		WhereIn("status in (?)", enum.TicketableStatus).
		Update()

	if err != nil {
		return false, xerror.Wrap(err)
	}

	return updated.RowsAffected() > 0, nil

}

func (t *Order) UpdateToPayed(ctx context.Context, tx *pg.Tx, id string, byKeeper bool) (bool, error) {

	order := Order{Id: id}

	q := tx.Model(&order).Context(ctx).
		Set("status = ?", enum.OrderStatus_Payed.Value)

	if byKeeper {
		q = q.Set("extra =  jsonb_set(extra, ?, ?)", "{byKeeper}", "1")
	}

	updated, err := q.Where("id = ?", id).
		WhereIn("status in (?)", enum.PayableStatus).
		Update()

	if err != nil {
		return false, xerror.Wrap(err)
	}

	return updated.RowsAffected() > 0, nil

}

func (t *Order) UpdateToPrized(ctx context.Context, tx *pg.Tx, orderIds []string) (bool, error) {

	//q := tx.Model(t).Context(ctx).
	//	Set("status = ?", enum.OrderStatus_Prized.Value)
	////
	////if reward != "" {
	////	q = q.Set("extra =  jsonb_set(extra, ?, ?)", "{reward}", fmt.Sprintf("\"%s\"", reward))
	////}

	updated, err := tx.Model(t).Context(ctx).
		Set("status = ?", enum.OrderStatus_Prized.Value).
		WhereIn("id in (?)", orderIds).
		WhereIn("status in (?)", enum.PrizableStatus).
		Update()
	if err != nil {
		return false, err
	}

	return updated.RowsAffected() >= len(orderIds), nil
}

func (t *Order) UpdateExtra(ctx context.Context, tx *pg.Tx, orderId, field string, value any) (bool, error) {

	q := tx.Model(t).Context(ctx).Where("id = ?", orderId)
	update, err := pgd.SetJSONField(q, "extra", field, value).
		Update()
	if err != nil {
		return false, err
	}

	return update.RowsAffected() > 0, nil
}

func (t *Order) FindShouldBePrizedOrders(ctx context.Context, itemId string, limit int) (Orders, error) {

	var rsp Orders
	err := comp.SDK().Postgres().Model(&rsp).Context(ctx).
		Where("item_id = ?", itemId).
		Where("status = ?", enum.OrderStatus_Ticketed.Value).
		OrderExpr("id").
		Limit(limit).
		Select()

	if err != nil {
		return nil, err
	}

	return rsp, nil
}
