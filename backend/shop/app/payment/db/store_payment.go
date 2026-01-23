package db

import (
	"context"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-pg/pg/v10"
	"time"
)

type StorePayment struct {
	tableName   struct{} `pg:"p_store_payments"`
	Id          string
	StoreId     string
	BizId       string
	BizCategory string
	Category    string
	Amount      float64
	CreatedAt   time.Time
	Status      string
	Extra       StorePaymentExtra
}

type StorePaymentExtra struct {
	FromStoreId string `json:"from_store_id"`
}

type StorePayments []*StorePayment

func (ts StorePayments) Ids() ([]string, []string) {
	tmp1 := mapset.NewSet[string]()
	tmp2 := mapset.NewSet[string]()
	for _, t := range ts {
		tmp1.Add(t.Id)
		tmp2.Add(t.StoreId)
	}

	return tmp1.ToSlice(), tmp2.ToSlice()
}

type ListStorePaymentsParams struct {
	StoreId, Month string
	Page, Size     int64
}

func (t *StorePayment) List(ctx context.Context, params ListStorePaymentsParams) (StorePayments, int64, error) {

	var tmp StorePayments
	q := comp.SDK().Postgres().Model(&tmp).Context(ctx)

	if params.StoreId != "" {
		q = q.Where("store_id = ?", params.StoreId)
	}

	if params.Month != "" {
		q = q.Where("to_char(created_at, 'yyyy-mm') = ?", params.Month)
	}

	if params.Size > 0 {
		page := params.Page
		if page <= 0 {
			page = 1
		}

		q = q.Limit(int(params.Size)).Offset(int((page - 1) * params.Size))
	}

	count, err := q.OrderExpr("id desc").
		SelectAndCount()

	if err != nil {
		return nil, 0, err
	}

	return tmp, int64(count), nil
}

func (t *StorePayment) Insert(ctx context.Context, tx *pg.Tx) error {

	insert, err := tx.Model(t).Context(ctx).
		OnConflict("(biz_category, biz_id, store_id) where biz_id != '' do nothing").
		Insert()

	if err != nil {
		return err
	}

	if insert.RowsAffected() > 0 {
		update, err := tx.Model((*StoreWallet)(nil)).Context(ctx).
			Where("store_id = ?", t.StoreId).
			Set("balance = balance + ?", t.Amount).Update()

		if err != nil {
			return xerror.Wrap(err)
		}

		if update.RowsAffected() == 0 {
			dbWallet := &StoreWallet{
				StoreId: t.StoreId,
				Balance: t.Amount,
			}

			_, err = tx.Model(dbWallet).
				OnConflict("(store_id) do nothing").
				Insert()
			if err != nil {
				return err
			}

		}
	}

	return nil
}
