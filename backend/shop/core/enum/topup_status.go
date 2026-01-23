package enum

var (
	TopupStatus_Submitted = Enum{Value: "submitted", Name: "已提交未支付", Color: "orange"}
	TopupStatus_Payed     = Enum{Value: "payed", Name: "已支付", Color: "green"}
	TopupStatus_Canceled  = Enum{Value: "canceled", Name: "已取消", Color: "red"}
	TopupStatus_Timeout   = Enum{Value: "timeout", Name: "已超时", Color: "red"}

	AllTopupStatus = []Enum{
		TopupStatus_Submitted,
		TopupStatus_Payed,
		TopupStatus_Canceled,
		TopupStatus_Timeout,
	}

	PayableTopupStatus    = []string{TopupStatus_Submitted.Value}
	CancelableTopupStatus = []string{TopupStatus_Submitted.Value}
)

func TopupStatus(value string) Enum {
	for _, x := range AllTopupStatus {
		if x.Value == value {
			return x
		}
	}

	return unknown
}
