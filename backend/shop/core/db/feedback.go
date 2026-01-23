package db

import (
	"context"
	"gitee.com/meepo/backend/kit/components/comp"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-pg/pg/v10"
	"time"
)

type Feedback struct {
	tableName struct{} `pg:"t_feedbacks"`
	Id        string
	StoreId   string
	UserId    string
	Text      string
	CreatedAt time.Time
	Status    string
	Extra     string
}

type Feedbacks []*Feedback

func (ts Feedbacks) Ids() ([]string, []string, []string) {

	tmp1 := mapset.NewSet[string]()
	tmp2 := mapset.NewSet[string]()
	tmp3 := mapset.NewSet[string]()
	for _, t := range ts {
		tmp1.Add(t.Id)
		tmp2.Add(t.UserId)
		tmp3.Add(t.StoreId)
	}

	return tmp1.ToSlice(), tmp2.ToSlice(), tmp3.ToSlice()
}

func (ts Feedbacks) AsMap() map[string]*Feedback {
	rsp := make(map[string]*Feedback, len(ts))

	for _, t := range ts {
		rsp[t.Id] = t
	}

	return rsp
}

type ListFeedbackParams struct {
	Id         string
	UserId     string
	StoreId    string
	Page, Size int64
}

func (t *Feedback) CountByStatus(ctx context.Context, storeId string, status []string) (int64, error) {
	total, err := comp.SDK().Postgres().WithContext(ctx).Model(t).
		WhereIn("status in (?)", status).
		Where("store_id = ?", storeId).
		Count()

	if err != nil {
		return 0, err
	}

	return int64(total), nil
}

func (t *Feedback) List(ctx context.Context, params ListFeedbackParams) (Feedbacks, int64, error) {

	var tmp Feedbacks
	q := comp.SDK().Postgres().Model(&tmp).Context(ctx)

	if params.Id != "" {
		q = q.Where("id = ?", params.Id)
	}
	if params.StoreId != "" {
		q = q.Where("store_id = ?", params.StoreId)
	}

	if params.UserId != "" {
		q = q.Where("user_id = ?", params.UserId)
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

	return tmp, int64(count), nil
}

func (t *Feedback) Add(ctx context.Context) (*Feedback, error) {

	_, err := comp.SDK().Postgres().Model(t).Context(ctx).
		Insert()

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (t *Feedback) UpdateStatus(ctx context.Context, tx *pg.Tx, id, storeId, toStatus string) (bool, error) {

	update, err := tx.Model(t).Context(ctx).
		Where("id = ?", id).
		Where("store_id = ?", storeId).
		//WhereIn("status in (?)", enum.DeletableFeedbackStatus).
		Set("status = ?", toStatus).
		Update()
	if err != nil {
		return false, err
	}

	return update.RowsAffected() > 0, nil
}
