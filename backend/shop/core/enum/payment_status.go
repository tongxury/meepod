package enum

var (
	//  只能有这2种状态 payment 表只负责记录 不能有状态机
	PaymentStatus_Payed    = Enum{Value: "payed", Name: "已消费", Color: "gray"}
	PaymentStatus_Reverted = Enum{Value: "reverted", Name: "已撤销", Color: "red"}

	AllPaymentStatus = []Enum{PaymentStatus_Payed, PaymentStatus_Reverted}

	CanRevertPayStatus = []string{PaymentStatus_Payed.Value}
)

func PaymentStatus(value string) Enum {
	for _, x := range AllPaymentStatus {
		if x.Value == value {
			return x
		}
	}

	return unknown
}
