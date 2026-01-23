package match_adapter

import (
	"gitee.com/meepo/backend/shop/core/issue/match_adapter/lottery"
	"gitee.com/meepo/backend/shop/core/issue/match_adapter/okooo"
)

func GetLotteryAdapter() Adapter {
	return &lottery.Adapter{}
}

func GetOkoooAdapter() Adapter {
	return &okooo.Adapter{}
}
