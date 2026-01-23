package enum

var (
	OrderStatus_Submitted = Enum{Value: "submitted", Name: "已提交待支付", Color: "orange"}
	OrderStatus_Payed     = Enum{Value: "payed", Name: "已付款待接单", Color: "orange"}
	OrderStatus_Accepted  = Enum{Value: "accepted", Name: "已接单待出票", Color: "orange"}
	OrderStatus_Ticketed  = Enum{Value: "ticketed", Name: "已出票待开奖", Color: "rgb(11,178,19)"}
	OrderStatus_Prized    = Enum{Value: "prized", Name: "已开奖", Color: "gray"}
	//OrderStatus_Rewarded  = Enum{Value: "rewarded", Name: "已兑奖", Color: "gray"}
	//OrderStatus_Rewarded  = Enum{Value: "rewarded", Name: "已兑奖", Color: "gray"}
	OrderStatus_Timeout  = Enum{Value: "timeout", Name: "已超时", Color: "red"}
	OrderStatus_Canceled = Enum{Value: "canceled", Name: "已撤销", Color: "red"}
	OrderStatus_Rejected = Enum{Value: "rejected", Name: "已拒单", Color: "red"}

	SubmittedOrderStatus = []string{
		OrderStatus_Payed.Value,
		OrderStatus_Accepted.Value,
		OrderStatus_Rejected.Value,
		OrderStatus_Canceled.Value,
		OrderStatus_Ticketed.Value,
		OrderStatus_Prized.Value,
		OrderStatus_Timeout.Value,
		//OrderStatus_Rewarded.Value,
	}
	AllOrderStatus = []Enum{
		OrderStatus_Submitted,
		OrderStatus_Payed,
		OrderStatus_Accepted,
		OrderStatus_Ticketed,
		OrderStatus_Prized,
		OrderStatus_Timeout,
		OrderStatus_Canceled,
		OrderStatus_Rejected,
		//OrderStatus_Rewarded,
	}
	FollowableStatus = []string{OrderStatus_Ticketed.Value, OrderStatus_Prized.Value}
	PayableStatus    = []string{OrderStatus_Submitted.Value}
	AcceptableStatus = []string{OrderStatus_Payed.Value}
	CancelableStatus = []string{OrderStatus_Submitted.Value, OrderStatus_Payed.Value}
	RejectableStatus = []string{OrderStatus_Submitted.Value, OrderStatus_Payed.Value}
	TicketableStatus = []string{OrderStatus_Accepted.Value}
	TicketedStatus   = []string{OrderStatus_Ticketed.Value, OrderStatus_Prized.Value}
	PrizableStatus   = []string{OrderStatus_Ticketed.Value}
	// 需要给推广员方法奖励的状态
	ProxyRewardableStatus = []string{OrderStatus_Ticketed.Value, OrderStatus_Prized.Value}
	// 需要扣减在合作店铺的余额的状态
	CoStorePayableStatus = []string{OrderStatus_Ticketed.Value, OrderStatus_Prized.Value}

	ToHandleByKeeperOrderStatus = []string{
		OrderStatus_Submitted.Value,
		OrderStatus_Payed.Value,
		OrderStatus_Accepted.Value,
	}
)

func OrderStatus(value string) Enum {
	for _, x := range AllOrderStatus {
		if x.Value == value {
			return x
		}
	}

	return unknown
}
