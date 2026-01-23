package types

import (
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"time"
)

type Feedback struct {
	Id          string    `json:"id"`
	Store       *Store    `json:"store"`
	User        *User     `json:"user"`
	Text        string    `json:"text"`
	CreatedAt   string    `json:"created_at"`
	CreatedAtTs int64     `json:"created_at_ts"`
	Status      enum.Enum `json:"status"`
	Resolvable  bool      `json:"resolvable"`
}

type Feedbacks []*Feedback

func FromDbFeedback(x *db.Feedback, user *db.User, store *db.Store, storeOwner *db.User) *Feedback {

	if x == nil {
		return nil
	}

	rsp := Feedback{
		Id:          x.Id,
		Store:       FromDbStore(store, storeOwner),
		User:        FromDbUser(user),
		Text:        x.Text,
		CreatedAt:   x.CreatedAt.Format(time.DateTime),
		CreatedAtTs: x.CreatedAt.Unix(),
		Status:      enum.FeedbackStatus(x.Status),
		Resolvable:  helper.InSlice(x.Status, enum.ResolvableFeedbackStatus),
	}
	return &rsp
}
