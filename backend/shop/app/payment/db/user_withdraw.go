package db

import (
	"context"
	"fmt"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/shop/core/enum"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-pg/pg/v10"
	"time"
)

type Withdraw struct {
	tableName struct{} `pg:"p_user_withdraws"`
	Id        string
	UserId    string
	StoreId   string
	Amount    float64
	CreatedAt time.Time
	Status    string
	Extra     WithdrawExtra
}

type WithdrawExtra struct {
	Remark string `json:"remark"`
	Image  string `json:"image"`
}

type Withdraws []*Withdraw

func (ts Withdraws) Ids() ([]string, []string, []string) {
	tmp1 := mapset.NewSet[string]()
	tmp2 := mapset.NewSet[string]()
	tmp3 := mapset.NewSet[string]()
	for _, t := range ts {
		tmp1.Add(t.Id)
		tmp2.Add(t.StoreId)
		tmp3.Add(t.UserId)
	}

	return tmp1.ToSlice(), tmp2.ToSlice(), tmp3.ToSlice()
}

func (t *Withdraw) Insert(ctx context.Context, tx *pg.Tx) (bool, error) {

	insert, err := tx.Model(t).Context(ctx).Insert()
	if err != nil {
		return false, err
	}

	if err != nil {
		return false, err
	}

	return insert.RowsAffected() > 0, nil
}

func (t *Withdraw) UpdateExtra(ctx context.Context, tx *pg.Tx, id, field, value string) (bool, error) {

	update, err := tx.Model(t).Context(ctx).
		Where("id = ?", id).
		Set("extra =  jsonb_set(extra, ?, ?)", fmt.Sprintf("{%s}", field), fmt.Sprintf("\"%s\"", value)).
		Update()
	if err != nil {
		return false, err
	}

	return update.RowsAffected() > 0, nil
}

func (t *Withdraw) UpdateToAccepted(ctx context.Context, tx *pg.Tx, id, image string) (bool, error) {

	update, err := tx.Model(t).Context(ctx).
		Where("id = ?", id).
		//Where("user_id = ?", userId).
		WhereIn("status in (?)", enum.AcceptableWithdrawStatus).
		Set("status = ?", enum.WithdrawStatus_Accepted.Value).
		Set("extra = ?", conv.S2J(WithdrawExtra{Image: image})).
		Update()
	if err != nil {
		return false, err
	}

	return update.RowsAffected() > 0, nil
}

func (t *Withdraw) UpdateToRejected(ctx context.Context, tx *pg.Tx, id, remark string) (bool, float64, error) {

	update, err := tx.Model(t).Context(ctx).
		Where("id = ?", id).
		//Where("user_id = ?", userId).
		WhereIn("status in (?)", enum.AcceptableWithdrawStatus).
		Set("status = ?", enum.WithdrawStatus_Rejected.Value).
		Set("extra =  jsonb_set(extra, ?, ?)", "{remark}", fmt.Sprintf("\"%s\"", remark)).

		//Set("extra = ?", WithdrawExtra{Remark: remark}).
		Update()
	if err != nil {
		return false, 0, err
	}

	if err != nil {
		return false, 0, err
	}

	var tmp Withdraws
	err = tx.Model(&tmp).Context(ctx).Where("id = ?", id).Select()
	if err != nil {
		return false, 0, err
	}

	if len(tmp) == 0 {
		return false, 0, nil
	}

	return update.RowsAffected() > 0, tmp[0].Amount, nil
}

func (t *Withdraw) UpdateToCanceled(ctx context.Context, tx *pg.Tx, orderId, userId string) (bool, float64, error) {

	update, err := tx.Model(t).Context(ctx).
		Where("id = ?", orderId).
		Where("user_id = ?", userId).
		WhereIn("status in (?)", enum.CancelableWithdrawStatus).
		Set("status = ?", enum.WithdrawStatus_Canceled.Value).
		Update()

	if err != nil {
		return false, 0, err
	}

	var tmp Withdraws
	err = tx.Model(&tmp).Context(ctx).Where("id = ?", orderId).Select()
	if err != nil {
		return false, 0, err
	}

	if len(tmp) == 0 {
		return false, 0, nil
	}

	return update.RowsAffected() > 0, tmp[0].Amount, nil
}

func (t *Withdraw) FindById(ctx context.Context, id string) (*Withdraw, error) {
	var orders Withdraws
	err := comp.SDK().Postgres().Model(&orders).Context(ctx).
		Where("id = ?", id).
		Select()
	if err != nil {
		return nil, err
	}
	if len(orders) == 0 {
		return nil, nil
	}

	return orders[0], nil
}

func (t *Withdraw) RequireByIdAndStoreId(ctx context.Context, id, storeId string) (*Withdraw, error) {
	var orders Withdraws
	err := comp.SDK().Postgres().Model(&orders).Context(ctx).
		Where("id = ?", id).
		Where("store_id = ?", storeId).
		Select()
	if err != nil {
		return nil, err
	}
	if len(orders) == 0 {
		return nil, fmt.Errorf("no order found by id: %s, %s", id, storeId)
	}

	return orders[0], nil
}

func (t *Withdraw) ExistsByStatus(ctx context.Context, userId, storeId, status string) (bool, error) {
	var orders Withdraws
	err := comp.SDK().Postgres().Model(&orders).Context(ctx).
		Where("user_id = ?", userId).
		Where("store_id = ?", storeId).
		Where("status = ?", status).
		Select()
	if err != nil {
		return false, err
	}

	return len(orders) > 0, nil
}

func (t *Withdraw) List(ctx context.Context, userId, storeId string, status []string, page, size int64) (Withdraws, int64, error) {

	var orders Withdraws
	q := comp.SDK().Postgres().Model(&orders).Context(ctx)

	if userId != "" {
		q = q.Where("user_id = ?", userId)
	}

	if storeId != "" {
		q = q.Where("store_id = ?", storeId)
	}

	if len(status) > 0 {
		q = q.WhereIn("status in (?)", status)
	}

	if page > 0 && size > 0 {
		q = q.Limit(int(size)).Offset(int((page - 1) * size))
	}

	count, err := q.OrderExpr("id desc").
		SelectAndCount()

	if err != nil {
		return nil, 0, err
	}

	return orders, int64(count), nil
}

func (t *Withdraw) CountByStatus(ctx context.Context, storeId string, status []string) (int64, error) {
	total, err := comp.SDK().Postgres().WithContext(ctx).Model(t).
		WhereIn("status in (?)", status).
		Where("store_id = ?", storeId).
		Count()

	if err != nil {
		return 0, err
	}

	return int64(total), nil
}
