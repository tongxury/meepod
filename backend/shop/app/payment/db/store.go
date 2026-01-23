package db

import (
	"context"
	"fmt"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/third/alipay"
	"gitee.com/meepo/backend/kit/services/util/pgd"
	"gitee.com/meepo/backend/shop/core/enum"
	mapset "github.com/deckarep/golang-set/v2"
	"time"
)

type Store struct {
	tableName struct{} `pg:"p_stores"`
	StoreId   string
	Status    string
	Alipay    *Alipay
	Xinsh     *Xinsh
	CreatedAt time.Time
}

type Xinsh struct {
	MerchantNo string `json:"merchant_no,omitempty"`
	RequestId  string `json:"request_id,omitempty"`
	State      string `json:"state,omitempty"`
	Status     string `json:"status,omitempty"`
}

type Alipay struct {
	AppId     string            `json:"app_id"`
	AuthToken *alipay.AuthToken `json:"auth_token"`
}

type Stores []*Store

func (ts Stores) Ids() []string {

	tmp1 := mapset.NewSet[string]()

	for _, t := range ts {
		tmp1.Add(t.StoreId)
	}

	return tmp1.ToSlice()
}

func (t *Store) Insert(ctx context.Context) (bool, error) {

	insert, err := comp.SDK().Postgres().Model(t).Context(ctx).
		OnConflict("(store_id) do nothing").
		Insert()
	if err != nil {
		return false, err
	}

	return insert.RowsAffected() > 0, nil
}

func (t *Store) RequireByStoreId(ctx context.Context, storeId string) (*Store, error) {

	var tmp Stores
	err := comp.SDK().Postgres().Model(&tmp).Context(ctx).
		Where("store_id = ?", storeId).
		Select()
	if err != nil {
		return nil, err
	}

	if len(tmp) == 0 {
		return nil, fmt.Errorf("no store found by id: %s", storeId)
	}

	return tmp[0], nil
}

func (t *Store) UpdateXinsh(ctx context.Context, storeId string, xinsh Xinsh) error {
	_, err := comp.SDK().Postgres().Model(t).Context(ctx).
		Where("store_id = ?", storeId).Set("xinsh = ?", xinsh).Update()
	if err != nil {
		return err
	}
	return nil
}

func (t *Store) UpdateXinshField(ctx context.Context, storeId string, field string, value any) error {

	q := comp.SDK().Postgres().Model(t).Context(ctx).
		Where("store_id = ?", storeId)

	q = pgd.SetJSONField(q, "xinsh", field, value)
	_, err := q.Update()
	if err != nil {
		return err
	}

	return nil
}

func (t *Store) UpdateAuthToken(ctx context.Context, authAppId string, authToken alipay.AuthToken) error {

	_, err := comp.SDK().Postgres().Model(t).Context(ctx).
		Where("app_id = ?", authAppId).
		Set("auth_token = ?", authToken).
		Set("status = ?", enum.PaymentStoreStatus_Authed.Value).
		Update()
	if err != nil {
		return err
	}

	return nil
}

type ListStoresParams struct {
	StoreId     string
	XinshStatus string
	Page, Size  int64
}

func (t *Store) List(ctx context.Context, params ListStoresParams) (Stores, int64, error) {

	var orders Stores
	q := comp.SDK().Postgres().Model(&orders).Context(ctx)

	if params.StoreId != "" {
		q = q.Where("store_id = ?", params.StoreId)
	}

	if params.XinshStatus != "" {
		q = q.Where("xinsh ->> 'status' = ?", params.XinshStatus)
	}

	if params.Size > 0 {
		page := params.Page
		if page <= 0 {
			page = 1
		}

		q = q.Limit(int(params.Size)).Offset(int((page - 1) * params.Size))
	}

	count, err := q.OrderExpr("created_at desc").
		SelectAndCount()

	if err != nil {
		return nil, 0, err
	}

	return orders, int64(count), nil
}
