package enum

var (
	// 理论上这里每一项都会对应一张表, 和biz_id联合使用 用户定位哪张表
	BizCategory_Order       = Enum{Value: "order", Name: "彩票订单", Color: "orange"}
	BizCategory_OrderGroup  = Enum{Value: "orderGroup", Name: "合买订单", Color: "orange"}
	BizCategory_GroupShare  = Enum{Value: "groupShare", Name: "合买份额", Color: "green"}
	BizCategory_ProxyReward = Enum{Value: "proxyReward", Name: "推广佣金", Color: "green"}
	//BizCategory_DecrByKeeper = Enum{Value: "decrByKeeper", Name: "店主扣减", Color: "red"}

	AllBizCategories = []Enum{
		BizCategory_Order,
		BizCategory_OrderGroup,
		BizCategory_GroupShare,
		BizCategory_ProxyReward,
	}
)

func BizCategory(value string) *Enum {
	for _, x := range AllBizCategories {
		if x.Value == value {
			return &x
		}
	}

	return nil
}
