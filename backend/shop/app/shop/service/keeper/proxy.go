package keeperservice

import (
	"context"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/kit/services/util/gind/errorx"
	"gitee.com/meepo/backend/shop/app/shop/service"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/types"
	"github.com/go-pg/pg/v10"
	"github.com/lithammer/shortuuid/v4"
)

type ProxyService struct {
	service.ProxyService
}

func (t *ProxyService) UpdateRewardRate(ctx context.Context, keeperId, storeId, proxyId string, rewardRate float64) error {
	if rewardRate < enum.MinProxyRewardRate || rewardRate > enum.MaxProxyRewardRate {
		return errorx.ParamErrorf("rewardRate is require between %f to %f", enum.MinProxyRewardRate, enum.MaxProxyRewardRate)
	}

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		_, err := new(db.Proxy).UpdateRewardRate(ctx, tx, proxyId, storeId, rewardRate)
		if err != nil {
			return err
		}

		return err
	})

	return xerror.Wrap(err)
}

func (t *ProxyService) Recover(ctx context.Context, keeperId, storeId, proxyId string) error {

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		_, err := new(db.Proxy).UpdateStatus(ctx, tx, proxyId, storeId, enum.ProxyStatus_Confirmed.Value)
		if err != nil {
			return err
		}

		return err
	})

	return xerror.Wrap(err)
}

func (t *ProxyService) Delete(ctx context.Context, keeperId, storeId, proxyId string) error {

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		_, err := new(db.Proxy).UpdateStatus(ctx, tx, proxyId, storeId, enum.ProxyStatus_Deleted.Value)
		if err != nil {
			return err
		}

		return err
	})

	return xerror.Wrap(err)
}

func (t *ProxyService) ListProxies(ctx context.Context, keeperId, storeId string, page, size int64) (types.Proxies, int64, error) {

	dbItems, total, err := new(db.Proxy).List(ctx, db.ListProxyParams{
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

func (t *ProxyService) AddProxy(ctx context.Context, keeperId, storeId, userId string, rewardRate float64) error {

	if rewardRate < enum.MinProxyRewardRate || rewardRate > enum.MaxProxyRewardRate {
		return errorx.ParamErrorf("rewardRate is require between %f to %f", enum.MinProxyRewardRate, enum.MaxProxyRewardRate)
	}

	dbProxy := db.Proxy{
		Id:         shortuuid.New()[:8],
		StoreId:    storeId,
		UserId:     userId,
		RewardRate: rewardRate,
		Status:     enum.ProxyStatus_Confirmed.Value,
	}

	_, err := dbProxy.CreateNX(ctx)
	if err != nil {
		return xerror.Wrap(err)
	}
	return nil
}
