package enum

var (
	WithdrawStatus_Submitted = Enum{Value: "submitted", Name: "已提交", Color: "orange"}
	WithdrawStatus_Accepted  = Enum{Value: "accepted", Name: "已受理", Color: "green"}
	WithdrawStatus_Canceled  = Enum{Value: "canceled", Name: "已取消", Color: "red"}
	WithdrawStatus_Rejected  = Enum{Value: "rejected", Name: "已拒绝", Color: "red"}

	AllWithdrawStatus = []Enum{
		WithdrawStatus_Submitted,
		WithdrawStatus_Accepted,
		WithdrawStatus_Canceled,
		WithdrawStatus_Rejected,
	}

	ToHandleByKeeperWithdrawStatus = []string{
		WithdrawStatus_Submitted.Value,
	}

	AcceptableWithdrawStatus = []string{WithdrawStatus_Submitted.Value}
	RejectableWithdrawStatus = []string{WithdrawStatus_Submitted.Value}
	CancelableWithdrawStatus = []string{WithdrawStatus_Submitted.Value}
)

func WithdrawStatus(value string) Enum {
	for _, x := range AllWithdrawStatus {
		if x.Value == value {
			return x
		}
	}

	return unknown
}
