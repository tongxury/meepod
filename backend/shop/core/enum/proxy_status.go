package enum

var (
	ProxyStatus_Pending   = Enum{Value: "pending", Name: "审核中", Color: "orange"}
	ProxyStatus_Confirmed = Enum{Value: "confirmed", Name: "正常", Color: "green"}
	ProxyStatus_Deleted   = Enum{Value: "deleted", Name: "已删除", Color: "red"}

	DeletableProxyStatus = []string{ProxyStatus_Pending.Value, ProxyStatus_Confirmed.Value}
	UpdatableProxyStatus = []string{ProxyStatus_Confirmed.Value}
	AddableProxyStatus   = []string{ProxyStatus_Confirmed.Value}

	AllProxyStatus = []Enum{ProxyStatus_Pending, ProxyStatus_Confirmed, ProxyStatus_Deleted}
)

func ProxyStatus(value string) Enum {
	for _, x := range AllProxyStatus {
		if x.Value == value {
			return x
		}
	}

	return unknown
}
