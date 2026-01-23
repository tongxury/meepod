package enum

var (
	MemberLevel_Normal   = Enum{Value: "normal", Name: "普通会员", Color: "orange"}
	MemberLevel_Advanced = Enum{Value: "advanced", Name: "高级会员", Color: "orange"}

	AllMemberLevels = []Enum{MemberLevel_Normal, MemberLevel_Advanced}
)

func MemberLevel(value string) Enum {
	for _, x := range AllMemberLevels {
		if x.Value == value {
			return x
		}
	}

	return unknown
}
