package enum

var (
	AccountStatus_Normal = Enum{Value: "normal", Name: "正常"}

	AllAccountStatus = []Enum{AccountStatus_Normal}
)

func AccountStatus(value string) Enum {
	for _, x := range AllAccountStatus {
		if x.Value == value {
			return x
		}
	}

	return unknown
}
