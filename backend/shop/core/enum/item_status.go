package enum

var (
	ItemStatus_Unstart = Enum{Value: "unstart", Name: "下一期未开始", Color: "orange", Desc: "下一期未开始"}
	ItemStatus_Ongoing = Enum{Value: "ongoing", Name: "进行中", Color: "rgb(11,178,19)"}                 // db 用到
	ItemStatus_NoData  = Enum{Value: "ondata", Name: "暂无场次", Color: "rgb(186, 26, 26)", Desc: "暂无场次"} // db 用到
	ItemStatus_Closed  = Enum{Value: "closed", Name: "本期已截止", Color: "rgb(186, 26, 26)", Desc: "本期已截止"}
	ItemStatus_Unable  = Enum{Value: "unable", Name: "不可用", Color: "orange", Desc: "请先选择商户"}
	ItemStatus_Invalid = Enum{Value: "invalid", Name: "店铺状态异常", Color: "orange", Desc: "请确认商户状态"}

	AllItemStatus = []Enum{ItemStatus_Unstart, ItemStatus_Ongoing, ItemStatus_Closed, ItemStatus_Unable, ItemStatus_Invalid}
)

func ItemStatus(value string) Enum {
	for _, x := range AllItemStatus {
		if x.Value == value {
			return x
		}
	}

	return unknown
}
