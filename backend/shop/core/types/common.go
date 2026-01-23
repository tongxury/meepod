package types

type Extra struct {
	Type  string `json:"type"` // countdown text
	Value any    `json:"value"`
}

type OptionItem struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type OptionItems []OptionItem
