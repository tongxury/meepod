package fetcher_adapter

import (
	"gitee.com/meepo/backend/shop/core/issue/fetcher_adapter/cn8200"
	"gitee.com/meepo/backend/shop/core/issue/fetcher_adapter/cwlgov"
	"gitee.com/meepo/backend/shop/core/issue/fetcher_adapter/lottery"
)

func GetCn8200Adapter() Adapter {
	return &cn8200.Adapter{}
}

func GetLotteryAdapter() Adapter {
	return &lottery.Adapter{}
}

func GetCwlGovAdapter() Adapter {
	return &cwlgov.Adapter{}
}
