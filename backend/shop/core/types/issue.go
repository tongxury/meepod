package types

import (
	"gitee.com/meepo/backend/kit/components/sdk/helper/mathd"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/issue/adapter"
	"time"
)

type Issue struct {
	Item          *Item     `json:"item"`
	Index         string    `json:"index"`
	Result        any       `json:"result"`
	PrizedAt      int64     `json:"prized_at"`
	StartedAt     int64     `json:"started_at"`
	CloseAt       int64     `json:"close_at"`
	CloseAtTime   time.Time `json:"close_at_time"`
	CloseTimeLeft int64     `json:"close_time_left"`
	PrizeTimeLeft int64     `json:"prize_time_left"`
	Status        enum.Enum `json:"status"`
	Extra         []any     `json:"extra"`
	//EmptyExtra    bool      `json:"-"`
}

type Issues []*Issue

func FromDbIssue(issue *db.Issue, item *db.Item) *Issue {

	if issue == nil {
		return nil
	}

	var result any
	if issue.Result != "" && issue.Result != "[]" {
		results, _, _, _ := adapter.Parse(item.Id, issue.Result)
		if len(results) > 0 {
			result = results[0]
		}
	}

	closeAtTime := issue.CloseAt.Add(-5 * time.Minute)
	prizedAt := issue.PrizedAt.Unix()
	//
	//var extraTap map[string]any
	//
	//err := conv.J2M(issue.Extra, &extraTap)
	//if err != nil {
	//	slf.WithError(err).Errorw("J2M err", slf.Reflect("src", issue.Extra))
	//}

	rsp := Issue{
		Item:          FromDbItem(item),
		Index:         issue.Index,
		CloseTimeLeft: mathd.Max(int64(0), closeAtTime.Unix()-time.Now().Unix()),
		PrizeTimeLeft: mathd.Max(int64(0), prizedAt-time.Now().Unix()),
		Result:        result,
		PrizedAt:      issue.PrizedAt.Unix(),
		StartedAt:     issue.StartedAt.Unix(),
		CloseAt:       closeAtTime.Unix(),
		CloseAtTime:   closeAtTime,
		Status:        enum.IssueStatus(issue.Status),
		//Extra:         extraTap,
	}

	return &rsp
}
