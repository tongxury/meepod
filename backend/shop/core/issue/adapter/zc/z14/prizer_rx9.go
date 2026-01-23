package z14

import (
	"gitee.com/meepo/backend/kit/components/sdk/helper/mathd"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/core/issue/types"
)

type Rx9Prizer struct {
}

func (t *Rx9Prizer) Prize(src string, matchTarget types.MatchTarget) (types.PrizeRecords, error) {

	tickets, _, _, err := new(Parser).Parse(src)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	var rsp []types.PrizeRecord
	for _, x := range tickets {
		rsp = append(rsp, x.(*Ticket).prizeRx9(matchTarget)...)
	}

	return nil, nil
}

// 所选择的任意9场竞猜场次中每场比赛的胜平负结果一致即中一等奖
func (t *Ticket) prizeRx9(matchTarget types.MatchTarget) []types.PrizeRecord {

	if len(t.Options) < 9 {
		return nil
	}

	var prizeRecords types.PrizeRecords

	matchesMap := matchTarget.Matches.AsMap()

	var unhitMatches []Options
	var unhitDanMatches []Options
	var hitMatches []Options
	var hitDanMatches []Options

	for matchId, matchOptions := range t.Options {
		if matchResult, found := matchesMap[matchId]; found {

			if matchOptions.Odds.Items.Contains(matchResult.Result.Value) {
				if matchOptions.Dan {
					hitDanMatches = append(hitDanMatches, matchOptions)
				} else {
					hitMatches = append(hitMatches, matchOptions)
				}
			} else {
				if matchOptions.Dan {
					unhitDanMatches = append(unhitDanMatches, matchOptions)
				} else {
					unhitMatches = append(unhitMatches, matchOptions)

				}
			}
		}
	}

	// 胆里面只要有一个没中就不可能中奖
	if len(unhitDanMatches) > 0 {
		return nil
	}

	if len(hitMatches)+len(hitDanMatches) < 9 {
		return nil
	}

	count := mathd.Cmn(len(hitMatches), 9-len(hitDanMatches))

	prizeRecords = append(prizeRecords, types.PrizeRecord{
		Grade: 1, Count: count, Amount: matchTarget.PrizeGrades.GetGradeAmount(1),
	})

	return prizeRecords
}
