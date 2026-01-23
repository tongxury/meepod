package db

import (
	"context"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/shop/core/enum"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-pg/pg/v10"
	"time"
)

type ProxyReward struct {
	tableName    struct{} `pg:"t_proxy_rewards"`
	Id           string
	Month        string
	ProxyId      string
	ProxyUserId  string
	StoreId      string
	UserCount    int64
	OrderCount   int64
	OrderAmount  float64
	RewardRate   float64
	RewardAmount float64
	Status       string
	PayAt        time.Time
	CreatedAt    time.Time
	Extra        string
}

type ProxyRewards []*ProxyReward

func (ts ProxyRewards) AsProxyIdMap() map[string]*ProxyReward {

	rsp := make(map[string]*ProxyReward, len(ts))

	for _, t := range ts {
		rsp[t.ProxyId] = t
	}

	return rsp

}

func (ts ProxyRewards) Ids() ([]string, []string, []string) {

	tmp1 := mapset.NewSet[string]()
	tmp2 := mapset.NewSet[string]()
	tmp3 := mapset.NewSet[string]()
	for _, t := range ts {
		tmp1.Add(t.ProxyId)
		tmp2.Add(t.ProxyUserId)
		tmp3.Add(t.Id)
	}

	return tmp1.ToSlice(), tmp2.ToSlice(), tmp3.ToSlice()
}

func (t *ProxyReward) UpdatePayed(ctx context.Context, tx *pg.Tx, storeId, id string) (bool, error) {

	u, err := tx.Model(t).Context(ctx).
		Where("store_id = ?", storeId).
		Where("id = ?", id).
		WhereIn("status in (?)", enum.PayableProxyRewardStatus).
		Set("status = ? ", enum.ProxyRewardStatus_Payed.Value).
		Update()
	if err != nil {
		return false, err
	}

	return u.RowsReturned() > 0, nil
}

func (t *ProxyReward) UpdateConfirmed(ctx context.Context, month string) error {

	_, err := comp.SDK().Postgres().Model(t).Context(ctx).
		Where("month = ?", month).
		Set("status = ? ", enum.ProxyRewardStatus_Confirmed.Value).
		Update()
	if err != nil {
		return err
	}

	return nil
}

func (t *ProxyReward) AggregateRewards(ctx context.Context, proxyIds []string) (ProxyRewards, error) {

	if len(proxyIds) == 0 {
		return nil, nil
	}

	var tmp ProxyRewards

	err := comp.SDK().Postgres().Model(t).Context(ctx).
		ColumnExpr("proxy_id").
		ColumnExpr("sum(user_count) as user_count").
		ColumnExpr("sum(order_count) as order_count").
		ColumnExpr("sum(order_amount) as order_amount").
		ColumnExpr("sum(reward_amount) as reward_amount").
		//Where("store_id = ?", storeId).
		WhereIn("proxy_id in (?)", proxyIds).
		GroupExpr("proxy_id").
		Select(&tmp)
	if err != nil {
		return nil, err
	}

	return tmp, nil
}

func (t *ProxyReward) AggregateMonthRewards(ctx context.Context, month string) (ProxyRewards, error) {

	var rewards ProxyRewards

	err := comp.SDK().Postgres().Model((*Order)(nil)).Context(ctx).
		ColumnExpr("o.store_id as store_id").
		ColumnExpr("u.proxy_id as proxy_id").
		ColumnExpr("count(distinct u.user_id) as user_count").
		ColumnExpr("count(1) as order_count").
		ColumnExpr("sum(o.amount) as order_amount").
		Join("left join t_proxy_users u on o.user_id = u.user_id and u.proxy_id is not null and o.store_id = u.store_id").
		GroupExpr("o.store_id, u.proxy_id").
		WhereIn("o.status in (?)", enum.ProxyRewardableStatus).
		Where("to_char(o.created_at, 'yyyy-mm') = ?", month).
		Where("u.proxy_id is not null and u.proxy_id != ''").
		Select(&rewards)

	if err != nil {
		return nil, err
	}

	return rewards, nil
}

func (t *ProxyReward) Upsert(ctx context.Context) (*ProxyReward, error) {

	_, err := comp.SDK().Postgres().Model(t).Context(ctx).
		OnConflict("(month, proxy_id) do nothing").
		OnConflict("(month, proxy_id) do update").
		Set("user_count = ?", t.UserCount).
		Set("order_count = ?", t.OrderCount).
		Set("order_amount = ?", t.OrderAmount).
		Set("reward_rate = ?", t.RewardRate).
		Set("reward_amount = ?", t.RewardAmount).
		Insert()

	if err != nil {
		return nil, err
	}

	return t, nil
}

type ListProxyRewardsParams struct {
	ProxyId    string
	StoreId    string
	Month      string
	Status     string
	SyncPayed  int
	MStatus    []string
	Page, Size int64
}

func (t *ProxyReward) List(ctx context.Context, params ListProxyRewardsParams) (ProxyRewards, int, error) {

	var tmp ProxyRewards
	q := comp.SDK().Postgres().Model(&tmp).Context(ctx)

	if params.ProxyId != "" {
		q = q.Where("proxy_id = ?", params.ProxyId)
	}

	if params.StoreId != "" {
		q = q.Where("store_id = ?", params.StoreId)
	}
	if params.Month != "" {
		q = q.Where("month = ?", params.Month)
	}
	if params.Status != "" {
		q = q.Where("status = ?", params.Status)
	}
	if len(params.MStatus) > 0 {
		q = q.WhereIn("status in (?)", params.MStatus)
	}

	if params.SyncPayed > 0 {
		q = q.Where("sync_payed = ?", params.SyncPayed)

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

	return tmp, count, nil
}

//func (t *ProxyReward) UpdateToClean(ctx context.Context, id string) (bool, error) {
//	update, err := comp.SDK().Postgres().Model(t).Context(ctx).
//		Set("clean = ?", enum.Clean).
//		Where("clean = ?", enum.NotClean).
//		Where("id = ?", id).
//		Update()
//	if err != nil {
//		return false, err
//	}
//
//	return update.RowsAffected() > 0, nil
//}

func (t *ProxyReward) UpdateToSynced(ctx context.Context, ids []string, field string) (bool, error) {

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
