package db

import (
	"context"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-pg/pg/v10"
	"time"
)

type CoStorePayment struct {
	tableName   struct{} `pg:"p_co_store_payments"`
	Id          string
	StoreId     string
	CoStoreId   string
	BizId       string
	BizCategory string
	Category    string
	Amount      float64
	CreatedAt   time.Time
	Status      string
	Extra       CoStorePaymentExtra
}

type CoStorePaymentExtra struct {
	ProofImage string `json:"proof_image"`
}

type CoStorePayments []*CoStorePayment

func (ts CoStorePayments) Ids() ([]string, []string, []string) {
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

type ListCoStorePaymentsParams struct {
	Month      string
	StoreId    string
	CoStoreId  string
	Category   string
	Page, Size int64
}

func (t *CoStorePayment) List(ctx context.Context, params ListCoStorePaymentsParams) (CoStorePayments, int64, error) {

	var tmp CoStorePayments
	q := comp.SDK().Postgres().Model(&tmp).Context(ctx)

	if params.Month != "" {
		q = q.Where("to_char(created_at, 'yyyy-mm') = ?", params.Month)
	}

	if params.StoreId != "" {
		q = q.Where("store_id = ?", params.StoreId)
	}

	if params.CoStoreId != "" {
		q = q.Where("co_store_id = ?", params.CoStoreId)
	}

	if params.Category != "" {
		q = q.Where("category = ?", params.Category)
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

func (t *CoStorePayment) Insert(ctx context.Context, tx *pg.Tx) error {

	insert, err := tx.Model(t).Context(ctx).
		OnConflict("(category, biz_category, biz_id, store_id, co_store_id) where biz_id != '' do nothing").
		Insert()

	if err != nil {
		return err
	}

	if insert.RowsAffected() > 0 {
		update, err := tx.Model((*CoStoreWallet)(nil)).Context(ctx).
			Where("store_id = ?", t.StoreId).
			Where("co_store_id = ?", t.CoStoreId).
			Set("balance = balance + ?", t.Amount).Update()

		if err != nil {
			return xerror.Wrap(err)
		}

		if err != nil {
			return xerror.Wrap(err)
		}

		if update.RowsAffected() == 0 {
			dbWallet := &CoStoreWallet{
				StoreId:   t.StoreId,
				CoStoreId: t.CoStoreId,
				Balance:   t.Amount,
			}

			_, err = tx.Model(dbWallet).
				OnConflict("(store_id, co_store_id) do nothing").
				Insert()
			if err != nil {
				return err
			}

		}
	}

	return nil
}
