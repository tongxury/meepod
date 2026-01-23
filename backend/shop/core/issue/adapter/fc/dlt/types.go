package dlt

import (
	"gitee.com/meepo/backend/kit/components/sdk/helper/mathd"
	"gitee.com/meepo/backend/shop/core/issue/types"
)

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

type Tickets []*Ticket

// todo
func (t *Ticket) Split() Tickets {

	//_, amount := t.Amount()
	//
	//// 红球最大个数为16个 金额为 8736, 不超过20000。 一旦超过必是蓝复式
	//if amount > MaxAmountPerTicket {
	//	var rsp Tickets
	//
	//	// 将蓝复式拆成多个蓝单式
	//	for _, x := range t.Blue {
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

func (t *Ticket) Amount() (int, float64) {
	rl := len(t.Red)
	bl := len(t.Blue)
	drl := len(t.DRed)

	if bl < 2 {
		return 0, 0
	}

	if rl+drl < 5 {
		return 0, 0
	}

	optionCount := mathd.Cmn(rl, 5-drl) * mathd.Cmn(bl, 2)

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
