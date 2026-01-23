package types

import (
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/kit/components/sdk/helper/timed"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
)

type Reward struct {
	Id          string     `json:"id"`
	User        *User      `json:"user"`
	BizId       string     `json:"biz_id"`
	BizCategory *enum.Enum `json:"biz_category"`
	Store       *Store     `json:"store"`
	CreatedAt   string     `json:"created_at"`
	CreatedAtTs int64      `json:"created_at_ts"`
	Status      enum.Enum  `json:"status"`
	Amount      float64    `json:"amount"`
	Rewardable  bool       `json:"rewardable"`
	Rejectable  bool       `json:"rejectable"`
}

type Rewards []*Reward

func FromDbReward(reward *db.Reward, accountUser *db.User, store *db.Store, storeOwner *db.User) *Reward {

	rsp := Reward{
		Id:          reward.Id,
		User:        FromDbUser(accountUser),
		BizId:       reward.BizId,
		BizCategory: enum.BizCategory(reward.BizCategory),
		Store:       FromDbStore(store, storeOwner),
		CreatedAtTs: reward.CreatedAt.Unix(),
		CreatedAt:   timed.SmartTime(reward.CreatedAt.Unix()),
		Status:      enum.RewardStatus(reward.Status),
		Amount:      reward.Amount,
		Rewardable:  helper.InSlice(reward.Status, enum.RewardableRewardStatus),
		Rejectable:  helper.InSlice(reward.Status, enum.RejectableRewardStatus),
	}

	return &rsp
}
