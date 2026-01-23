package db

import (
	"context"
	"fmt"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/core/enum"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-pg/pg/v10"
	"time"
)

type Reward struct {
	tableName   struct{} `pg:"t_user_rewards"`
	Id          string
	BizId       string
	BizCategory string
	UserId      string
	StoreId     string
	Amount      float64
	CreatedAt   time.Time
	Status      string
	Extra       RewardExtra
}

type RewardExtra struct {
	Summary      string  `json:"summary"`
	TotalCount   int64   `json:"total_count"`
	TotalAmount  float64 `json:"total_amount"`
	Multiple     int64   `json:"multiple"`
	OrderGroupId string  `json:"order_group_id,omitempty"`
	TotalVolume  int64   `json:"total_volume,omitempty"`
	Volume       int64   `json:"volume,omitempty"`
}

type Rewards []*Reward

func (ts Rewards) Ids() ([]string, []string, []string) {

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

func (t *Reward) Insert(ctx context.Context, tx *pg.Tx) (bool, error) {

	insert, err := tx.Model(t).Context(ctx).
		OnConflict("(biz_id, biz_category) DO NOTHING").
		Insert()
	if err != nil {
		return false, err
	}

	return insert.RowsAffected() > 0, nil
}

func (ts Rewards) InsertBatch(ctx context.Context, tx *pg.Tx) (bool, error) {

	if len(ts) == 0 {
		return false, nil
	}

	insert, err := tx.Model(&ts).Context(ctx).
		OnConflict("(biz_id, biz_category) DO NOTHING").
		Insert()
	if err != nil {
		return false, err
	}

	return insert.RowsAffected() >= len(ts), nil
}

type ListRewardsParams struct {
	StoreId    string
	MStatus    []string
	SyncPay    int
	Page, Size int64
}

func (t *Reward) List(ctx context.Context, params ListRewardsParams) (Rewards, int64, error) {

	var rsp Rewards
	q := comp.SDK().Postgres().Model(&rsp).Context(ctx)
	if params.StoreId != "" {
		q = q.Where("store_id = ?", params.StoreId)
	}
	if params.SyncPay > 0 {
		q = q.Where("sync_pay = ?", params.SyncPay)
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
	q = q.OrderExpr("id desc")
	count, err := q.SelectAndCount()

	return rsp, int64(count), err

}

func (t *Reward) RequireByIdAndStoreId(ctx context.Context, id, storeId string) (*Reward, error) {

	var tmp Rewards

	err := comp.SDK().Postgres().Model(&tmp).Context(ctx).
		Where("id = ?", id).
		Where("store_id = ?", storeId).
		Select()

	if err != nil {
		return nil, err
	}

	if len(tmp) == 0 {
		return nil, fmt.Errorf("no account found by id, storeId: %s,%s", id, storeId)
	}

	return tmp[0], nil
}

func (t *Reward) UpdateToRewarded(ctx context.Context, tx *pg.Tx, id string) (bool, error) {

	reward := Reward{Id: id}

	updated, err := tx.Model(&reward).Context(ctx).
		Set("status = ?", enum.RewardStatus_Rewarded.Value).
		//Set("extra =  jsonb_set(extra, ?, ?)", "{rejectReasonId}", reasonId).
		Where("id = ?", id).
		WhereIn("status in (?)", enum.RewardableRewardStatus).
		Update()

	if err != nil {
		return false, xerror.Wrap(err)
	}

	return updated.RowsAffected() > 0, nil
}

func (t *Reward) UpdateToRejected(ctx context.Context, tx *pg.Tx, id, reason string) (bool, error) {

	reward := Reward{Id: id}

	updated, err := tx.Model(&reward).Context(ctx).
		Set("status = ?", enum.RewardStatus_Rejected.Value).
		Set("extra =  jsonb_set(extra, ?, ?)", "{reason}", reason).
		Where("id = ?", id).
		WhereIn("status in (?)", enum.RejectableRewardStatus).
		Update()

	if err != nil {
		return false, xerror.Wrap(err)
	}

	return updated.RowsAffected() > 0, nil
}

func (t *Reward) CountByStatus(ctx context.Context, storeId string, status []string) (int64, error) {
	total, err := comp.SDK().Postgres().WithContext(ctx).Model(t).
		WhereIn("status in (?)", status).
		Where("store_id = ?", storeId).
		Count()

	if err != nil {
		return 0, err
	}

	return int64(total), nil
}

func (t *Reward) UpdateToSynced(ctx context.Context, ids []string, field string) (bool, error) {

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
