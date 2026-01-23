package pl5

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

/*
一等奖，所选号码与中奖号码全部相同且顺序一致。例如：中奖号码为43751，则排列5的中奖结果为：43751。
*/
func (t *Ticket) prize(target *Ticket) []types.PrizeRecord {

	var prizeRecords []types.PrizeRecord

	if helper.Contains(t.Wan, target.Wan[0]) &&
		helper.Contains(t.Qian, target.Qian[0]) &&
		helper.Contains(t.Bai, target.Bai[0]) &&
		helper.Contains(t.Shi, target.Shi[0]) &&
		helper.Contains(t.Gen, target.Gen[0]) {

		prizeRecords = append(prizeRecords, types.PrizeRecord{
			Ticket: t,
			Grade:  0,
			Count:  1,
			Amount: RewardPerTicket,
		})
	}

	return prizeRecords

}

func (t *Prizer) FetchTarget(index string) (*types.PrizeResult, error) {

	result, err := fetcher_adapter.GetCn8200Adapter().FindPl5ResultByIndex(index)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	var target = Ticket{
		Wan:         result.Result[0:1],
		Qian:        result.Result[1:2],
		Bai:         result.Result[2:3],
		Shi:         result.Result[3:4],
		Gen:         result.Result[4:5],
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
