package types

import (
	"gitee.com/meepo/backend/shop/core/enum"
)

type CoStorePayment struct {
	Id          string     `json:"id"`
	Store       *Store     `json:"store"`
	CoStore     *Store     `json:"co_store"`
	Amount      float64    `json:"amount"`
	CreatedAt   string     `json:"created_at"`
	CreatedAtTs int64      `json:"created_at_ts"`
	Status      enum.Enum  `json:"status"`
	Category    enum.Enum  `json:"category"`
	BizCategory *enum.Enum `json:"biz_category"`
	BizId       string     `json:"biz_id"`
	Payable     bool       `json:"payable"`
	Cancelable  bool       `json:"cancelable"`
	Payed       bool       `json:"payed"`
}

type CoStorePayments []*CoStorePayment
