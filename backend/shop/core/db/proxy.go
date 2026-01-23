package db

import (
	"context"
	"fmt"
	"gitee.com/meepo/backend/kit/components/comp"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-pg/pg/v10"
	"time"
)

type Proxy struct {
	tableName  struct{} `pg:"t_proxies"`
	Id         string
	StoreId    string
	UserId     string
	RewardRate float64
	CreatedAt  time.Time
	Status     string
	Extra      string
}

type Proxies []*Proxy

func (ts Proxies) AsMap() map[string]*Proxy {
	rsp := make(map[string]*Proxy, len(ts))

	for _, t := range ts {
		rsp[t.Id] = t
	}

	return rsp
}

func (ts Proxies) Ids() ([]string, []string) {

	tmp1 := mapset.NewSet[string]()
	tmp2 := mapset.NewSet[string]()
	for _, t := range ts {
		tmp1.Add(t.Id)
		tmp2.Add(t.UserId)
	}

	return tmp1.ToSlice(), tmp2.ToSlice()
}

type ListProxyParams struct {
	Ids        []string
	StoreId    string
	Page, Size int64
}

func (t *Proxy) List(ctx context.Context, params ListProxyParams) (Proxies, int64, error) {

	var tmp Proxies
	q := comp.SDK().Postgres().Model(&tmp).Context(ctx)

	if len(params.Ids) > 0 {
		q = q.WhereIn("id in (?)", params.Ids)
	}

	if params.StoreId != "" {
		q = q.Where("store_id = ?", params.StoreId)
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

func (t *Proxy) CreateNX(ctx context.Context) (*Proxy, error) {

	_, err := comp.SDK().Postgres().Model(t).Context(ctx).
		OnConflict("(store_id, user_id) do nothing").
		Insert()

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (t *Proxy) UpdateRewardRate(ctx context.Context, tx *pg.Tx, id, storeId string, rewardRate float64) (bool, error) {

	update, err := tx.Model(t).Context(ctx).
		Where("id = ?", id).
		Where("store_id = ?", storeId).
		//WhereIn("status in (?)", enum.DeletableProxyStatus).
		Set("reward_rate = ?", rewardRate).
		Update()
	if err != nil {
		return false, err
	}

	return update.RowsAffected() > 0, nil
}

func (t *Proxy) UpdateStatus(ctx context.Context, tx *pg.Tx, id, storeId, toStatus string) (bool, error) {

	update, err := tx.Model(t).Context(ctx).
		Where("id = ?", id).
		Where("store_id = ?", storeId).
		//WhereIn("status in (?)", enum.DeletableProxyStatus).
		Set("status = ?", toStatus).
		Update()
	if err != nil {
		return false, err
	}

	return update.RowsAffected() > 0, nil
}

func (t *Proxy) RequireById(ctx context.Context, id string) (*Proxy, error) {
	proxy, err := t.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	if proxy == nil {
		return nil, fmt.Errorf("no proxy found by id: %s", id)
	}

	return proxy, nil
}

func (t *Proxy) FindById(ctx context.Context, id string) (*Proxy, error) {

	var tmp Proxies
	err := comp.SDK().Postgres().Model(&tmp).Context(ctx).
		Where("id = ?", id).
		Select()

	if err != nil {
		return nil, err
	}

	if len(tmp) == 0 {
		return nil, nil
	}

	return tmp[0], nil
}

func (t *Proxy) FindByUserId(ctx context.Context, storeId, userId string) (*Proxy, error) {

	var tmp Proxies
	err := comp.SDK().Postgres().Model(&tmp).Context(ctx).
		Where("store_id = ?", storeId).
		Where("user_id = ?", userId).
		Select()

	if err != nil {
		return nil, err
	}

	if len(tmp) == 0 {
		return nil, nil
	}

	return tmp[0], nil
}

func (t *Proxy) Update(ctx context.Context, tx *pg.Tx, id, field, value string) (bool, error) {

	update, err := tx.Model((*Proxy)(nil)).Context(ctx).
		Where("id = ?", id).
		Set("extra =  jsonb_set(extra, ?, ?)", fmt.Sprintf("{%s}", field), fmt.Sprintf("\"%s\"", value)).
		Update()

	if err != nil {
		return false, err
	}

	return update.RowsAffected() > 0, nil
}
