package types

import (
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"time"
)

type Proxy struct {
	Id           string    `json:"id"`
	User         *User     `json:"user"`
	RewardRate   float64   `json:"reward_rate"`
	RewardAmount float64   `json:"reward_amount"`
	UserCount    int64     `json:"user_count"`
	OrderCount   int64     `json:"order_count"`
	OrderAmount  float64   `json:"order_amount"`
	CreatedAt    string    `json:"created_at"`
	CreatedAtTs  int64     `json:"created_at_ts"`
	Status       enum.Enum `json:"status"`
	Tags         Tags      `json:"tags"`
	Deletable    bool      `json:"deletable"`
	Recoverable  bool      `json:"recoverable"`
	Addable      bool      `json:"addable"`
	Updatable    bool      `json:"updatable"`
}

type Proxies []*Proxy

func FromDbProxy(x *db.Proxy, user *db.User) *Proxy {

	if x == nil {
		return nil
	}

	rsp := Proxy{
		Id:          x.Id,
		User:        FromDbUser(user),
		CreatedAt:   x.CreatedAt.Format(time.DateOnly),
		CreatedAtTs: x.CreatedAt.Unix(),
		RewardRate:  x.RewardRate,
		Status:      enum.ProxyStatus(x.Status),
		Tags:        nil,
		Deletable:   helper.InSlice(x.Status, enum.DeletableProxyStatus),
		Recoverable: helper.InSlice(x.Status, []string{enum.ProxyStatus_Deleted.Value}),
		Addable:     helper.InSlice(x.Status, enum.AddableProxyStatus),
		Updatable:   helper.InSlice(x.Status, enum.UpdatableProxyStatus),
	}
	return &rsp
}
