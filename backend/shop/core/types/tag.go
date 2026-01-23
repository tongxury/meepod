package types

type Tag struct {
	Title string `json:"title"`
	Color string `json:"color"`
}

type Tags []*Tag

var SelfTag = &Tag{
	Title: "本人",
	Color: "orange",
}

var NewTag = &Tag{
	Title: "新用户",
	Color: "#1890ff",
}

var SwitchOutTag = &Tag{
	Title: "转出",
	Color: "#1890ff",
}

var SwitchInTag = &Tag{
	Title: "转入",
	Color: "#1890ff",
}

var FollowTag = &Tag{
	Title: "跟",
	Color: "orange",
}
