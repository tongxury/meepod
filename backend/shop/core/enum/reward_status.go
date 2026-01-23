package enum

var (
	RewardStatus_Confirmed = Enum{Value: "confirmed", Name: "已确认待发放", Color: "orange"}
	RewardStatus_Rewarded  = Enum{Value: "rewarded", Name: "已发放", Color: "rgb(11,178,19)"}
	RewardStatus_Rejected  = Enum{Value: "rejected", Name: "已拒绝", Color: "red"}

	AllRewardStatus = []Enum{RewardStatus_Confirmed, RewardStatus_Rewarded, RewardStatus_Rejected}

	RewardableRewardStatus = []string{RewardStatus_Confirmed.Value}
	RejectableRewardStatus = []string{RewardStatus_Confirmed.Value}

	ToHandleByKeeperRewardStatus = []string{
		RewardStatus_Confirmed.Value,
	}
)

func RewardStatus(value string) Enum {
	for _, x := range AllRewardStatus {
		if x.Value == value {
			return x
		}
	}

	return unknown
}
