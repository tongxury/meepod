package enum

var (
	StorePaymentStatus_Confirmed = Enum{Value: "confirmed", Name: "已确认", Color: "gray"}

	AllStorePaymentStatus = []Enum{
		StorePaymentStatus_Confirmed,
	}
)

func StorePaymentStatus(value string) Enum {
	for _, x := range AllStorePaymentStatus {
		if x.Value == value {
			return x
		}
	}

	return unknown
}
