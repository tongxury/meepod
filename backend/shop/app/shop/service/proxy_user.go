package service

import (
	"context"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/kit/services/util/gind/errorx"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/types"
	"time"
)

type ProxyUserService struct {
}

func (t *ProxyUserService) ListProxyUsers(ctx context.Context, proxyUserId, storeId string, page, size int64) (types.ProxyUsers, int, error) {

	dbProxyUsers, total, err := new(db.ProxyUser).List(ctx, db.ListProxyUsersParams{
		ProxyUserId: proxyUserId, StoreId: storeId, Page: page, Size: size,
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

func (t *ProxyUserService) Assemble(ctx context.Context, storeId string, items db.ProxyUsers) (types.ProxyUsers, error) {

	userIds := items.Ids()

	dbUsers, err := new(db.User).ListByIds(ctx, userIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	dbUsersMap := dbUsers.AsMap()

	orderAggs, err := new(db.Order).AggregateByUserIds(ctx, storeId, userIds)
	if err != nil {
		return nil, err
	}
	orderAggsMap := orderAggs.AsMap()

	var proxyUsers types.ProxyUsers
	for _, x := range items {

		y := types.ProxyUser{
			User:        types.FromDbUser(dbUsersMap[x.UserId]),
			CreatedAt:   x.CreatedAt.Format(time.DateOnly),
			CreatedAtTs: x.CreatedAt.Unix(),
			OrderCount:  0,
			OrderAmount: 0,
		}

		if agg, found := orderAggsMap[x.UserId]; found {
			y.OrderCount = agg.OrderCount
			y.OrderAmount = agg.OrderAmount
		}

		proxyUsers = append(proxyUsers, &y)
	}

	return proxyUsers, nil
}

func (t *ProxyUserService) AddUser(ctx context.Context, storeId, proxyId, userId string) error {

	proxy, err := new(db.Proxy).FindById(ctx, proxyId)
	if err != nil {
		return xerror.Wrap(err)
	}

	if proxy == nil {
		return errorx.UserMessage("未知推荐人")
	}

	if proxy.UserId == userId {
		return errorx.UserMessage("不能添加自己为推广用户")
	}

	dbProxyUser := &db.ProxyUser{
		ProxyId:     proxyId,
		ProxyUserId: proxy.UserId,
		StoreId:     proxy.StoreId,
		UserId:      userId,
	}

	_, err = dbProxyUser.CreateNX(ctx)
	if err != nil {
		return xerror.Wrap(err)
	}

	return nil
}
