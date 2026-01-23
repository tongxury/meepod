package enum

var (
	PlanStatus_Saved   = Enum{Value: "saved", Name: "已保存"}
	PlanStatus_Binding = Enum{Value: "binding", Name: "已绑定"}

	AllPlanStatus = []Enum{PlanStatus_Saved, PlanStatus_Binding}
)

func PlanStatus(value string) Enum {
	for _, x := range AllPlanStatus {
		if x.Value == value {
			return x
		}
	}

	return unknown
}
