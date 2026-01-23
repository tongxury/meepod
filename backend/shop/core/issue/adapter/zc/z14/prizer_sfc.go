package z14

import (
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/core/issue/types"
)

type SfcPrizer struct {
}

func (t *SfcPrizer) Prize(src string, matchTarget types.MatchTarget) (types.PrizeRecords, error) {

	tickets, _, _, err := new(Parser).Parse(src)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	var rsp []types.PrizeRecord
	for _, x := range tickets {
		rsp = append(rsp, x.(*Ticket).prizeSfc(matchTarget)...)
	}

	return nil, nil
}

// 一等奖 猜中全部14场比赛的胜平负结果；
// 二等奖 猜中其中13场比赛的胜平负结果。
func (t *Ticket) prizeSfc(matchTarget types.MatchTarget) []types.PrizeRecord {

	if len(t.Options) < 14 {
		return nil
	}

	var prizeRecords types.PrizeRecords

	matchesMap := matchTarget.Matches.AsMap()

	var unhitMatch *Options

	for matchId, matchOptions := range t.Options {
		if matchResult, found := matchesMap[matchId]; found {

			if matchOptions.Odds.Items.Contains(matchResult.Result.Value) {
				continue
			}

			// 超过两个没命中就不会中奖
			if unhitMatch != nil {
				return prizeRecords
			}

			tmp := matchOptions
			unhitMatch = &tmp
		}
	}

	if unhitMatch == nil {

		prizeRecords = append(prizeRecords, types.PrizeRecord{
			Grade: 1, Count: 1, Amount: matchTarget.PrizeGrades.GetGradeAmount(1),
		})
	} else {
		prizeRecords = append(prizeRecords, types.PrizeRecord{
			Grade: 2, Count: len(unhitMatch.Odds.Items), Amount: matchTarget.PrizeGrades.GetGradeAmount(2),
		})
	}

	return prizeRecords
}
