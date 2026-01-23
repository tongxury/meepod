package service

import (
	"context"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/types"
	"github.com/go-pg/pg/v10"
)

type ProxyService struct {
}

func (t *ProxyService) ListProxies(ctx context.Context, page, size int64) (types.Proxies, int64, error) {

	dbProxies, total, err := new(db.Proxy).List(ctx, db.ListProxyParams{
		Page: page, Size: size,
	})
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	users, err := t.Assemble(ctx, dbProxies)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	return users, total, nil

}

func (t *ProxyService) FindById(ctx context.Context, id string) (*types.Proxy, error) {

	dbProxy, err := new(db.Proxy).FindById(ctx, id)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	if dbProxy == nil {
		return nil, nil
	}

	proxies, err := t.Assemble(ctx, db.Proxies{dbProxy})
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return proxies[0], nil
}

func (t *ProxyService) FindByUserId(ctx context.Context, storeId, userId string) (*types.Proxy, error) {

	dbProxy, err := new(db.Proxy).FindByUserId(ctx, storeId, userId)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	if dbProxy == nil {
		return nil, nil
	}

	proxies, err := t.Assemble(ctx, db.Proxies{dbProxy})
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return proxies[0], nil
}

func (t *ProxyService) UpdateProxy(ctx context.Context, userId, field string, value string) error {

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		_, err := new(db.Proxy).Update(ctx, tx, userId, field, value)
		return err
	})

	return xerror.Wrap(err)
}

func (t *ProxyService) Assemble(ctx context.Context, proxies db.Proxies) (types.Proxies, error) {

	ids, userIds := proxies.Ids()

	dbUsers, err := new(db.User).ListByIds(ctx, userIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	usersMap := dbUsers.AsMap()

	var tmp types.Proxies

	rewards, err := new(db.ProxyReward).AggregateRewards(ctx, ids)
	if err != nil {
		return nil, err
	}
	rewardsMap := rewards.AsProxyIdMap()

	for _, x := range proxies {
		y := types.FromDbProxy(x, usersMap[x.UserId])

		reward := rewardsMap[x.Id]
		if reward != nil {
			y.UserCount = reward.UserCount
			y.OrderCount = reward.OrderCount
			y.OrderAmount = reward.OrderAmount
			y.RewardAmount = reward.RewardAmount
		}

		tmp = append(tmp, y)
	}

	return tmp, nil

}
