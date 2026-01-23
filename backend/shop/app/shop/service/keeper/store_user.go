package keeperservice

import (
	"context"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/app/shop/service"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/types"
	"github.com/go-pg/pg/v10"
)

type StoreUserService struct {
	service.StoreUserService
}

func (t *StoreUserService) ListStoreUsers(ctx context.Context, keeperId, storeId, phone string, page, size int64) (types.StoreUsers, int64, error) {

	dbItems, total, err := new(db.StoreUser).List(ctx, db.ListStoreUsersParams{
		Phone:   phone,
		StoreId: storeId,
		Page:    page, Size: size,
	})
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	items, err := t.Assemble(ctx, dbItems)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	return items, total, nil

}

func (t *StoreUserService) FindStoreUserById(ctx context.Context, storeId, userId string) (*types.StoreUser, error) {

	storeUsers, _, err := new(db.StoreUser).List(ctx, db.ListStoreUsersParams{
		StoreId: storeId, UserIds: []string{userId},
	})
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	if len(storeUsers) == 0 {
		return nil, nil
	}

	users, err := t.Assemble(ctx, storeUsers)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return users[0], nil
}

func (t *StoreUserService) UpdateUser(ctx context.Context, storeId, userId, field string, value string) error {

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		_, err := new(db.StoreUser).Update(ctx, tx, storeId, userId, field, value)
		return err
	})

	return xerror.Wrap(err)
}
