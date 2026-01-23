package types

import (
	"gitee.com/meepo/backend/shop/core/enum"
)

type PaymentTopup struct {
	Id          string    `json:"id"`
	User        *User     `json:"user"`
	Store       *Store    `json:"store"`
	Amount      float64   `json:"amount"`
	CreatedAt   string    `json:"created_at"`
	CreatedAtTs int64     `json:"created_at_ts"`
	Status      enum.Enum `json:"status"`
	Category    enum.Enum `json:"category"`
	Payable     bool      `json:"payable"`
	Cancelable  bool      `json:"cancelable"`
	PayUrl      string    `json:"pay_url"`
	QrCode      string    `json:"qr_code"`
	PayMethod   string    `json:"pay_method"`
	TimeLeft    int64     `json:"time_left"`
	Payed       bool      `json:"payed"`
}

type PaymentTopups []*PaymentTopup
