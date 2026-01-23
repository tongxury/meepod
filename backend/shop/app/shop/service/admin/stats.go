package adminserivce

import (
	"context"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/core/db"
)

type StatsService struct {
}

func (t *StatsService) GetStats(ctx context.Context) (any, error) {

	pg := comp.SDK().Postgres()

	storeCount, err := pg.Model((*db.Store)(nil)).Context(ctx).Count()
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	userCount, err := pg.Model((*db.User)(nil)).Context(ctx).Count()
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	orderCount, err := pg.Model((*db.Order)(nil)).Context(ctx).Count()
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return map[string]int{
		"store_count": storeCount,
		"user_count":  userCount,
		"order_count": orderCount,
	}, nil
}
