package enum

var (
	//OrderGroupStatus_Pending   = Enum{Value: "pending", Name: "未开始", Color: "yellow"}
	OrderGroupStatus_Submitted = Enum{Value: "submitted", Name: "合买中", Color: "orange"}
	OrderGroupStatus_Payed     = Enum{Value: "payed", Name: "已付款待接单", Color: "orange"}
	OrderGroupStatus_Accepted  = Enum{Value: "accepted", Name: "已接单待出票", Color: "orange"}
	OrderGroupStatus_Rejected  = Enum{Value: "rejected", Name: "已拒单", Color: "red"}
	OrderGroupStatus_Ticketed  = Enum{Value: "ticketed", Name: "已出票待开奖", Color: "rgb(11,178,19)"}
	OrderGroupStatus_Prized    = Enum{Value: "prized", Name: "已开奖", Color: "gray"}
	OrderGroupStatus_Timeout   = Enum{Value: "timeout", Name: "已过期", Color: "red"}

	JoinableOrderGroupStatus   = []string{OrderGroupStatus_Submitted.Value}
	RejectableOrderGroupStatus = []string{OrderGroupStatus_Payed.Value, OrderGroupShareStatus_Submitted.Value}
	AcceptableOrderGroupStatus = []string{OrderGroupStatus_Payed.Value}
	TicketableOrderGroupStatus = []string{OrderGroupStatus_Accepted.Value}

	ToHandleByKeeperOrderGroupStatus = []string{
		OrderGroupStatus_Payed.Value,
		OrderGroupStatus_Accepted.Value,
	}
	PrizableGroupStatus = []string{OrderGroupStatus_Ticketed.Value}

	AllOrderGroupStatus = []Enum{OrderGroupStatus_Submitted, OrderGroupStatus_Payed,
		OrderGroupStatus_Accepted, OrderGroupStatus_Rejected, OrderGroupStatus_Ticketed,
		OrderGroupStatus_Prized, OrderGroupStatus_Timeout}
)

func OrderGroupStatus(value string) Enum {
	for _, x := range AllOrderGroupStatus {
		if x.Value == value {
			return x
		}
	}

	return unknown
}
