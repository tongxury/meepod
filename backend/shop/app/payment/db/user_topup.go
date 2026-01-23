package db

import (
	"context"
	"fmt"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/shop/core/enum"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-pg/pg/v10"
	"time"
)

type Topup struct {
	tableName   struct{} `pg:"p_user_topups"`
	Id          string
	UserId      string
	StoreId     string
	Amount      float64
	CreatedAt   time.Time
	Status      string
	Category    string
	BizId       string
	BizCategory string
	Extra       *TopupExtra
}

type TopupExtra struct {
	PayUrl       string `json:"pay_url,omitempty"`
	QrCode       string `json:"qr_code,omitempty"`
	PayMethod    string `json:"pay_method,omitempty"`
	OrderId      string `json:"order_id,omitempty"`
	Category     string `json:"category,omitempty"`
	MerchantNo   string `json:"merchant_no"`
	TradeNo      string `json:"trade_no"`
	SellerId     string `json:"seller_id"`
	BuyerId      string `json:"buyer_id"`
	BuyerLogonId string `json:"buyer_logon_id"`
}

type Topups []*Topup

func (ts Topups) Ids() ([]string, []string, []string) {
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

func (t *Topup) Insert(ctx context.Context, tx *pg.Tx) (bool, error) {

	insert, err := tx.Model(t).Context(ctx).
		OnConflict("(biz_category, biz_id, store_id)  where biz_id != '' do nothing").
		Insert()
	if err != nil {
		return false, err
	}

	return insert.RowsAffected() > 0, nil
}

func (t *Topup) CleanTimeoutTopups(ctx context.Context, timeout time.Duration) error {

	_, err := comp.SDK().Postgres().Model(t).Context(ctx).
		Where("created_at < ?", time.Now().Add(-timeout)).
		Where("status = ?", enum.TopupStatus_Submitted.Value).
		Set("status = ?", enum.TopupStatus_Timeout.Value).Update()
	if err != nil {
		return err
	}

	return nil
}

func (t *Topup) SetExtra(ctx context.Context, tx *pg.Tx, id string, extra TopupExtra) (bool, error) {

	update, err := tx.Model(t).Context(ctx).
		Where("id = ?", id).
		Set("extra = ?", extra).
		Update()
	if err != nil {
		return false, err
	}

	return update.RowsAffected() > 0, nil
}

func (t *Topup) UpdateExtra(ctx context.Context, tx *pg.Tx, id, field, value string) (bool, error) {

	update, err := tx.Model(t).Context(ctx).
		Where("id = ?", id).
		Set("extra =  jsonb_set(extra, ?, ?)", fmt.Sprintf("{%s}", field), fmt.Sprintf("\"%s\"", value)).
		Update()
	if err != nil {
		return false, err
	}

	return update.RowsAffected() > 0, nil
}

func (t *Topup) UpdateToPayed(ctx context.Context, tx *pg.Tx, topupId, userId string) (bool, error) {

	update, err := tx.Model(t).Context(ctx).
		Where("id = ?", topupId).
		Where("user_id = ?", userId).
		WhereIn("status in (?)", enum.PayableTopupStatus).
		Set("status = ?", enum.TopupStatus_Payed.Value).
		Update()
	if err != nil {
		return false, err
	}

	return update.RowsAffected() > 0, nil
}

func (t *Topup) UpdateToCanceled(ctx context.Context, tx *pg.Tx, orderId, userId string) (bool, error) {

	update, err := tx.Model(t).Context(ctx).
		Where("id = ?", orderId).
		Where("user_id = ?", userId).
		WhereIn("status in (?)", enum.CancelableTopupStatus).
		Set("status = ?", enum.TopupStatus_Canceled.Value).Update()
	if err != nil {
		return false, err
	}

	return update.RowsAffected() > 0, nil
}

func (t *Topup) FindById(ctx context.Context, id string) (*Topup, error) {
	var orders Topups
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

func (t *Topup) RequireByIdAndStoreId(ctx context.Context, id, storeId string) (*Topup, error) {
	var orders Topups
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

func (t *Topup) ExistsByStatus(ctx context.Context, userId, storeId, status string) (bool, error) {
	var orders Topups
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

type ListTopupsParams struct {
	UserId      string
	StoreId     string
	MStatus     []string
	SyncSettled int
	Page, Size  int64
}

func (t *Topup) List(ctx context.Context, params ListTopupsParams) (Topups, int64, error) {

	var orders Topups
	q := comp.SDK().Postgres().Model(&orders).Context(ctx)

	if params.UserId != "" {
		q = q.Where("user_id = ?", params.UserId)
	}

	if params.StoreId != "" {
		q = q.Where("store_id = ?", params.StoreId)
	}

	if params.SyncSettled > 0 {
		q = q.Where("sync_settled = ?", params.SyncSettled)
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

	count, err := q.OrderExpr("id desc").
		SelectAndCount()

	if err != nil {
		return nil, 0, err
	}

	return orders, int64(count), nil
}

func (t *Topup) UpdateToSynced(ctx context.Context, ids []string, field string) (bool, error) {

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
