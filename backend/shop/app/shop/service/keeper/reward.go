package keeperservice

import (
	"context"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/types"
	"github.com/go-pg/pg/v10"
	redisV9 "github.com/redis/go-redis/v9"
)

type RewardService struct {
}

func (t *RewardService) RequireByIdAndStoreId(ctx context.Context, id, storeId string) (*types.Reward, error) {

	reward, err := new(db.Reward).RequireByIdAndStoreId(ctx, id, storeId)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	rewards, err := t.Assemble(ctx, db.Rewards{reward})
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return rewards[0], nil
}

func (t *RewardService) Assemble(ctx context.Context, rewards db.Rewards) (types.Rewards, error) {

	var userIds []string
	_, accountUserIds, storeIds := rewards.Ids()
	userIds = append(userIds, accountUserIds...)

	dbStores, err := new(db.Store).ListByIds(ctx, storeIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	dbStoresTap := dbStores.AsMap()
	_, storeOwnerIds := dbStores.Ids()
	userIds = append(userIds, storeOwnerIds...)

	dbUsers, err := new(db.User).ListByIds(ctx, userIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	dbUsersTap := dbUsers.AsMap()

	var rsp types.Rewards
	for _, x := range rewards {

		store := dbStoresTap[x.StoreId]
		accountUser := dbUsersTap[x.UserId]
		storeOwner := dbUsersTap[store.OwnerId]

		y := types.FromDbReward(x, accountUser, store, storeOwner)

		rsp = append(rsp, y)
	}

	return rsp, nil
}

func (t *RewardService) Reject(ctx context.Context, rewardId, storeId, reason string) error {
	_, err := new(db.Reward).RequireByIdAndStoreId(ctx, rewardId, storeId)
	if err != nil {
		return xerror.Wrap(err)
	}

	err = comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		rewarded, err := new(db.Reward).UpdateToRejected(ctx, tx, rewardId, reason)
		if err != nil {
			return err
		}

		if !rewarded {
			return xerror.Wrapf("cannot be payed: %s", rewardId)
		}

		return nil
	})

	return xerror.Wrap(err)
}

func (t *RewardService) Reward(ctx context.Context, rewardId, storeId string) error {

	reward, err := new(db.Reward).RequireByIdAndStoreId(ctx, rewardId, storeId)
	if err != nil {
		return xerror.Wrap(err)
	}

	err = comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		rewarded, err := new(db.Reward).UpdateToRewarded(ctx, tx, rewardId)
		if err != nil {
			return err
		}

		if !rewarded {
			return xerror.Wrapf("cannot be payed: %s", rewardId)
		}

		return nil
	})

	err = comp.SDK().Redis().XAdd(ctx, &redisV9.XAddArgs{
		Stream: "reward.pay.event",
		Values: map[string]interface{}{
			"Id":          reward.Id,
			"UserId":      reward.UserId,
			"StoreId":     reward.StoreId,
			"Amount":      reward.Amount,
			"BizId":       reward.BizId,
			"BizCategory": reward.BizCategory,
		},
	}).Err()
	if err != nil {
		slf.WithError(err).Errorw("XAdd err")
	}

	return xerror.Wrap(err)
}

func (t *RewardService) ListRewards(ctx context.Context, storeId string, page, size int64) (types.Rewards, int64, error) {

	dbRewards, total, err := new(db.Reward).List(ctx, db.ListRewardsParams{
		StoreId: storeId, Page: page, Size: size,
	})
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	rsp, err := t.Assemble(ctx, dbRewards)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	return rsp, total, nil
}
