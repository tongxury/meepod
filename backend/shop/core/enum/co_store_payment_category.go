package enum

var (
	CoStorePaymentCategory_TopUp      = Enum{Value: "topUp", Name: "充值", Color: "green"}
	CoStorePaymentCategory_SwitchDecr = Enum{Value: "switchDecr", Name: "出票扣减", Color: "green"}
	CoStorePaymentCategory_SwitchFee  = Enum{Value: "switchFee", Name: "转单佣金", Color: "green"}
	CoStorePaymentCategory_Return     = Enum{Value: "return", Name: "退还预充", Color: "green"}

	AllCoStorePaymentCategories = []Enum{
		CoStorePaymentCategory_TopUp,
		CoStorePaymentCategory_SwitchDecr,
		CoStorePaymentCategory_SwitchFee,
		CoStorePaymentCategory_Return,
	}
)

func CoStorePaymentCategory(value string) Enum {
	for _, x := range AllCoStorePaymentCategories {
		if x.Value == value {
			return x
		}
	}

	return unknown
}
