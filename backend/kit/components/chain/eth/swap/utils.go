package swap

import (
	"gitee.com/meepo/backend/kit/components/chain/eth"
)

// 按照token0为买卖币种，token1作为稳定币 方便后续处理
// 稳定系数
func getStability(token string) int {
	for i, x := range eth.STABLE_COINS {
		if x == token {
			return i + 1
		}
	}
	return 0
}
