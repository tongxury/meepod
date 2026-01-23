package service

import (
	"context"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/types"
	"github.com/go-pg/pg/v10"
)

type StoreUserService struct {
}

func (t *StoreUserService) AddStoreUser(ctx context.Context, storeId, userId string) error {

	dbStoreUser := db.StoreUser{
		StoreId: storeId,
		UserId:  userId,
	}

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {
		_, err := dbStoreUser.Insert(ctx, tx)
		if err != nil {
			return err
		}
		return nil
	})

	return xerror.Wrap(err)

}

func (t *StoreUserService) Assemble(ctx context.Context, storeUsers db.StoreUsers) (types.StoreUsers, error) {

	if len(storeUsers) == 0 {
		return nil, nil
	}

	_, userIds := storeUsers.Ids()

	users, err := new(db.User).ListByIds(ctx, userIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	usersMap := users.AsMap()

	var rsp types.StoreUsers

	for _, x := range storeUsers {

		y := types.FromDbStoreUser(x, usersMap[x.UserId])
		rsp = append(rsp, y)
	}

	return rsp, nil
}
