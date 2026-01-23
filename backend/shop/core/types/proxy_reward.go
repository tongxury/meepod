package types

import (
	"gitee.com/meepo/backend/shop/core/enum"
)

type ProxyReward struct {
	Id           string    `json:"id"`
	User         *User     `json:"user"`
	Month        string    `json:"month"`
	UserCount    int64     `json:"user_count"`
	OrderCount   int64     `json:"order_count"`
	OrderAmount  float64   `json:"order_amount"`
	RewardRate   float64   `json:"reward_rate"`
	RewardAmount float64   `json:"reward_amount"`
	Status       enum.Enum `json:"status"`
	PayAt        string    `json:"pay_at"`
	CreatedAt    string    `json:"created_at"`
	CreatedAtTs  int64     `json:"created_at_ts"`
	Tags         Tags      `json:"tags"`
	Payable      bool      `json:"payable"`
}

type ProxyRewards []*ProxyReward

//func FromDbProxy(x *db.Proxy, user *db.User) *Proxy {
//
//	if x == nil {
//		return nil
//	}
//
//	rsp := Proxy{
//		Id:          x.Id,
//		User:        FromDbUser(user),
//		CreatedAt:   timed.SmartTime(user.CreatedAt.Unix()),
//		CreatedAtTs: user.CreatedAt.Unix(),
//		RewardRate:  x.RewardRate,
//		MStatus:      enum.ProxyStatus(x.MStatus),
//		Tags:        nil,
//		Deletable:   helper.InSlice(x.MStatus, enum.DeletableProxyStatus),
//		Recoverable: helper.InSlice(x.MStatus, []string{enum.ProxyStatus_Deleted.Value}),
//		Addable:     helper.InSlice(x.MStatus, enum.AddableProxyStatus),
//		Updatable:   helper.InSlice(x.MStatus, enum.UpdatableProxyStatus),
//	}
//	return &rsp
//}
