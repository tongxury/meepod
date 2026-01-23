package types

import (
	"gitee.com/meepo/backend/shop/core/enum"
)

type PaymentStore struct {
	//AppId     string    `json:"app_id"`
	Store     *Store    `json:"store"`
	CreatedAt string    `json:"created_at"`
	Status    enum.Enum `json:"status"`
	Xinsh     any       `json:"xinsh"`
	Aliyun    any       `json:"aliyun"`
}

type PaymentStores []*PaymentStore
