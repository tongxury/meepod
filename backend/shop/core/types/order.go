package types

import (
	"gitee.com/meepo/backend/shop/core/enum"
)

type OrderForm struct {
	PlanId     string `binding:"required"`
	NeedUpload bool
}

type CreateGroupOrderForm struct {
	PlanId      string `binding:"required"`
	TotalVolume int64  `binding:"required"`
	Volume      int64
	RewardRate  float64
	Remark      string
	Floor       int64
	needUpload  bool
}

type FollowOrderForm struct {
	FollowOrderId string `binding:"required"`
}

type Order struct {
	Id            string      `json:"id"`
	Plan          *Plan       `json:"plan"`
	Store         *Store      `json:"store"`
	ToStore       *Store      `json:"to_store"`
	User          *User       `json:"user"`
	Volume        int64       `json:"volume"`
	Amount        float64     `json:"amount"`
	Group         *OrderGroup `json:"group"`
	FollowOrderId string      `json:"follow_order_id"`
	CreatedAt     string      `json:"created_at"`
	CreatedAtTs   int64       `json:"created_at_ts"`
	Status        enum.Enum   `json:"status"`
	Cancelable    bool        `json:"cancelable"`
	Rejectable    bool        `json:"rejectable"`
	Acceptable    bool        `json:"acceptable"`
	Ticketable    bool        `json:"ticketable"`
	Switchable    bool        `json:"switchable"`
	Followable    bool        `json:"followable"`
	Payable       bool        `json:"payable"`
	KeeperPayable bool        `json:"keeper_payable"`
	TicketImages  []string    `json:"ticket_images,omitempty"`
	Prized        bool        `json:"prized,omitempty"`
	NeedUpload    bool        `json:"need_upload,omitempty"`
	Tags          Tags        `json:"tags,omitempty"`
}

type Orders []*Order
