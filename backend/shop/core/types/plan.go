package types

import (
	"gitee.com/meepo/backend/kit/components/sdk/helper/timed"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/issue/adapter"
)

type PlanForm struct {
	ItemId   string `binding:"required"`
	Content  string `binding:"required"`
	Multiple int64
}

type Plan struct {
	Id           string    `json:"id"`
	Item         *Item     `json:"item"`
	Issue        *Issue    `json:"issue"`
	Content      string    `json:"content"`
	Tickets      []any     `json:"tickets"`
	SplitTickets []any     `json:"split_tickets"`
	Multiple     int64     `json:"multiple"`
	Amount       float64   `json:"amount"`
	User         *User     `json:"user"`
	CreatedAt    string    `json:"created_at"`
	CreatedAtTs  int64     `json:"created_at_ts"`
	Status       enum.Enum `json:"status"`
}

type Plans []*Plan

func FromDbPlan(plan *db.Plan, item *db.Item, planCreator *db.User, issue *db.Issue) *Plan {

	tickets, splitTickets, _, _ := adapter.Parse(item.Id, plan.Content)

	rsp := Plan{
		Id:           plan.Id,
		Item:         FromDbItem(item),
		Issue:        FromDbIssue(issue, item),
		Content:      plan.Content,
		Tickets:      tickets,
		SplitTickets: splitTickets,
		Multiple:     plan.Multiple,
		Amount:       plan.Amount,
		User:         FromDbUser(planCreator),
		CreatedAtTs:  plan.CreatedAt.Unix(),
		CreatedAt:    timed.SmartTime(plan.CreatedAt.Unix()),
		Status:       enum.PlanStatus(plan.Status),
	}

	return &rsp
}
