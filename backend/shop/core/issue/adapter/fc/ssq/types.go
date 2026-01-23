package ssq

import (
	"gitee.com/meepo/backend/kit/components/sdk/helper/mathd"
	"gitee.com/meepo/backend/shop/core/issue/types"
)

// SingleTicket 单式
type SingleTicket Ticket
type SingleTickets []*SingleTicket

// Ticket 复式
type Ticket struct {
	Red  []string `json:"red"`
	Blue []string `json:"blue"`
	DRed []string `json:"redD,omitempty"`
	// 只有开奖结果有这3个字段
	Sales       string            `json:"sales,omitempty"`
	PoolMoney   string            `json:"pool_money,omitempty"`
	PrizeGrades types.PrizeGrades `json:"prize_grades,omitempty"`
}

type SsqTickets []*Ticket

func (t *Ticket) Split() SsqTickets {

	_, amount := t.Amount()

	// 红球最大个数为16个 金额为 16016, 不超过20000。 一旦超过必是蓝复式或全复式
	if amount > MaxAmountPerTicket {
		var rsp SsqTickets

		// 将蓝复式拆成多个蓝单式
		for _, x := range t.Blue {
			rsp = append(rsp, &Ticket{
				Red:  t.Red,
				Blue: []string{x},
				DRed: t.DRed,
			})
		}

		return rsp
	}

	return SsqTickets{t}
}

func (t *Ticket) Amount() (int, float64) {
	rl := len(t.Red)
	bl := len(t.Blue)
	drl := len(t.DRed)

	if bl < 1 {
		return 0, 0
	}

	if rl+drl < 6 {
		return 0, 0
	}

	optionCount := (mathd.Factorial(rl) / (mathd.Factorial(6-drl) * mathd.Factorial(rl-(6-drl)))) * bl

	return optionCount, float64(optionCount * 2)
}

//func (t *Ticket) prize(target *SingleTicket) []types.PrizeRecord {
//
//	singleTickets := t.SingleTickets()
//
//	var rsp []types.PrizeRecord
//	for _, x := range singleTickets {
//		if grade := x.prize(target); grade > 0 {
//
//			y := types.PrizeRecord{
//				Ticket: x,
//				Grade:  grade,
//				Count:  1,
//			}
//
//			for _, prizeGrade := range target.PrizeGrades {
//				if prizeGrade.Grade == grade {
//					y.Amount = conv.Float64(prizeGrade.Amount)
//				}
//			}
//
//			rsp = append(rsp, y)
//		}
//	}
//
//	return rsp
//}
