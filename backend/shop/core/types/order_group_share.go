package types

import "gitee.com/meepo/backend/shop/core/enum"

type GroupShare struct {
	User        *User       `json:"user"`
	Volume      int64       `json:"volume"`
	Id          string      `json:"id"`
	Amount      float64     `json:"amount"`
	Group       *OrderGroup `json:"group"`
	CreatedAtTs int64       `json:"created_at_ts"`
	CreatedAt   string      `json:"created_at"`
	Status      enum.Enum   `json:"status"`
	Tags        []Tag       `json:"tags"`
	//RewardSummary string      `json:"reward_summary"`
}

type GroupShares []*GroupShare
