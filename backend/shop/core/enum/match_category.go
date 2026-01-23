package enum

var (
	MatchCategory_Z14 = Enum{Value: "z14", Name: "足彩14场"}
	MatchCategory_Zjc = Enum{Value: "zjc", Name: "竞彩足球"}

	AllMatchCategories = []Enum{MatchCategory_Z14, MatchCategory_Zjc}

	ZcCategories = []string{MatchCategory_Z14.Value, MatchCategory_Zjc.Value}
)

func MatchCategory(value string) Enum {
	for _, x := range AllMatchCategories {
		if x.Value == value {
			return x
		}
	}

	return unknown
}
