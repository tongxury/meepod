package keeperservice

import (
	"context"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/types"
	"github.com/go-pg/pg/v10"
	"time"
)

type ProxyRewardService struct {
}

func (t *ProxyRewardService) FindById(ctx context.Context, id string) (*types.ProxyReward, error) {
	return nil, nil
}

func (t *ProxyRewardService) Pay(ctx context.Context, keeperId, storeId, rewardId string) error {

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		_, err := new(db.ProxyReward).UpdatePayed(ctx, tx, storeId, rewardId)
		if err != nil {
			return err
		}

		return nil
	})

	return xerror.Wrap(err)
}

func (t *ProxyRewardService) ListProxyRewards(ctx context.Context, storeId, month, status string, page, size int64) (types.ProxyRewards, int, error) {

	dbItems, total, err := new(db.ProxyReward).List(ctx, db.ListProxyRewardsParams{
		StoreId: storeId, Month: month, Status: status,
		Page: page, Size: size,
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

func (t *ProxyRewardService) Assemble(ctx context.Context, items db.ProxyRewards) (types.ProxyRewards, error) {

	_, proxyUserIds, _ := items.Ids()

	dbUsers, err := new(db.User).ListByIds(ctx, proxyUserIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	dbUsersMap := dbUsers.AsMap()

	var proxyUsers types.ProxyRewards
	for _, x := range items {

		y := types.ProxyReward{
			Id:           x.Id,
			User:         types.FromDbUser(dbUsersMap[x.ProxyUserId]),
			Month:        x.Month,
			UserCount:    x.UserCount,
			OrderCount:   x.OrderCount,
			OrderAmount:  x.OrderAmount,
			RewardRate:   x.RewardRate,
			RewardAmount: x.RewardAmount,
			Status:       enum.ProxyRewardStatus(x.Status),
			PayAt:        x.PayAt.Format(time.DateOnly),
			CreatedAt:    x.CreatedAt.Format(time.DateOnly),
			CreatedAtTs:  x.CreatedAt.Unix(),
			Payable:      helper.InSlice(x.Status, enum.PayableProxyRewardStatus),
			Tags:         nil,
		}

		proxyUsers = append(proxyUsers, &y)
	}

	return proxyUsers, nil
}
