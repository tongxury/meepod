package enum

var (
	PaymentStoreStatus_Pending   = Enum{Value: "pending", Name: "未授权", Color: "red"}
	PaymentStoreStatus_Authed    = Enum{Value: "authed", Name: "已授权", Color: "rgb(11,178,19)"}
	PaymentStoreStatus_Confirmed = Enum{Value: "confirmed", Name: "已通过", Color: "rgb(11,178,19)"}
	PaymentStoreStatus_Rejected  = Enum{Value: "rejected", Name: "已驳回", Color: "red"}

	AllPaymentStoreStatus = []Enum{PaymentStoreStatus_Pending, PaymentStoreStatus_Authed, PaymentStoreStatus_Confirmed, PaymentStoreStatus_Rejected}
)

func PaymentStoreStatus(value string) Enum {
	for _, x := range AllPaymentStoreStatus {
		if x.Value == value {

			x.List = AllPaymentStoreStatus

			return x
		}
	}

	return unknown
}
