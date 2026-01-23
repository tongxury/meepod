package enum

var (
	ProxyRewardStatus_Pending   = Enum{Value: "pending", Name: "核算中", Color: "gray"}
	ProxyRewardStatus_Confirmed = Enum{Value: "confirmed", Name: "待结算", Color: "orange"}
	ProxyRewardStatus_Payed     = Enum{Value: "payed", Name: "已结算", Color: "green"}

	PayableProxyRewardStatus = []string{ProxyRewardStatus_Confirmed.Value}

	AllProxyRewardStatus = []Enum{ProxyRewardStatus_Pending, ProxyRewardStatus_Confirmed, ProxyRewardStatus_Payed}
)

func ProxyRewardStatus(value string) Enum {
	for _, x := range AllProxyRewardStatus {
		if x.Value == value {
			return x
		}
	}

	return unknown
}
