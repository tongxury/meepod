package enum

var (
	StoreStatus_Pending   = Enum{Value: "pending", Name: "审核中", Color: "orange"}
	StoreStatus_Confirmed = Enum{Value: "confirmed", Name: "正常", Color: "rgb(11,178,19)"}
	StoreStatus_Deleted   = Enum{Value: "deleted", Name: "已删除", Color: "red"}

	AllStoreStatus = []Enum{StoreStatus_Pending, StoreStatus_Confirmed, StoreStatus_Deleted}
)

func StoreStatus(value string) Enum {
	for _, x := range AllStoreStatus {
		if x.Value == value {
			return x
		}
	}

	return unknown
}
