package types

import (
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/issue/types"
	"time"
)

type Match struct {
	Id           string       `json:"id"`
	League       string       `json:"league"`
	HomeTeam     string       `json:"home_team"`
	HomeTeamTag  string       `json:"home_team_tag"`
	GuestTeam    string       `json:"guest_team"`
	GuestTeamTag string       `json:"guest_team_tag"`
	Category     enum.Enum    `json:"category"`
	Issue        *Issue       `json:"issue"`
	StartAt      string       `json:"start_at"`
	StartAtTs    int64        `json:"start_at_ts"`
	CloseAtTs    int64        `json:"close_at_ts"`
	Result       types.Result `json:"result"`
	Status       enum.Enum    `json:"status"`
	Odds         types.Odds   `json:"odds"`
	RealOdds     types.Odds   `json:"real_odds"`
}

type Matches []*Match

func FromDbMatch(x *db.Match, issue *db.Issue, item *db.Item) *Match {

	y := Match{
		Id:           x.Id,
		League:       x.League,
		HomeTeam:     x.HomeTeam,
		HomeTeamTag:  x.HomeTeamTag,
		GuestTeam:    x.GuestTeam,
		GuestTeamTag: x.GuestTeamTag,
		Category:     enum.MatchCategory(x.Category),
		//Issue:        FromDbIssue(issue, item),
		StartAt:   x.StartAt.In(time.Local).Format("01-02 15:04"),
		StartAtTs: x.StartAt.Unix(),
		CloseAtTs: x.CloseAt.Unix(),
		Result:    x.Result,
		Status:    enum.MatchStatus(x.Status),
		Odds:      x.Odds,
		RealOdds:  x.RealOdds,
	}

	return &y
}
