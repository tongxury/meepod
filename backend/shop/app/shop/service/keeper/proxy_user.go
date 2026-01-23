package keeperservice

import (
	"context"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/app/shop/service"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/types"
)

type ProxyUserService struct {
	service.ProxyUserService
}

func (t *ProxyUserService) ListProxyUsers(ctx context.Context, proxyId, storeId string, page, size int64) (types.ProxyUsers, int, error) {

	dbProxyUsers, total, err := new(db.ProxyUser).List(ctx, db.ListProxyUsersParams{
		ProxyId: proxyId, StoreId: storeId, Page: page, Size: size,
	})
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	users, err := t.Assemble(ctx, storeId, dbProxyUsers)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	return users, total, nil
}
