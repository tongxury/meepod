package db

import (
	"context"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/core/enum"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-pg/pg/v10"
	"time"
)

type Payment struct {
	tableName   struct{} `pg:"p_user_payments"`
	Id          string
	UserId      string
	StoreId     string
	BizId       string
	BizCategory string
	Category    string
	Amount      float64
	CreatedAt   time.Time
	Status      string
	Extra       PaymentExtra
}

type PaymentExtra struct {
	Remark string
}

type Payments []*Payment

func (ts Payments) Ids() ([]string, []string, []string) {
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

func (t *Payment) List(ctx context.Context, userId, storeId string, page, size int64) (Payments, int64, error) {

	var tmp Payments
	q := comp.SDK().Postgres().Model(&tmp).Context(ctx)

	if userId != "" {
		q = q.Where("user_id = ?", userId)
	}

	if storeId != "" {
		q = q.Where("store_id = ?", storeId)
	}

	if page > 0 && size > 0 {
		q = q.Limit(int(size)).Offset(int((page - 1) * size))
	}

	count, err := q.OrderExpr("id desc").
		SelectAndCount()

	if err != nil {
		return nil, 0, err
	}

	return tmp, int64(count), nil
}

func (t *Payment) Insert(ctx context.Context, tx *pg.Tx) error {

	insert, err := tx.Model(t).Context(ctx).
		OnConflict("(biz_category, biz_id, store_id) where biz_id != '' do nothing").
		Insert()

	if err != nil {
		return err
	}

	if insert.RowsAffected() > 0 {
		_, err := tx.Model((*Account)(nil)).Context(ctx).
			Where("user_id = ?", t.UserId).
			Where("store_id = ?", t.StoreId).
			Set("balance = balance - ?", t.Amount).Update()

		if err != nil {
			return xerror.Wrap(err)
		}
	}

	return nil
}

func (t *Payment) Revert(ctx context.Context, tx *pg.Tx, orderId string, bizCategory string) (bool, error) {

	var tmp Payments
	err := tx.Model(&tmp).Context(ctx).
		Where("biz_id = ?", orderId).
		Where("biz_category = ?", bizCategory).
		WhereIn("status in (?)", enum.CanRevertPayStatus).
		Select()

	if err != nil {
		return false, err
	}

	if len(tmp) == 0 {
		return false, nil
	}

	update, err := tx.Model(t).Context(ctx).
		Where("biz_id = ?", orderId).
		Where("biz_category = ?", bizCategory).
		WhereIn("status in (?)", enum.CanRevertPayStatus).
		Set("status = ?", enum.PaymentStatus_Reverted.Value).
		Update()
	if err != nil {
		return false, err
	}

	if update.RowsAffected() > 0 {
		_, err := tx.Model((*Account)(nil)).Context(ctx).
			Where("user_id = ?", tmp[0].UserId).
			Where("store_id = ?", tmp[0].StoreId).
			Set("balance = balance + ?", tmp[0].Amount).Update()

		if err != nil {
			return false, err
		}
	}

	return true, nil
}
