package enum

type Enum struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
	Desc  string `json:"desc,omitempty"`
	List  []Enum `json:"list,omitempty"`
}

var unknown = Enum{Name: "未知", Value: "unknown", Color: "gray"}
