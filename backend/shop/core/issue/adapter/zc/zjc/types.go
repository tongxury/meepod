package zjc

import (
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/shop/core/issue/types"
	"strings"
)

// 最大金额6000
// 最大倍数50

//type Options struct {
//	Spf          types.OddsItems `json:"spf,omitempty"`
//	Rspf         types.OddsItems `json:"rspf,omitempty"`
//	Goals        types.OddsItems `json:"goals,omitempty"`
//	ScoreVictory types.OddsItems `json:"score_victory,omitempty"`
//	ScoreDogfall types.OddsItems `json:"score_dogfall,omitempty"`
//	ScoreDefeat  types.OddsItems `json:"score_defeat,omitempty"`
//	HalfFull     types.OddsItems `json:"half_full,omitempty"`
//	Dan          types.OddsItems `json:"dan,omitempty"`
//}

type Options struct {
	Odds types.Odds `json:"odds,omitempty"`
	Dan  bool       `json:"dan,omitempty"`
}

type Ticket struct {
	Key     string   `json:"key"`
	Modes   []string `json:"modes"`
	Matches map[string]struct {
		Id       string `json:"id"`
		RCount   int    `json:"r_count"`
		HomeTeam string `json:"home_team"`
	} `json:"matches"`
	Options map[string]Options `json:"options"`
	Amount  float64            `json:"amount"`
}

type Tickets []*Ticket

func (t *Ticket) Split() Tickets {
	// todo
	return Tickets{t}
}

func (t *Ticket) CalcAmount() (int64, float64) {

	// todo 胆
	matchOptionCounts := make(map[string]int, len(t.Options))
	for id, x := range t.Options {
		matchOptionCounts[id] = len(x.Odds.Items) + len(x.Odds.RItems) + len(x.Odds.GoalsItems) + len(x.Odds.ScoreVictoryItems) +
			len(x.Odds.ScoreDogfallItems) + len(x.Odds.ScoreDefeatItems) + len(x.Odds.HalfFullItems)
	}

	var matchIds []string
	for k, _ := range t.Options {
		matchIds = append(matchIds, k)
	}

	factors := map[string][]int{
		"2-1":   {2},
		"3-1":   {3},
		"3-3":   {2},
		"3-4":   {2, 4},
		"4-1":   {4},
		"4-4":   {3},
		"4-5":   {3, 4},
		"4-6":   {2},
		"4-11":  {2, 3, 4},
		"5-1":   {5},
		"5-5":   {4},
		"5-6":   {4, 5},
		"5-10":  {3},
		"5-16":  {3, 4, 5},
		"5-20":  {2, 3},
		"5-26":  {2, 3, 4, 5},
		"6-1":   {6},
		"6-6":   {5},
		"6-7":   {5, 6},
		"6-15":  {4},
		"6-20":  {3},
		"6-22":  {4, 5, 6},
		"6-35":  {3, 4},
		"6-42":  {3, 4, 5, 6},
		"6-50":  {2, 3, 4},
		"6-57":  {2, 3, 4, 5, 6},
		"7-1":   {7},
		"7-7":   {6},
		"7-8":   {6, 7},
		"7-21":  {5},
		"7-35":  {4},
		"7-120": {2, 3, 4, 5, 6, 7},
		"8-1":   {8},
		"8-8":   {7},
		"8-9":   {7, 8},
		"8-28":  {6},
		"8-56":  {5},
		"8-70":  {4},
		"8-247": {2, 3, 4, 5, 6, 7, 8},
	}

	var count int

	for _, mode := range t.Modes {

		parts := strings.Split(mode, "-")
		m := conv.Int(parts[0])

		for _, _matchIds := range helper.Combine(matchIds, m) {

			for _, factor := range factors[mode] {

				for _, ___matchIds := range helper.Combine(_matchIds, factor) {
					var tmpCount = 1
					for _, id := range ___matchIds {
						tmpCount *= matchOptionCounts[id]
					}

					count += tmpCount
				}

			}

		}
	}

	return int64(count), float64(2 * count)
}
