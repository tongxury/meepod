package types

type ProxyUser struct {
	User        *User   `json:"user"`
	CreatedAt   string  `json:"created_at"`
	CreatedAtTs int64   `json:"created_at_ts"`
	Tags        Tags    `json:"tags"`
	OrderCount  int64   `json:"order_count"`
	OrderAmount float64 `json:"order_amount"`
}

type ProxyUsers []*ProxyUser

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
