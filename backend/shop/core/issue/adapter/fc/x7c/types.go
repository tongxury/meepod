package x7c

import (
	"gitee.com/meepo/backend/shop/core/issue/types"
)

type Tickets []*Ticket

type Ticket struct {
	Swan []string `json:"swan"`
	Wan  []string `json:"wan"`
	Qian []string `json:"qian"`
	Bai  []string `json:"bai"`
	Shi  []string `json:"shi"`
	Gen  []string `json:"gen"`
	Last []string `json:"last"`
	// 只有开奖结果有这3个字段
	Sales       string            `json:"sales,omitempty"`
	PoolMoney   string            `json:"pool_money,omitempty"`
	PrizeGrades types.PrizeGrades `json:"prize_grades,omitempty"`
}

func (t *Ticket) Amount() (int, float64) {

	count := len(t.Gen) * len(t.Shi) * len(t.Bai) * len(t.Qian) * len(t.Wan) * len(t.Wan) * len(t.Last)

	return count, float64(count) * 2
}

func (t *Ticket) Split() Tickets {
	//_, amount := t.Amount()
	//
	//// 红球最大个数为16个 金额为 16016, 不超过20000。 一旦超过必是蓝复式或全复式
	//if amount > MaxAmountPerTicket {
	//	var rsp Tickets
	//
	//	// 将蓝复式拆成多个蓝单式
	//	for _, x := range t.Last {
	//		rsp = append(rsp, &Ticket{
	//			Red:  t.Red,
	//			Blue: []string{x},
	//			DRed: t.DRed,
	//		})
	//	}
	//
	//	return rsp
	//}

	return Tickets{t}
}
