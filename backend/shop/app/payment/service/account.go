package service

import (
	"context"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/app/payment/db"
	coredb "gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/types"
)

type AccountService struct {
}

func (t *AccountService) RequireByIdAndStoreId(ctx context.Context, id, storeId string) (*types.Account, error) {

	account, err := new(db.Account).RequireByIdAndStoreId(ctx, id, storeId)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	accounts, err := t.AssembleAccounts(ctx, db.Accounts{account})
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return accounts[0], nil
}

func (t *AccountService) RequireByUserIdAndStoreId(ctx context.Context, userId, storeId string) (*types.Account, error) {

	account, err := new(db.Account).RequireByUserIdAndStoreId(ctx, userId, storeId)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	accounts, err := t.AssembleAccounts(ctx, db.Accounts{account})
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return accounts[0], nil
}

func (t *AccountService) AssembleAccounts(ctx context.Context, accounts db.Accounts) (types.Accounts, error) {

	var userIds []string
	_, accountUserIds, storeIds := accounts.Ids()
	userIds = append(userIds, accountUserIds...)

	dbStores, err := new(coredb.Store).ListByIds(ctx, storeIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	dbStoresTap := dbStores.AsMap()
	_, storeOwnerIds := dbStores.Ids()
	userIds = append(userIds, storeOwnerIds...)

	dbUsers, err := new(coredb.User).ListByIds(ctx, userIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	dbUsersTap := dbUsers.AsMap()

	var rsp types.Accounts
	for _, x := range accounts {

		store := dbStoresTap[x.StoreId]
		accountUser := dbUsersTap[x.UserId]
		storeOwner := dbUsersTap[store.OwnerId]

		y := types.FromDbAccount(x, accountUser, store, storeOwner)

		rsp = append(rsp, y)
	}

	return rsp, nil
}
