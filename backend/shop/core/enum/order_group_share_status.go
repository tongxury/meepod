package enum

var (
	OrderGroupShareStatus_Submitted = Enum{Value: "submitted", Name: "已提交", Color: "green"}
	OrderGroupShareStatus_Payed     = Enum{Value: "payed", Name: "已付款", Color: "orange"}
	//OrderGroupShareStatus_Timeout   = Enum{Value: "timeout", Name: "已过期", Color: "grey"}

	PayableOrderGroupShareStatus = []string{OrderGroupShareStatus_Submitted.Value}

	AllOrderGroupShareStatus = []Enum{
		OrderGroupShareStatus_Submitted,
		OrderGroupShareStatus_Payed,
		//OrderGroupShareStatus_Timeout,
	}
)

func OrderGroupShareStatus(value string) Enum {
	for _, x := range AllOrderGroupShareStatus {
		if x.Value == value {
			return x
		}
	}

	return unknown
}
