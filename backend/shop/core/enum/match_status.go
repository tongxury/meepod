package enum

var (
	MatchStatus_Pending = Enum{Value: "pending", Name: "gray"} // 爬取到了 但未生成对应的issue
	MatchStatus_UnStart = Enum{Value: "unStart", Name: "未开始", Color: "orange"}
	MatchStatus_End     = Enum{Value: "end", Name: "已结束", Color: "rgb(186, 26, 26)"}

	AllMatchStatus = []Enum{MatchStatus_Pending, MatchStatus_UnStart, MatchStatus_End}
)

func MatchStatus(value string) Enum {
	for _, x := range AllMatchStatus {
		if x.Value == value {
			return x
		}
	}

	return unknown
}
