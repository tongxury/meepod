package types

import (
	"time"
)

type Match struct {
	Id           string
	League       string
	HomeTeam     string
	HomeTeamTag  string
	GuestTeam    string
	GuestTeamTag string
	Category     string
	Issue        string
	StartAt      time.Time
	CloseAt      time.Time
	CreatedAt    time.Time
	Status       string
	Odds         Odds
	RealOdds     Odds
	Result       Result
	RCount       int
}

type Matches []*Match

func (ts Matches) AsMap() map[string]*Match {
	tmp := make(map[string]*Match, len(ts))

	for _, x := range ts {
		tmp[x.Id] = x
	}

	return tmp
}

type OddsItem struct {
	Name   string  `json:"name,omitempty"`
	Result string  `json:"result,omitempty"` // 3 1 0
	Value  float64 `json:"value,omitempty"`
	//Goal   int     `json:"goal"` // 让球数  只有主胜有值
}

type OddsItems []OddsItem

func (ts OddsItems) Contains(resultValue string) bool {
	for _, t := range ts {
		if t.Result == resultValue {
			return true
		}
	}

	return false
}

type Result struct {
	Value     string `json:"value,omitempty"`
	HalfValue string `json:"half_value,omitempty"`
	Goals     string `json:"goals,omitempty"`
	HalfGoals string `json:"half_goals,omitempty"`
}

type Odds struct {
	Items             OddsItems `json:"items,omitempty"`
	RItems            OddsItems `json:"r_items,omitempty"`
	ScoreVictoryItems OddsItems `json:"score_victory_items,omitempty"`
	ScoreDogfallItems OddsItems `json:"score_dogfall_items,omitempty"`
	ScoreDefeatItems  OddsItems `json:"score_defeat_items,omitempty"`
	GoalsItems        OddsItems `json:"goals_items,omitempty"`
	HalfFullItems     OddsItems `json:"half_full_items,omitempty"`
}

type MatchTarget struct {
	Matches     Matches
	PrizeGrades PrizeGrades
}
