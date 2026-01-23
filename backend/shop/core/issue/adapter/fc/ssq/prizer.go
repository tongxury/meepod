package ssq

import (
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/shop/core/issue/fetcher_adapter"
	"gitee.com/meepo/backend/shop/core/issue/types"
)

type Prizer struct {
}

func (t *Prizer) Prize(src, target string) (types.PrizeRecords, error) {
	tickets, _, _, err := new(Parser).Parse(src)
	if err != nil {
		return nil, err
	}

	targets, _, _, err := new(Parser).Parse(target)
	if err != nil {
		return nil, err
	}

	if len(targets) == 0 {
		return nil, fmt.Errorf("invalid format: %s", target)
	}

	targ := targets[0].(*Ticket)

	var rsp []types.PrizeRecord
	for _, x := range tickets {
		rsp = append(rsp, x.(*Ticket).prize(targ)...)
	}
	return rsp, nil
}

// Prize
// 一等奖：7个号码相符(6个红色球号码和1个蓝色球号码)(红色球号码顺序不限，下同)
// 二等奖：6个红色球号码相符；
// 三等奖：5个红色球号码和1个蓝色球号码相符；
// 四等奖：5个红色球号码或4个红色球号码和1个蓝色球号码相符；
// 五等奖：4个红色球号码或3个红色球号码和1个蓝色球号码相符；
// 六等奖：1个蓝色球号码相符(有无红色球号码相符均可)。
func (t *SingleTicket) prize(target *Ticket) int {

	redHits, _ := helper.Intersect(t.Red, target.Red)
	blueHits, _ := helper.Intersect(t.Blue, target.Blue)

	redHitsCount := len(redHits)
	blueHitsCount := len(blueHits)

	if redHitsCount == 6 && blueHitsCount == 1 {
		return 1
	} else if redHitsCount == 6 {
		return 2
	} else if redHitsCount == 5 && blueHitsCount == 1 {
		return 3
	} else if redHitsCount == 5 || (redHitsCount == 4 && blueHitsCount == 1) {
		return 4
	} else if redHitsCount == 4 || (redHitsCount == 3 && blueHitsCount == 1) {
		return 5
	} else if blueHitsCount == 1 {
		return 6
	}

	return 0
}

func (t *Ticket) prize(target *Ticket) []types.PrizeRecord {

	singleTickets := t.SingleTickets()

	var rsp []types.PrizeRecord
	for _, x := range singleTickets {
		if grade := x.prize(target); grade > 0 {

			y := types.PrizeRecord{
				Ticket: x,
				Grade:  grade,
				Count:  1,
			}

			for _, prizeGrade := range target.PrizeGrades {
				if prizeGrade.Grade == grade {
					y.Amount = conv.Float64(prizeGrade.Amount)
				}
			}

			rsp = append(rsp, y)
		}
	}

	return rsp
}

// SingleTickets 复式拆成多个单式
func (t *Ticket) SingleTickets() SingleTickets {

	redOptions := helper.Combine(t.Red, 6-len(t.DRed))

	var rsp SingleTickets
	for _, redOption := range redOptions {
		for _, blueOption := range t.Blue {
			rsp = append(rsp, &SingleTicket{
				Red:  append(t.DRed, redOption...),
				Blue: []string{blueOption},
			})
		}
	}

	return rsp
}

func (t *Prizer) FetchTarget(index string) (*types.PrizeResult, error) {
	result, err := fetcher_adapter.GetCn8200Adapter().FindSsqResultByIndex(index)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	var ticket = SingleTicket{
		Red:         result.Red,
		Blue:        []string{result.Blue},
		PrizeGrades: result.PrizeGrades,
	}

	return &types.PrizeResult{
		Issue:   index,
		Result:  conv.S2J(SingleTickets{&ticket}),
		PrizeAt: result.PrizeAt(PrizeHour, PrizeMinute),
	}, nil

}
