package enum

var (
	StorePaymentCategory_TopUp    = Enum{Value: "topUp", Name: "平台充值", Color: "green"}
	StorePaymentCategory_SwichFee = Enum{Value: "switchFee", Name: "转单服务费", Color: "red"}

	AllStorePaymentCategories = []Enum{
		StorePaymentCategory_TopUp,
		StorePaymentCategory_SwichFee,
	}
)

func StorePaymentCategory(value string) Enum {
	for _, x := range AllStorePaymentCategories {
		if x.Value == value {
			return x
		}
	}

	return unknown
}
