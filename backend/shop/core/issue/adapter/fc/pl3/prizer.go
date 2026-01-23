package pl3

import (
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/shop/core/issue/fetcher_adapter"
	"gitee.com/meepo/backend/shop/core/issue/types"
	mapset "github.com/deckarep/golang-set/v2"
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
3D的设奖奖金占彩票销售总额的53%。其中：当期奖金为销售总额的52%；调节基金为销售总额的1%。
3D设置奖池，奖池由设奖奖金与实际中出奖金的差额组成。当期中出奖金未达设奖金额时，余额进入奖池；当期中出奖金超出当期奖金时，差额由奖池补充。当奖池总额不足时，由调节基金补足，调节基金不足时，用发行经费垫支。3D在各省(区、市)保留各自奖池。
3D采用固定设奖。当期奖金设"单选"、"组选3"、"组选6"三个奖等，各奖等中奖奖金额按固定奖金结构设置，规定如下：
"单选"投注：中奖金额每注1040元；
"组选3"投注：中奖号码中任意两位数字相同(开奖号码中的3个数字有两个相同，即为组三)，所选号码与中奖号码相同且顺序不限，则该注彩票中奖。例如，中 奖号码为 544 ，则中奖结果为： 544 、 454 、 445 之一均可。中奖金额为每注346元。
"组选6"投注：中奖金额为每注173元。

1.单选（又叫直选）： 单选投注的号码与当期公布的中奖号码的3位数字及排列顺序都相同，即为中单选奖。如投注单选号码“123”，开奖号码为“123”即为中奖。
2.组选三： 当期开出的中奖号码3位数中有任意两位数字相同，且投注号码的三位数字与中奖号码相同，顺序不限，即中得“组选3”奖。如投注组选3号码“112”，开奖号码为“112”、“121”和“211”中的任何一个都为中奖。
3.组选6： 当期开出的中奖号码3位数各不相同，且投注号码的三位数与中奖号码相同，顺序不限，即中得“组选6”奖。如投注组选6号码“123”，开奖号码为“123”、“132”、“312”、“321”、“213”、“231”中的任何一个都为中奖。
*/

func (t *Ticket) prize(target *Ticket) []types.PrizeRecord {

	var prizeRecords []types.PrizeRecord

	switch t.Cat {
	case "z1":
		if helper.Contains(t.Bai, target.Bai[0]) &&
			helper.Contains(t.Shi, target.Shi[0]) &&
			helper.Contains(t.Gen, target.Gen[0]) {

			prizeRecords = append(prizeRecords, types.PrizeRecord{
				Ticket: t,
				Grade:  0,
				Count:  1,
				Amount: Z1RewardPerTicket,
			})
		}
	case "z3":
		tmp := mapset.NewSet[string]()

		tmp.Add(target.Bai[0])
		tmp.Add(target.Shi[0])
		tmp.Add(target.Gen[0])

		targetSlice := tmp.ToSlice()

		if len(targetSlice) == 2 {
			if helper.ContainsAll(t.Ton, targetSlice...) {

				prizeRecords = append(prizeRecords, types.PrizeRecord{
					Ticket: t,
					Grade:  0,
					Count:  1,
					Amount: Z3RewardPerTicket,
				})
			}
		}
	case "z6":
		tmp := mapset.NewSet[string]()

		tmp.Add(target.Bai[0])
		tmp.Add(target.Shi[0])
		tmp.Add(target.Gen[0])

		targetSlice := tmp.ToSlice()

		if len(targetSlice) == 3 {
			if helper.ContainsAll(t.Ton, targetSlice...) {
				prizeRecords = append(prizeRecords, types.PrizeRecord{
					Ticket: t,
					Grade:  0,
					Count:  1,
					Amount: Z6RewardPerTicket,
				})
			}
		}

	}

	return prizeRecords

}

func (t *Prizer) FetchTarget(index string) (*types.PrizeResult, error) {

	result, err := fetcher_adapter.GetCn8200Adapter().FindPl3ResultByIndex(index)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	var target = Ticket{
		Cat:         "result",
		Bai:         result.Result[0:1],
		Shi:         result.Result[1:2],
		Gen:         result.Result[2:3],
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
