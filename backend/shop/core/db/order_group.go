package db

import (
	"context"
	"fmt"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/helper/mathd"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/kit/services/util/pgd"
	"gitee.com/meepo/backend/shop/core/enum"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-pg/pg/v10"
	"strings"
	"time"
)

type OrderGroup struct {
	tableName     struct{} `pg:"t_order_groups"`
	Id            string
	UserId        string
	PlanId        string
	ItemId        string
	IssueIndex    string
	StoreId       string
	ToStoreId     string
	Volume        int64
	VolumeOrdered int64
	Floor         int64
	Remark        string
	RewardRate    float64
	CreatedAt     time.Time
	Status        string
	Shares        map[string]string
	Extra         OrderGroupExtra
}

func (t *OrderGroup) IssueId() string {
	return t.ItemId + "-" + t.IssueIndex
}

type OrderGroupExtra struct {
	Tickets string `json:"tickets,omitempty"`
	Summary string `json:"summary,omitempty"`
}

type GroupExtra struct {
}

func (t *OrderGroup) VolumeLeft() int64 {
	return mathd.Max(t.Volume-t.VolumeOrdered, 0)
}

type OrderGroups []*OrderGroup

func (ts OrderGroups) AsMap() map[string]*OrderGroup {
	rsp := make(map[string]*OrderGroup, len(ts))

	for _, t := range ts {
		rsp[t.Id] = t
	}

	return rsp
}

func (ts OrderGroups) Ids() ([]string, []string, []string, []string, []string) {
	tmp1 := mapset.NewSet[string]()
	tmp2 := mapset.NewSet[string]()
	tmp3 := mapset.NewSet[string]()
	tmp4 := mapset.NewSet[string]()
	tmp5 := mapset.NewSet[string]()
	for _, t := range ts {
		tmp1.Add(t.Id)
		tmp2.Add(t.PlanId)
		tmp3.Add(t.UserId)
		tmp4.Add(t.StoreId)
		if t.ToStoreId != "" {
			tmp4.Add(t.ToStoreId)
		}
		tmp5.Add(t.IssueId())
	}

	return tmp1.ToSlice(), tmp2.ToSlice(), tmp3.ToSlice(), tmp4.ToSlice(), tmp5.ToSlice()
}

func (t *OrderGroup) UpdateToStoreId(ctx context.Context, tx *pg.Tx, id, toStoreId string) (bool, error) {
	order := OrderGroup{Id: id}

	updated, err := tx.Model(&order).Context(ctx).
		Set("to_store_id = ?", toStoreId).
		Where("id = ?", id).
		Where("to_store_id is null").
		WhereIn("status in (?)", enum.TicketableOrderGroupStatus).
		Update()

	if err != nil {
		return false, xerror.Wrap(err)
	}

	return updated.RowsAffected() > 0, nil

}

func (t *OrderGroup) RequireById(ctx context.Context, id string) (*OrderGroup, error) {

	var tmp OrderGroups
	err := comp.SDK().Postgres().WithContext(ctx).Model(&tmp).
		Where("id = ?", id).
		Select()

	if err != nil {
		return nil, err
	}

	if len(tmp) == 0 {
		return nil, fmt.Errorf("no order found by: %s", id)
	}

	return tmp[0], nil
}

func (t *OrderGroup) FindByPlanId(ctx context.Context, planId string) (OrderGroups, error) {

	var rsp OrderGroups
	err := comp.SDK().Postgres().WithContext(ctx).Model(&rsp).
		Where("plan_id = ?", planId).
		Select()

	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func (t *OrderGroup) RequireByIdAndStoreId(ctx context.Context, id, storeId string) (*OrderGroup, error) {

	var tmp OrderGroups
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

func (t *OrderGroup) CountByStatus(ctx context.Context, storeId string, status []string) (int64, error) {
	total, err := comp.SDK().Postgres().WithContext(ctx).Model(t).
		WhereIn("status in (?)", status).
		Where("store_id = ?", storeId).
		Count()

	if err != nil {
		return 0, err
	}

	return int64(total), nil
}

func (t *OrderGroup) RequireByIdAndStoreIdOrToStoreId(ctx context.Context, id, storeId string) (*OrderGroup, error) {

	var tmp OrderGroups
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

func (t *OrderGroup) ListMyGroups(ctx context.Context, userId, storeId string, status, ids []string, page, size int64) (OrderGroups, int64, error) {

	if len(ids) == 0 {
		return nil, 0, nil
	}

	var rsp OrderGroups
	q := comp.SDK().Postgres().Model(&rsp).Context(ctx).
		Where("user_id = ?", userId).
		Where("store_id = ?", storeId).
		WhereOr("id in (?)", pg.In(ids))

	if len(status) > 0 {
		q = q.WhereIn("status in (?)", status)
	}

	if page > 0 && size > 0 {
		q = q.Limit(int(size)).Offset(int((page - 1) * size))
	}
	count, err := q.OrderExpr("id desc").SelectAndCount()

	if err != nil {
		return nil, 0, err
	}

	return rsp, int64(count), nil
}

func (t *OrderGroup) ListNotCleanOrdersByStatus(ctx context.Context, status []string, limit int) (OrderGroups, error) {
	var tmp OrderGroups
	err := comp.SDK().Postgres().Model(&tmp).Context(ctx).
		WhereIn("status in (?)", status).
		Where("clean = 0").
		Limit(limit).
		Select()

	if err != nil {
		return nil, err
	}

	return tmp, nil
}

func (t *OrderGroup) UpdateToClean(ctx context.Context, id string) (bool, error) {
	update, err := comp.SDK().Postgres().Model(t).Context(ctx).
		Set("clean = ?", 1).
		Where("clean = 0").
		Where("id = ?", id).
		Update()
	if err != nil {
		return false, err
	}

	return update.RowsAffected() > 0, nil
}

func (t *OrderGroup) ListKeeperOrders(ctx context.Context, userId, storeId string, status, ids []string, page, size int64) (OrderGroups, int64, error) {

	var rsp OrderGroups
	q := comp.SDK().Postgres().Model(&rsp).Context(ctx)

	if userId != "" {
		q = q.Where("user_id = ?", userId)
	}

	if storeId != "" {
		q = q.Where("store_id = ? or to_store_id = ?", storeId, storeId)
	}

	if len(status) > 0 {
		q = q.WhereIn("status in (?)", status)
	}

	if len(ids) > 0 {
		q = q.WhereIn("id in (?)", ids)
	}

	if page > 0 && size > 0 {
		q = q.Limit(int(size)).Offset(int((page - 1) * size))
	}
	count, err := q.OrderExpr("id desc").SelectAndCount()

	if err != nil {
		return nil, 0, err
	}

	return rsp, int64(count), nil
}

type ListOrderGroupsParams struct {
	Ids                 []string
	StoreIdOrToStoreId  string
	ItemId              string
	ExcludeIssueIndexes []string
	UserId              string
	StoreId             string
	MStatus             []string
	Page, Size          int64
}

func (t *OrderGroup) List(ctx context.Context, params ListOrderGroupsParams) (OrderGroups, int64, error) {

	var rsp OrderGroups
	q := comp.SDK().Postgres().Model(&rsp).Context(ctx)

	if params.UserId != "" {
		q = q.Where("user_id = ?", params.UserId)
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

	if params.StoreId != "" {
		q = q.Where("store_id = ?", params.StoreId)
	}

	if len(params.MStatus) > 0 {
		q = q.WhereIn("status in (?)", params.MStatus)
	}

	if len(params.Ids) > 0 {
		q = q.WhereIn("id in (?)", params.Ids)
	}

	if params.Size > 0 {
		page := params.Page
		if page <= 0 {
			page = 1
		}

		q = q.Limit(int(params.Size)).Offset(int((page - 1) * params.Size))
	}

	count, err := q.OrderExpr("id desc").SelectAndCount()

	if err != nil {
		return nil, 0, err
	}

	return rsp, int64(count), nil
}

func (t *OrderGroup) Insert(ctx context.Context, tx *pg.Tx) (bool, error) {

	inserted, err := tx.Model(t).Context(ctx).
		OnConflict("(plan_id) do nothing").
		Insert()
	if err != nil {
		return false, err
	}

	return inserted.RowsAffected() > 0, nil
}

func (t *OrderGroup) ListByIds(ctx context.Context, ids []string) (OrderGroups, error) {

	if len(ids) == 0 {
		return nil, nil
	}

	var rsp OrderGroups

	q := comp.SDK().Postgres().Model(&rsp).Context(ctx).
		WhereIn("id in (?)", ids)

	err := q.Select()

	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func (t *OrderGroup) UpdateToAccepted(ctx context.Context, tx *pg.Tx, id string) (bool, error) {
	order := OrderGroup{Id: id}

	updated, err := tx.Model(&order).Context(ctx).
		Set("status = ?", enum.OrderGroupStatus_Accepted.Value).
		Where("id = ?", id).
		WhereIn("status in (?)", enum.AcceptableOrderGroupStatus).
		Update()

	if err != nil {
		return false, xerror.Wrap(err)
	}

	return updated.RowsAffected() > 0, nil

}

func (t *OrderGroup) UpdateToRejected(ctx context.Context, tx *pg.Tx, id, reasonId string) (bool, error) {
	order := OrderGroup{Id: id}

	updated, err := tx.Model(&order).Context(ctx).
		Set("status = ?", enum.OrderGroupStatus_Rejected.Value).
		Set("extra =  jsonb_set(extra, ?, ?)", "{rejectReasonId}", reasonId).
		Where("id = ?", id).
		WhereIn("status in (?)", enum.RejectableOrderGroupStatus).
		Update()

	if err != nil {
		return false, xerror.Wrap(err)
	}

	return updated.RowsAffected() > 0, nil

}

func (t *OrderGroup) UpdateToTicketed(ctx context.Context, tx *pg.Tx, id string, images []string) (bool, error) {
	order := OrderGroup{Id: id}

	q := tx.Model(&order).Context(ctx).
		Set("status = ?", enum.OrderGroupStatus_Ticketed.Value)

	if len(images) > 0 {
		q = q.Set("extra =  jsonb_set(extra, ?, ?)", "{tickets}", fmt.Sprintf("\"%s\"", strings.Join(images, ",")))
	}

	updated, err := q.Where("id = ?", id).
		WhereIn("status in (?)", enum.TicketableOrderGroupStatus).
		Update()

	if err != nil {
		return false, xerror.Wrap(err)
	}

	return updated.RowsAffected() > 0, nil

}

func (t *OrderGroup) FindUnPrizedOrders(ctx context.Context, itemId string, limit int) (OrderGroups, error) {

	// 非合买订单
	var rsp OrderGroups
	err := comp.SDK().Postgres().Model(&rsp).Context(ctx).
		Where("item_id = ?", itemId).
		Where("status = ?", enum.OrderGroupStatus_Ticketed.Value).
		OrderExpr("id").
		Limit(limit).
		Select()

	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func (t *OrderGroup) UpdateToPrized(ctx context.Context, tx *pg.Tx, orderIds []string) (bool, error) {

	//q := tx.Model(t).Context(ctx).
	//	Set("status = ?", enum.OrderStatus_Prized.Value)
	////
	////if reward != "" {
	////	q = q.Set("extra =  jsonb_set(extra, ?, ?)", "{reward}", fmt.Sprintf("\"%s\"", reward))
	////}

	updated, err := tx.Model(t).Context(ctx).
		Set("status = ?", enum.OrderGroupStatus_Prized.Value).
		WhereIn("id in (?)", orderIds).
		WhereIn("status in (?)", enum.PrizableGroupStatus).
		Update()
	if err != nil {
		return false, err
	}

	return updated.RowsAffected() >= len(orderIds), nil
}

func (t *OrderGroup) UpdateExtra(ctx context.Context, tx *pg.Tx, orderId, field string, value any) (bool, error) {

	q := tx.Model(t).Context(ctx).Where("id = ?", orderId)
	update, err := pgd.SetJSONField(q, "extra", field, value).
		Update()
	if err != nil {
		return false, err
	}

	return update.RowsAffected() > 0, nil
}

func (t *OrderGroup) UpdateToSynced(ctx context.Context, ids []string, field string) (bool, error) {

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

//
//func (t *OrderGroup) UpdateToPrized(ctx context.Context, tx *pg.Tx, orderId string, reward string) (bool, error) {
//
//	plan := OrderGroup{Id: orderId}
//
//	q := tx.Model(&plan).Context(ctx).
//		Set("status = ?", enum.OrderGroupStatus_Prized.Value)
//
//	if reward != "" {
//		q = q.Set("extra =  jsonb_set(extra, ?, ?)", "{reward}", fmt.Sprintf("\"%s\"", reward))
//	}
//
//	updated, err := q.WhereIn("status in (?)", enum.PrizableGroupStatus).
//		Update()
//	if err != nil {
//		return false, err
//	}
//
//	return updated.RowsAffected() > 0, nil
//}
