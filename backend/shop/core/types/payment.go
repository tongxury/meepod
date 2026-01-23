package types

import (
	"gitee.com/meepo/backend/shop/core/enum"
)

type PaymentForm struct {
	OrderId  string `binding:"required"`
	Category string
}

type Payment struct {
	Id          string     `json:"id"`
	User        *User      `json:"user"`
	Store       *Store     `json:"store"`
	BizId       string     `json:"biz_id"`
	BizCategory *enum.Enum `json:"biz_category"`
	Category    enum.Enum  `json:"category"`
	Amount      float64    `json:"amount"`
	CreatedAtTs int64      `json:"created_at_ts"`
	CreatedAt   string     `json:"created_at"`
	Status      enum.Enum  `json:"status"`
	Remark      string     `json:"remark"`
}

type Payments []*Payment

type PayMethod struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type PayMethods []*PayMethod
