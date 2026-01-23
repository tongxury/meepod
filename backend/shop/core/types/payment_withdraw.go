package types

import (
	"gitee.com/meepo/backend/shop/core/enum"
)

type PaymentWithdraw struct {
	Id          string    `json:"id"`
	User        *User     `json:"user"`
	Store       *Store    `json:"store"`
	Amount      float64   `json:"amount"`
	CreatedAt   string    `json:"created_at"`
	CreatedAtTs int64     `json:"created_at_ts"`
	Status      enum.Enum `json:"status"`
	Acceptable  bool      `json:"acceptable"`
	Rejectable  bool      `json:"rejectable"`
	Cancelable  bool      `json:"cancelable"`
	Remark      string    `json:"remark"`
	Image       string    `json:"image,omitempty"`
}

type PaymentWithdraws []*PaymentWithdraw
