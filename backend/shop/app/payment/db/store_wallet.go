package db

import (
	"context"
	"fmt"
	"gitee.com/meepo/backend/kit/components/comp"
	"time"
)

type StoreWallet struct {
	tableName struct{} `pg:"p_store_wallets"`
	StoreId   string
	Balance   float64
	CreatedAt time.Time
	Status    string
}

type StoreWallets []*StoreWallet

func (ts StoreWallets) AsMap() map[string]*StoreWallet {

	rsp := make(map[string]*StoreWallet, len(ts))

	for _, t := range ts {
		rsp[t.StoreId] = t
	}

	return rsp
}

func (t *StoreWallet) FindByStoreId(ctx context.Context, storeId string) (*StoreWallet, error) {

	var tmp StoreWallets
	err := comp.SDK().Postgres().Model(&tmp).Context(ctx).
		Where("store_id = ?", storeId).
		Select()
	if err != nil {
		return nil, err
	}

	if len(tmp) == 0 {
		return nil, nil
	}

	return tmp[0], nil
}

func (t *StoreWallet) RequireByStoreId(ctx context.Context, storeId string) (*StoreWallet, error) {

	var tmp StoreWallets
	err := comp.SDK().Postgres().Model(&tmp).Context(ctx).
		Where("store_id = ?", storeId).
		Select()
	if err != nil {
		return nil, err
	}

	if len(tmp) == 0 {
		return nil, fmt.Errorf("no store wallet found by id :%s", storeId)
	}

	return tmp[0], nil
}

type ListStoreWalletsParams struct {
	StoreIds   []string
	Page, Size int64
}

func (t *StoreWallet) List(ctx context.Context, params ListStoreWalletsParams) (StoreWallets, int64, error) {

	var tmp StoreWallets
	q := comp.SDK().Postgres().Model(&tmp).Context(ctx)

	if len(params.StoreIds) > 0 {
		q.WhereIn("store_id in (?)", params.StoreIds)
	}

	if params.Size > 0 {
		page := params.Page
		if page <= 0 {
			page = 1
		}

		q = q.Limit(int(params.Size)).Offset(int((page - 1) * params.Size))
	}

	count, err := q.SelectAndCount()

	if err != nil {
		return nil, 0, err
	}

	return tmp, int64(count), nil
}
