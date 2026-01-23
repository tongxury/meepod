package enum

var (
	FeedbackStatus_Submitted = Enum{Value: "submitted", Name: "已提交", Color: "orange"}
	FeedbackStatus_Resolved  = Enum{Value: "resolved", Name: "已解决", Color: "gray"}

	ResolvableFeedbackStatus = []string{FeedbackStatus_Submitted.Value}

	AllFeedbackStatus = []Enum{FeedbackStatus_Submitted, FeedbackStatus_Resolved}
)

func FeedbackStatus(value string) Enum {
	for _, x := range AllFeedbackStatus {
		if x.Value == value {
			return x
		}
	}

	return unknown
}
