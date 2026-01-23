package db

import (
	"context"
	"fmt"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/shop/core/enum"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-pg/pg/v10"
	"time"
)

type Plan struct {
	tableName struct{} `pg:"t_plans"`
	Id        string
	ItemId    string
	Issue     string
	Content   string
	Multiple  int64
	//Volume    int64
	Amount    float64
	Type      string
	UserId    string
	StoreId   string
	CreatedAt time.Time
	Status    string
}

func (t *Plan) IssueId() string {
	return t.ItemId + "-" + t.Issue
}

type Plans []*Plan

func (ts Plans) Ids() ([]string, []string, []string, []string) {
	tmp1 := mapset.NewSet[string]()
	tmp2 := mapset.NewSet[string]()
	tmp3 := mapset.NewSet[string]()
	tmp4 := mapset.NewSet[string]()
	for _, t := range ts {
		tmp1.Add(t.Id)
		tmp2.Add(t.ItemId)
		tmp3.Add(t.UserId)
		tmp4.Add(t.IssueId())
	}

	return tmp1.ToSlice(), tmp2.ToSlice(), tmp3.ToSlice(), tmp4.ToSlice()
}

func (t *Plan) RealAmount() float64 {
	return t.Amount * float64(t.Multiple)
}

func (ts Plans) AsMap() map[string]*Plan {

	rsp := make(map[string]*Plan, len(ts))

	for _, t := range ts {
		rsp[t.Id] = t
	}

	return rsp
}

func (t *Plan) Delete(ctx context.Context, id, userId string) (bool, error) {

	result, err := comp.SDK().Postgres().Model(t).Context(ctx).
		Where("id = ?", id).
		Where("user_id = ?", userId).
		Where("status = ?", enum.PlanStatus_Saved.Value).Delete()
	if err != nil {
		return false, err
	}

	if result.RowsAffected() > 0 {
		return true, nil
	}

	return false, nil
}

func (t *Plan) UpdateStatus(ctx context.Context, tx *pg.Tx, id, status string) (bool, error) {

	u, err := tx.Model(t).Context(ctx).Where("id = ?", id).
		Set("status = ?", status).
		Update()

	if err != nil {
		return false, err
	}

	return u.RowsAffected() > 0, nil
}

func (t *Plan) Insert(ctx context.Context) (*Plan, error) {

	_, err := comp.SDK().Postgres().Model(t).Context(ctx).Insert()
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (t *Plan) RequireById(ctx context.Context, id string) (*Plan, error) {

	var rsp Plans
	err := comp.SDK().Postgres().Model(&rsp).Context(ctx).
		Where("id = ?", id).
		//Where("store_id = ?", storeId).
		Select()

	if err != nil {
		return nil, err
	}

	if len(rsp) == 0 {
		return nil, fmt.Errorf("no plan found by id : %s", id)
	}

	return rsp[0], nil
}

func (t *Plan) ListByIds(ctx context.Context, ids []string) (Plans, error) {

	if len(ids) == 0 {
		return nil, nil
	}

	var rsp Plans

	err := comp.SDK().Postgres().Model(&rsp).Context(ctx).
		WhereIn("id in (?)", ids).
		Select()

	if err != nil {
		return nil, err
	}

	return rsp, nil
}

type ListPlansParams struct {
	Ids        []string
	UserId     string
	StoreId    string
	MStatus    []string
	Page, Size int64
}

func (t *Plan) List(ctx context.Context, params ListPlansParams) (Plans, int64, error) {

	var rsp Plans
	q := comp.SDK().Postgres().Model(&rsp).Context(ctx)

	if len(params.Ids) > 0 {
		q = q.WhereIn("id in (?)", params.Ids)
	}
	if params.UserId != "" {
		q = q.Where("user_id = ?", params.UserId)
	}

	if params.StoreId != "" {
		q = q.Where("store_id = ?", params.StoreId)
	}

	if len(params.MStatus) > 0 {
		q = q.WhereIn("status in (?)", params.MStatus)
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

func (t *Plan) FindByStatus(ctx context.Context, status string, page, limit int64) (Plans, error) {

	var rsp Plans
	q := comp.SDK().Postgres().Model(&rsp).Context(ctx).
		Where("status = ?", status)

	if page > 0 && limit > 0 {
		q = q.Limit(int(limit)).Offset(int((page - 1) * limit))
	}
	err := q.Select()

	if err != nil {
		return nil, err
	}

	return rsp, nil
}
