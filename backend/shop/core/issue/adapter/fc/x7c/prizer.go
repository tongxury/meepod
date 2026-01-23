package x7c

import (
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
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

/*
一等奖：投注号码的全部数字与开奖号码对应位置数字均相同，即中奖；
二等奖：投注号码的前6位数字与开奖号码对应位置数字相同，即中奖；
三等奖：投注号码前6位中的任意5个数字与开奖号码对应位置数字相同且最后一个数字与开奖号码对应位置数字相同，即中奖；
四等奖：投注号码中任意5个数字与开奖号码对应位置数字相同，即中奖；
五等奖：投注号码中任意4个数字与开奖号码对应位置数字相同，即中奖；
六等奖：投注号码中任意3个数字与开奖号码对应位置数字相同，或者投注号码前6位中的任意1个数字与开奖号码对应位置数字相同且最后一个数字与开奖号码对应位置数字相同，或者仅最后一个数字与开奖号码对应位置数字相同，即中奖。
*/

func (t *Ticket) prize(target *Ticket) []types.PrizeRecord {

	var prizeRecords []types.PrizeRecord

	// 对应位置的数字个数
	allCountStats := []int{len(t.Swan), len(t.Wan), len(t.Qian), len(t.Bai), len(t.Shi), len(t.Gen), len(t.Last)}
	front6HitStats := []bool{
		helper.InSlice(target.Swan[0], t.Swan),
		helper.InSlice(target.Wan[0], t.Wan),
		helper.InSlice(target.Qian[0], t.Qian),
		helper.InSlice(target.Bai[0], t.Bai),
		helper.InSlice(target.Shi[0], t.Shi),
		helper.InSlice(target.Gen[0], t.Gen),
	}

	allHitStats := append(front6HitStats, helper.InSlice(target.Last[0], t.Last))

	var front6HitsCount int
	for _, x := range front6HitStats {
		if x {
			front6HitsCount += 1
		}
	}

	var lastHitsCount int
	if helper.InSlice(target.Last[0], t.Last) {
		lastHitsCount += 1
	}

	var grade, count int //注数
	if front6HitsCount == 6 && lastHitsCount == 1 {
		grade = 1
		count = 1
	} else if front6HitsCount == 6 {
		grade = 2
		count = len(t.Last)
	} else if front6HitsCount == 5 && lastHitsCount == 1 {
		grade = 3
		count = 1
		for i, x := range front6HitStats {
			if !x {
				count *= allCountStats[i]
			}
		}
	} else if front6HitsCount+lastHitsCount == 5 {
		grade = 4
		count = 1
		for i, x := range allHitStats {
			if !x {
				count *= allCountStats[i]
			}
		}
	} else if front6HitsCount+lastHitsCount == 4 {
		grade = 5
		count = 1
		for i, x := range allHitStats {
			if !x {
				count *= allCountStats[i]
			}
		}
	} else if front6HitsCount+lastHitsCount == 3 {
		grade = 6
		count = 1
		for i, x := range allHitStats {
			if !x {
				count *= allCountStats[i]
			}
		}
	} else if front6HitsCount == 0 && lastHitsCount == 1 {
		grade = 6
		count = 1
		for i, x := range front6HitStats {
			if !x {
				count *= allCountStats[i]
			}
		}
	} else if front6HitsCount == 1 && lastHitsCount == 1 {
		grade = 6
		count = 1

		for i, x := range front6HitStats {
			if !x {
				count *= allCountStats[i]
			}
		}

	}

	if pg, found := target.PrizeGrades.Find(grade); found {
		prizeRecords = append(prizeRecords, types.PrizeRecord{
			Ticket: t,
			Grade:  grade,
			Count:  count,
			Amount: conv.Float64(pg.Amount),
		})

	} else {
		slf.Errorw("[MAIN] no grade found", slf.Int("grade", grade), slf.Reflect("t", t), slf.Reflect("target", target))
	}

	return prizeRecords

}

func (t *Prizer) FetchTarget(index string) (*types.PrizeResult, error) {
	result, err := fetcher_adapter.GetCn8200Adapter().FindX7cResultByIndex(index)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	var target = Ticket{
		Swan:        result.Result[0:1],
		Wan:         result.Result[1:2],
		Qian:        result.Result[2:3],
		Bai:         result.Result[3:4],
		Shi:         result.Result[4:5],
		Gen:         result.Result[5:6],
		Last:        result.Result[6:7],
		Sales:       "",
		PoolMoney:   "",
		PrizeGrades: result.PrizeGrades,
	}

	return &types.PrizeResult{
		Issue:   index,
		Result:  conv.S2J(Tickets{&target}),
		PrizeAt: result.PrizeAt(PrizeHour, PrizeMinute),
	}, nil

}
