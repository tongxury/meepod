package dlt

import (
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/kit/components/sdk/helper/mathd"
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

// 一等奖：投注号码与当期开奖号码全部相同(顺序不限，下同)，即中奖；
// 二等奖：投注号码与当期开奖号码中的五个前区号码及任意一个后区号码相同，即中奖；
// 三等奖：投注号码与当期开奖号码中的五个前区号码相同，即中奖；
// 四等奖：投注号码与当期开奖号码中的任意四个前区号码及两个后区号码相同，即中奖；
// 五等奖：投注号码与当期开奖号码中的任意四个前区号码及任意一个后区号码相同，即中奖；
// 六等奖：投注号码与当期开奖号码中的任意三个前区号码及两个后区号码相同，即中奖；
// 七等奖：投注号码与当期开奖号码中的任意四个前区号码相同，即中奖；
// 八等奖：投注号码与当期开奖号码中的任意三个前区号码及任意一个后区号码相同，或者任意两个前区号码及两个后区号码相同，即中奖；
// 九等奖：投注号码与当期开奖号码中的任意三个前区号码相同，或者任意一个前区号码及两个后区号码相同，或者任意两个前区号码及任意一个后区号码相同，或者两个后区号码相同，即中奖。
func (t *Ticket) prize(target *Ticket) types.PrizeRecords {
	var rsp types.PrizeRecords

	redHits, redUnHits := helper.Intersect(t.Red, target.Red)
	dRedHits, dRedUnHits := helper.Intersect(t.DRed, target.Red)
	blueHits, _ := helper.Intersect(t.Blue, target.Blue)

	//一等奖
	if len(dRedUnHits) == 0 && len(redHits)+len(dRedHits) >= 5 && len(blueHits) >= 2 {
		rsp = append(rsp, types.PrizeRecord{
			Grade:  1,
			Count:  1,
			Amount: target.PrizeGrades.GetGradeAmount(1),
		})
		// 二等奖
	} else if len(dRedUnHits) == 0 && len(redHits)+len(dRedHits) >= 5 && len(blueHits) >= 1 {
		rsp = append(rsp, types.PrizeRecord{
			Grade:  2,
			Count:  len(t.Blue) - 1,
			Amount: target.PrizeGrades.GetGradeAmount(2),
		})
		// 三等奖
	} else if len(dRedUnHits) == 0 && len(redHits)+len(dRedHits) >= 5 {
		rsp = append(rsp, types.PrizeRecord{
			Grade:  3,
			Count:  len(t.Blue),
			Amount: target.PrizeGrades.GetGradeAmount(3),
		})
	} else {
		// 四等奖
		if len(dRedUnHits) <= 1 && len(redHits)+len(dRedHits) >= 4 && len(blueHits) >= 2 {
			rsp = append(rsp, types.PrizeRecord{
				Grade:  4,
				Count:  mathd.Cmn(len(redUnHits), 5-(len(t.DRed)+len(redHits))),
				Amount: target.PrizeGrades.GetGradeAmount(4),
			})
			// 五等奖
		} else if len(dRedHits) <= 1 && len(redHits)+len(dRedHits) >= 4 && len(blueHits) >= 1 {
			rsp = append(rsp, types.PrizeRecord{
				Grade:  5,
				Count:  mathd.Cmn(len(redUnHits), 5-(len(t.DRed)+len(redHits))) * (len(t.Blue) - 1),
				Amount: target.PrizeGrades.GetGradeAmount(5),
			})
			// 六等奖
		} else if len(dRedUnHits) <= 2 && len(redHits)+len(dRedHits) >= 3 && len(blueHits) >= 2 {
			rsp = append(rsp, types.PrizeRecord{
				Grade:  6,
				Count:  mathd.Cmn(len(redUnHits), 5-(len(t.DRed)+len(redHits))),
				Amount: target.PrizeGrades.GetGradeAmount(6),
			})
			// 七等奖
		} else if len(dRedUnHits) <= 1 && len(redHits)+len(dRedHits) >= 4 {
			rsp = append(rsp, types.PrizeRecord{
				Grade:  7,
				Count:  mathd.Cmn(len(redUnHits), 5-(len(t.DRed)+len(redHits))) * mathd.Cmn(len(t.Blue), 2),
				Amount: target.PrizeGrades.GetGradeAmount(7),
			})
			// 八等奖
		} else if len(dRedUnHits) <= 2 && len(redHits)+len(dRedHits) >= 3 && len(blueHits) >= 1 {
			rsp = append(rsp, types.PrizeRecord{
				Grade:  8,
				Count:  mathd.Cmn(len(redUnHits), 5-(len(t.DRed)+len(redHits))) * (len(t.Blue) - 1),
				Amount: target.PrizeGrades.GetGradeAmount(8),
			})

			// 八等奖
		} else if len(dRedUnHits) <= 3 && len(redHits)+len(dRedHits) >= 2 && len(blueHits) >= 2 {
			rsp = append(rsp, types.PrizeRecord{
				Grade:  8,
				Count:  mathd.Cmn(len(redUnHits), 5-(len(t.DRed)+len(redHits))),
				Amount: target.PrizeGrades.GetGradeAmount(8),
			})

			// 九等奖
		} else if len(dRedUnHits) <= 2 && len(redHits)+len(dRedHits) >= 3 {
			rsp = append(rsp, types.PrizeRecord{
				Grade:  9,
				Count:  mathd.Cmn(len(redUnHits), 5-(len(t.DRed)+len(redHits))) * mathd.Cmn(len(t.Blue), 2),
				Amount: target.PrizeGrades.GetGradeAmount(9),
			})
			// 九等奖
		} else if len(dRedUnHits) <= 4 && len(redHits)+len(dRedHits) >= 1 && len(blueHits) >= 2 {
			rsp = append(rsp, types.PrizeRecord{
				Grade:  9,
				Count:  mathd.Cmn(len(redUnHits), 5-(len(t.DRed)+len(redHits))),
				Amount: target.PrizeGrades.GetGradeAmount(9),
			})
			// 九等奖
		} else if len(dRedUnHits) <= 3 && len(redHits)+len(dRedHits) >= 2 && len(blueHits) >= 1 {
			rsp = append(rsp, types.PrizeRecord{
				Grade:  9,
				Count:  mathd.Cmn(len(redUnHits), 5-(len(t.DRed)+len(redHits))) * (len(t.Blue) - 1),
				Amount: target.PrizeGrades.GetGradeAmount(9),
			})
			// 九等奖
		} else if len(blueHits) >= 2 {
			rsp = append(rsp, types.PrizeRecord{
				Grade:  9,
				Count:  mathd.Cmn(len(redUnHits), 5-(len(t.DRed)+len(redHits))),
				Amount: target.PrizeGrades.GetGradeAmount(9),
			})

		}
	}

	return rsp
}

func (t *Prizer) FetchTarget(index string) (*types.PrizeResult, error) {
	result, err := fetcher_adapter.GetCn8200Adapter().FindDltResultByIndex(index)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	var ticket = Ticket{
		Red:         result.Red,
		Blue:        result.Blue,
		PrizeGrades: result.PrizeGrades,
	}

	return &types.PrizeResult{
		Issue:   index,
		Result:  conv.S2J(Tickets{&ticket}),
		PrizeAt: result.PrizeAt(PrizeHour, PrizeMinute),
	}, nil

}
