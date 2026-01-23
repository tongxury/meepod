package enum

var (
	CoStorePaymentStatus_Confirmed = Enum{Value: "confirmed", Name: "已确认", Color: "gray"}

	AllCoStorePaymentStatus = []Enum{CoStorePaymentStatus_Confirmed}
)

func CoStorePaymentStatus(value string) Enum {
	for _, x := range AllCoStorePaymentStatus {
		if x.Value == value {
			return x
		}
	}

	return unknown
}
