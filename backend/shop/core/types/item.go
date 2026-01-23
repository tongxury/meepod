package types

import (
	"gitee.com/meepo/backend/kit/services/util/oss"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
)

type Item struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type Items []*Item

type ItemState struct {
	Item
	Status      enum.Enum `json:"status"`
	LatestIssue *Issue    `json:"latest_issue"`
	Disabled    bool      `json:"disabled"`
	Extra       Extra     `json:"extra"`
}

type ItemStates []*ItemState

func FromDbItems(items db.Items) Items {
	var rsp Items
	for _, x := range items {
		rsp = append(rsp, FromDbItem(x))
	}

	return rsp
}

func FromDbItem(item *db.Item) *Item {
	if item == nil {
		return nil
	}
	return &Item{
		Id:   item.Id,
		Name: item.Name,
		Icon: oss.Resource(item.Icon),
	}
}
