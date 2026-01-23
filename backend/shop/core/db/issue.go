package db

import (
	"context"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/issue/types"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-pg/pg/v10"
	"time"
)

type Issue struct {
	tableName   struct{} `pg:"t_issues"`
	Id          string
	ItemId      string
	Index       string
	Result      string
	PrizeGrades types.PrizeGrades
	PrizedAt    time.Time
	CloseAt     time.Time
	StartedAt   time.Time
	Status      string
	Extra       string
}

type Issues []*Issue

func (ts Issues) Ids() ([]string, []string) {
	tmp := mapset.NewSet[string]()
	tmp2 := mapset.NewSet[string]()

	for _, t := range ts {
		tmp.Add(t.ItemId)
		tmp2.Add(t.Index)
	}

	return tmp.ToSlice(), tmp2.ToSlice()
}

func (ts Issues) AsMap() map[string]*Issue {
	rsp := make(map[string]*Issue, len(ts))

	for _, t := range ts {
		rsp[t.Id] = t
	}

	return rsp
}

func (t *Issue) Upsert(ctx context.Context) error {
	_, err := comp.SDK().Postgres().Model(t).Context(ctx).
		OnConflict("(item_id, index) do nothing").
		Insert()

	if err != nil {
		return err
	}

	return nil
}

func (t *Issue) FindLatest(ctx context.Context, itemId string) (*Issue, error) {

	var tmp Issues
	err := comp.SDK().Postgres().Model(&tmp).Context(ctx).
		Where("item_id = ?", itemId).
		OrderExpr("index desc").
		Limit(1).
		Select()

	if err != nil {
		return nil, err
	}

	if len(tmp) == 0 {
		return nil, nil
	}

	return tmp[0], nil
}

func (t *Issue) FindUnPrizedIssue(ctx context.Context, limit int) (Issues, error) {

	var tmp Issues
	err := comp.SDK().Postgres().Model(&tmp).Context(ctx).
		//Where("item_id = ?", itemId).
		Where("status != ?", enum.IssueStatus_Prized.Value).
		OrderExpr("index desc").
		Limit(limit).
		Select()

	if err != nil {
		return nil, err
	}

	return tmp, nil
}

func (t *Issue) FindLastUnPrized(ctx context.Context, itemId string) (*Issue, error) {

	var tmp Issues
	err := comp.SDK().Postgres().Model(&tmp).Context(ctx).
		Where("item_id = ?", itemId).
		Where("status != ?", enum.IssueStatus_Prized.Value).
		OrderExpr("index desc").
		Limit(1).
		Select()

	if err != nil {
		return nil, err
	}

	if len(tmp) == 0 {
		return nil, nil
	}

	return tmp[0], nil
}

func (t *Issue) FindByTimestamp(ctx context.Context, itemId string, ts time.Time) (*Issue, error) {

	var tmp Issues
	err := comp.SDK().Postgres().Model(&tmp).Context(ctx).
		Where("item_id = ?", itemId).
		Where("started_at < ?", ts).
		Where("prized_at > ?", ts).
		OrderExpr("index desc").
		Select()

	if err != nil {
		return nil, err
	}

	if len(tmp) == 0 {
		return nil, nil
	}

	return tmp[0], nil

}

func (ts Issues) UpsertBatch(ctx context.Context, tx *pg.Tx) error {

	if len(ts) == 0 {
		return nil
	}

	_, err := tx.Model(&ts).Context(ctx).
		OnConflict("(item_id, index) DO NOTHING").
		Insert()
	if err != nil {
		return err
	}

	return nil
}

func (t *Issue) ListByIds(ctx context.Context, ids []string) (Issues, error) {

	if len(ids) == 0 {
		return nil, nil
	}

	var rsp Issues

	q := comp.SDK().Postgres().Model(&rsp).Context(ctx).
		WhereIn("id in (?)", ids)

	err := q.Select()

	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func (t *Issue) UpdatePrizeResult(ctx context.Context, tx *pg.Tx, id string, result string, prizeGrades types.PrizeGrades, prizedAt time.Time) (bool, error) {

	q := tx.Model((*Issue)(nil)).Context(ctx).Where("id = ?", id)

	if result != "" {
		q = q.Set("result = ?", result)
	}

	if len(prizeGrades) > 0 {
		q = q.Set("prize_grades = ?", prizeGrades)
	}

	update, err := q.Set("prized_at = ?", prizedAt).
		Set("status = ?", enum.IssueStatus_Prized.Value).
		Update()
	if err != nil {
		return false, err
	}

	if err != nil {
		return false, err
	}

	return update.RowsAffected() > 0, nil
}

func (t *Issue) FindByItemIdAndIndexAndStatus(ctx context.Context, itemId, index, status string) (*Issue, error) {

	var rsp Issues
	err := comp.SDK().Postgres().Model(&rsp).Context(ctx).
		Where("item_id = ?", itemId).
		Where("index = ?", index).
		Where("status = ?", status).
		Select()

	if err != nil {
		return nil, err
	}

	if len(rsp) == 0 {
		return nil, nil
	}

	return rsp[0], nil
}

type ListIssuesParams struct {
	Ids            []string
	ItemIds        []string
	MStatus        []string
	MExcludeStatus []string
	Page, Size     int64
}

func (t *Issue) List(ctx context.Context, params ListIssuesParams) (Issues, int64, error) {

	var rsp Issues

	q := comp.SDK().Postgres().Model(&rsp).Context(ctx)

	if params.Ids != nil {
		q = q.WhereIn("id in (?)", params.Ids)
	}

	if params.ItemIds != nil {
		q = q.WhereIn("item_id in (?)", params.ItemIds)
	}
	if params.MStatus != nil {
		q = q.WhereIn("status in (?)", params.MStatus)
	}
	if params.MExcludeStatus != nil {
		q = q.WhereIn("status not in (?)", params.MExcludeStatus)
	}

	if params.Size > 0 {
		page := params.Page
		if page <= 0 {
			page = 1
		}

		q = q.Limit(int(params.Size)).Offset(int((page - 1) * params.Size))
	}

	q = q.OrderExpr("id desc")
	count, err := q.SelectAndCount(&rsp)

	if err != nil {
		return nil, 0, err
	}

	return rsp, int64(count), nil
}
