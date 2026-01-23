package enum

var (
	TopupCategory_Wallet      = Enum{Value: "wallet", Name: "钱包充值", Color: "orange"}
	TopupCategory_Buying      = Enum{Value: "buying", Name: "彩票购买", Color: "orange"}
	TopupCategory_ProxyReward = Enum{Value: "proxyReward", Name: "推广佣金", Color: "orange"}
	TopupCategory_Reward      = Enum{Value: "reward", Name: "彩票中奖", Color: "rgb(11,178,19)"}

	AllTopupCategories = []Enum{
		TopupCategory_Wallet,
		TopupCategory_Buying,
		TopupCategory_ProxyReward,
		TopupCategory_Reward,
	}
)

func TopupCategory(value string) Enum {
	for _, x := range AllTopupCategories {
		if x.Value == value {
			return x
		}
	}

	return unknown
}
