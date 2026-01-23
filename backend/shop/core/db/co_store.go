package db

import (
	"context"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/services/util/pgd"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-pg/pg/v10"
	"time"
)

type CoStore struct {
	tableName struct{} `pg:"t_co_stores"`
	Id        string
	StoreId   string
	CoStoreId string
	Items     map[string]float64
	CreatedAt time.Time
	Status    string
	Extra     CoStoreExtra
}

type CoStoreExtra struct {
	EndProof string `json:"end_proof"`
}

type CoStores []*CoStore

func (ts CoStores) Ids() ([]string, []string, []string) {

	tmp1 := mapset.NewSet[string]()
	tmp2 := mapset.NewSet[string]()
	tmp3 := mapset.NewSet[string]()

	for _, t := range ts {
		tmp1.Add(t.Id)
		tmp2.Add(t.StoreId)
		tmp3.Add(t.CoStoreId)
	}

	return tmp1.ToSlice(), tmp2.ToSlice(), tmp3.ToSlice()

}

func (ts CoStores) AsMap() map[string]*CoStore {
	rsp := make(map[string]*CoStore, len(ts))

	for _, t := range ts {
		rsp[t.Id] = t
	}

	return rsp
}

func (t *CoStore) CreateNX(ctx context.Context) (*CoStore, error) {

	_, err := comp.SDK().Postgres().Model(t).Context(ctx).
		OnConflict("(store_id, co_store_id) do nothing").
		Insert()

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (t *CoStore) UpdateStatus(ctx context.Context, tx *pg.Tx, storeId, coStoreId, toStatus string, fromStatus []string) (bool, error) {

	update, err := tx.Model(t).Context(ctx).
		Where("store_id = ?", storeId).
		Where("co_store_id = ?", coStoreId).
		WhereIn("status in (?)", fromStatus).
		Set("status = ?", toStatus).
		Update()
	if err != nil {
		return false, err
	}

	return update.RowsAffected() > 0, nil
}

func (t *CoStore) UpdateExtra(ctx context.Context, tx *pg.Tx, storeId, coStoreId, field string, value any) (bool, error) {

	q := tx.Model(t).Context(ctx).
		Where("store_id = ?", storeId).
		Where("co_store_id = ?", coStoreId)

	q = pgd.SetJSONField(q, "extra", field, value)

	update, err := q.Update()
	if err != nil {
		return false, err
	}
	if err != nil {
		return false, err
	}

	return update.RowsAffected() > 0, nil
}

func (t *CoStore) UpdateItems(ctx context.Context, tx *pg.Tx, storeId, coStoreId string, items map[string]float64) (bool, error) {

	update, err := tx.Model(t).Context(ctx).
		//Where("id = ?", id).
		Where("store_id = ?", storeId).
		Where("co_store_id = ?", coStoreId).
		//WhereIn("status in (?)", enum.DeletableProxyStatus).
		Set("items = ?", items).
		Update()
	if err != nil {
		return false, err
	}

	return update.RowsAffected() > 0, nil
}

type ListCoStoresParams struct {
	Id        string
	StoreId   string
	CoStoreId string
	MStatus   []string
	Page      int64
	Size      int64
}

func (t *CoStore) List(ctx context.Context, params ListCoStoresParams) (CoStores, int64, error) {

	var tmp CoStores
	q := comp.SDK().Postgres().Model(&tmp).Context(ctx)

	if params.Id != "" {
		q = q.Where("id = ?", params.Id)
	}

	if params.StoreId != "" {
		q = q.Where("store_id = ?", params.StoreId)
	}
	if params.CoStoreId != "" {
		q = q.Where("co_store_id = ?", params.CoStoreId)
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
	count, err := q.OrderExpr("created_at desc").SelectAndCount()

	if err != nil {
		return nil, 0, err
	}

	return tmp, int64(count), nil
}

func (t *CoStore) UpdateToSynced(ctx context.Context, storeId, coStoreId string, field string) (bool, error) {

	if storeId == "" || coStoreId == "" {
		return false, nil
	}

	update, err := comp.SDK().Postgres().Model(t).Context(ctx).
		Set(field+" = ?", 2).
		Where(field+" = 1").
		Where("store_id = ?", storeId).
		Where("co_store_id = ?", coStoreId).
		Update()
	if err != nil {
		return false, err
	}

	return update.RowsAffected() > 0, nil
}
