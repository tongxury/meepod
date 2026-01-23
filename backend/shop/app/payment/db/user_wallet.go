package db

import (
	"context"
	"fmt"
	"gitee.com/meepo/backend/kit/components/comp"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-pg/pg/v10"
	"time"
)

type Account struct {
	tableName struct{} `pg:"p_user_wallets"`
	Id        string
	UserId    string
	StoreId   string
	Balance   float64
	CreatedAt time.Time
	Status    string
}

type Accounts []*Account

func (ts Accounts) Ids() ([]string, []string, []string) {

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

func (t *Account) RequireByIdAndStoreId(ctx context.Context, id, storeId string) (*Account, error) {

	var tmp Accounts

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

func (t *Account) RequireByUserIdAndStoreId(ctx context.Context, userId, storeId string) (*Account, error) {

	acc := Account{
		UserId:  userId,
		StoreId: storeId,
	}

	_, err := comp.SDK().Postgres().Model(&acc).Context(ctx).
		Where("user_id = ?", userId).
		Where("store_id = ?", storeId).
		OnConflict("(user_id, store_id) do nothing").
		SelectOrInsert()

	if err != nil {
		return nil, err
	}

	return &acc, nil
}

func (t *Account) List(ctx context.Context, storeId string, page, size int64) (Accounts, int64, error) {

	var accounts Accounts
	q := comp.SDK().Postgres().Model(&accounts).Context(ctx)
	if storeId != "" {
		q = q.Where("store_id = ?", storeId)
	}

	if page > 0 && size > 0 {
		q = q.Limit(int(size)).Offset(int((page - 1) * size))
	}
	count, err := q.SelectAndCount()

	return accounts, int64(count), err

}

func (t *Account) GetStoreSummary(ctx context.Context, storeId string) (int64, float64, error) {

	var rsp struct {
		UserCount    int64
		TotalBalance float64
	}

	count, err := comp.SDK().Postgres().Model((*Account)(nil)).Context(ctx).
		Where("store_id = ?", storeId).Count()

	if err != nil {
		return 0, 0, err
	}

	if count == 0 {
		return 0, 0, nil
	}

	err = comp.SDK().Postgres().Model((*Account)(nil)).Context(ctx).
		ColumnExpr("count(user_id) as user_count").
		ColumnExpr("sum(balance) as total_balance").
		Where("store_id = ?", storeId).
		GroupExpr("store_id").
		Select(&rsp)

	if err != nil {
		return 0, 0, err
	}

	return rsp.UserCount, rsp.TotalBalance, nil
}

func (t *Account) Incr(ctx context.Context, tx *pg.Tx, userId, storeId string, amount float64) error {

	acc := &Account{
		UserId:  userId,
		StoreId: storeId,
		Balance: amount,
	}

	_, err := tx.Model(acc).Context(ctx).
		OnConflict("(user_id, store_id) do update").
		Set("balance = account.balance + ?", amount).
		Insert()

	if err != nil {
		return err
	}

	return nil
}
