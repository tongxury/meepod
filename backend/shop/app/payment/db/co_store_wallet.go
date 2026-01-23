package db

import (
	"context"
	"gitee.com/meepo/backend/kit/components/comp"
	mapset "github.com/deckarep/golang-set/v2"
	"time"
)

type CoStoreWallet struct {
	tableName struct{} `pg:"p_co_store_wallets"`
	Id        string
	StoreId   string
	CoStoreId string
	Balance   float64
	CreatedAt time.Time
	Status    string
}

type CoStoreWallets []*CoStoreWallet

func (ts CoStoreWallets) Ids() ([]string, []string, []string) {

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

func (ts CoStoreWallets) AsMap() map[string]*CoStoreWallet {
	rsp := map[string]*CoStoreWallet{}

	for _, x := range ts {
		rsp[x.StoreId+"-"+x.CoStoreId] = x
	}

	return rsp
}

//func (t *CoStoreWallet) RequireByIdAndCoStoreId(ctx context.Context, id, coStoreId string) (*Account, error) {
//
//	var tmp Accounts
//
//	err := comp.SDK().Postgres().Model(&tmp).Context(ctx).
//		Where("id = ?", id).
//		Where("co_store_id = ?", coStoreId).
//		Select()
//
//	if err != nil {
//		return nil, err
//	}
//
//	if len(tmp) == 0 {
//		return nil, fmt.Errorf("no account found by id, coStoreId: %s,%s", id, coStoreId)
//	}
//
//	return tmp[0], nil
//}

//func (t *CoStoreWallet) RequireByUserIdAndStoreId(ctx context.Context, storeId, coStoreId string) (*CoStoreWallet, error) {
//
//	acc := CoStoreWallet{
//		StoreId:   storeId,
//		CoStoreId: coStoreId,
//	}
//
//	_, err := comp.SDK().Postgres().Model(&acc).Context(ctx).
//		Where("store_id = ?", storeId).
//		Where("co_store_id = ?", coStoreId).
//		OnConflict("(store_id, co_store_id) do nothing").
//		SelectOrInsert()
//
//	if err != nil {
//		return nil, err
//	}
//
//	return &acc, nil
//}

type ListCoStoreWalletsParams struct {
	StoreId    string
	StoreIds   []string
	CoStoreIds []string
	Page, Size int64
}

func (t *CoStoreWallet) FindByStoreIdAndCoStoreId(ctx context.Context, storeId, coStoreId string) (*CoStoreWallet, error) {

	var tmp CoStoreWallets

	err := comp.SDK().Postgres().Model(&tmp).Context(ctx).
		Where("store_id = ?", storeId).
		Where("co_store_id = ?", coStoreId).
		Select()

	if err != nil {
		return nil, err
	}

	if len(tmp) == 0 {
		return nil, nil
	}
	return tmp[0], nil
}

func (t *CoStoreWallet) List(ctx context.Context, params ListCoStoreWalletsParams) (CoStoreWallets, int64, error) {

	var accounts CoStoreWallets
	q := comp.SDK().Postgres().Model(&accounts).Context(ctx)
	if params.StoreId != "" {
		q = q.Where("store_id = ?", params.StoreId)
	}

	if len(params.CoStoreIds) > 0 {
		q = q.WhereIn("co_store_id in (?)", params.CoStoreIds)
	}

	if len(params.StoreIds) > 0 {
		q = q.WhereIn("store_id in (?)", params.StoreIds)
	}

	if params.Size > 0 {
		page := params.Page
		if page <= 0 {
			page = 1
		}

		q = q.Limit(int(params.Size)).Offset(int((page - 1) * params.Size))
	}
	count, err := q.SelectAndCount()

	return accounts, int64(count), err

}

//func (t *CoStoreWallet) GetStoreSummary(ctx context.Context, storeId string) (int64, float64, error) {
//
//	var rsp struct {
//		UserCount    int64
//		TotalBalance float64
//	}
//
//	count, err := comp.SDK().Postgres().Model((*Account)(nil)).Context(ctx).
//		Where("store_id = ?", storeId).Count()
//
//	if err != nil {
//		return 0, 0, err
//	}
//
//	if count == 0 {
//		return 0, 0, nil
//	}
//
//	err = comp.SDK().Postgres().Model((*Account)(nil)).Context(ctx).
//		ColumnExpr("count(user_id) as user_count").
//		ColumnExpr("sum(balance) as total_balance").
//		Where("store_id = ?", storeId).
//		GroupExpr("store_id").
//		Select(&rsp)
//
//	if err != nil {
//		return 0, 0, err
//	}
//
//	return rsp.UserCount, rsp.TotalBalance, nil
//}

//func (t *CoStoreWallet) Incr(ctx context.Context, tx *pg.Tx, storeId, coStoreId string, amount float64) error {
//
//	acc := &CoStoreWallet{
//		StoreId:   storeId,
//		CoStoreId: coStoreId,
//		Balance:   amount,
//	}
//
//	_, err := tx.Model(acc).Context(ctx).
//		OnConflict("(store_id, co_store_id) do update").
//		Set("balance = co_store_wallet.balance + ?", amount).
//		Insert()
//
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
