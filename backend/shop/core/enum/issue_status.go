package enum

var (
	IssueStatus_Ongoing = Enum{Value: "ongoing", Name: "进行中"} //
	IssueStatus_Prized  = Enum{Value: "prized", Name: "已开奖"}

	AllIssueStatus = []Enum{IssueStatus_Ongoing, IssueStatus_Prized}
)

func IssueStatus(value string) Enum {
	for _, x := range AllIssueStatus {
		if x.Value == value {
			return x
		}
	}

	return unknown
}
