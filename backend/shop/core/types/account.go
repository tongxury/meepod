package types

import (
	"gitee.com/meepo/backend/kit/components/sdk/helper/timed"
	"gitee.com/meepo/backend/shop/app/payment/db"
	coredb "gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
)

type AccountSummary struct {
	UserCount    int64   `json:"user_count"`
	TotalBalance float64 `json:"total_balance"`
}

type Account struct {
	Id          string    `json:"id"`
	User        *User     `json:"user"`
	Store       *Store    `json:"store"`
	Balance     float64   `json:"balance"`
	CreatedAt   string    `json:"created_at"`
	CreatedAtTs int64     `json:"created_at_ts"`
	Status      enum.Enum `json:"status"`
	Decrable    bool      `json:"decrable"`
}

type Accounts []*Account

func FromDbAccount(account *db.Account, accountUser *coredb.User, store *coredb.Store, storeOwner *coredb.User) *Account {

	rsp := Account{
		Id:          account.Id,
		User:        FromDbUser(accountUser),
		Store:       FromDbStore(store, storeOwner),
		Balance:     account.Balance,
		CreatedAtTs: account.CreatedAt.Unix(),
		CreatedAt:   timed.SmartTime(account.CreatedAt.Unix()),
		Status:      enum.AccountStatus(account.Status),
		Decrable:    account.Balance > 0,
	}

	return &rsp
}
