package enum

var (
	PaymentCategory_DecrByKeeper = Enum{Value: "decrByKeeper", Name: "店主扣减", Color: "red"}
	PaymentCategory_BuyTicket    = Enum{Value: "buyTicket", Name: "彩票购买", Color: "orange"}

	AllPaymentCategories = []Enum{
		PaymentCategory_DecrByKeeper,
		PaymentCategory_BuyTicket,
	}
)

func PaymentCategory(value string) Enum {
	for _, x := range AllPaymentCategories {
		if x.Value == value {
			return x
		}
	}

	return unknown
}
