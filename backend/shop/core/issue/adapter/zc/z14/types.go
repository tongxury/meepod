package z14

import (
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/shop/core/issue/types"
)

// 最大金额6000
// 最大倍数50
type Options struct {
	Odds types.Odds `json:"odds,omitempty"`
	Dan  bool       `json:"dan,omitempty"`
}

type Ticket struct {
	Key     string `json:"key"`
	Matches map[string]struct {
		Id       string `json:"id"`
		RCount   int    `json:"r_count"`
		HomeTeam string `json:"home_team"`
	} `json:"matches"`
	Options map[string]Options `json:"options"`
	Amount  float64            `json:"amount"`
}

type Tickets []*Ticket

// 无须拆票
func (t *Ticket) Split() Tickets {
	return Tickets{t}
}

func (t *Ticket) GetAmount(min int) (int64, float64) {

	matchOptionCounts := make(map[string]int, len(t.Options))
	for id, x := range t.Options {
		matchOptionCounts[id] = len(x.Odds.Items)
	}

	var danMatchIds []string
	var notDanMatchIds []string

	for k, x := range t.Options {
		if x.Dan {
			danMatchIds = append(danMatchIds, k)
		} else {
			notDanMatchIds = append(notDanMatchIds, k)
		}
	}

	var danCount = 1
	for _, x := range danMatchIds {
		danCount *= matchOptionCounts[x]
	}

	var count = 0
	for _, _matchIds := range helper.Combine(notDanMatchIds, min-len(danMatchIds)) {

		var tmpCount = 1
		for _, id := range _matchIds {
			tmpCount *= matchOptionCounts[id]
		}

		count += tmpCount
	}

	count *= danCount

	return int64(count), float64(count * 2)
}
