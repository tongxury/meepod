package types

import (
	"gitee.com/meepo/backend/shop/core/enum"
)

type OrderGroup struct {
	Id            string    `json:"id"`
	Plan          *Plan     `json:"plan"`
	User          *User     `json:"user"`
	Store         *Store    `json:"store"`
	ToStore       *Store    `json:"to_store"`
	Amount        float64   `json:"amount"`
	Volume        int64     `json:"volume"`
	VolumeOrdered int64     `json:"volume_ordered"`
	Floor         int64     `json:"floor"`
	RewardRate    float64   `json:"reward_rate"`
	Remark        string    `json:"remark"`
	CreatedAt     string    `json:"created_at"`
	CreatedAtTs   int64     `json:"created_at_ts"`
	Status        enum.Enum `json:"status"`
	Tags          Tags      `json:"tags"`
	Joinable      bool      `json:"joinable"`
	JoinerCount   int64     `json:"joiner_count"`
	Rejectable    bool      `json:"rejectable"`
	Acceptable    bool      `json:"acceptable"`
	Ticketable    bool      `json:"ticketable"`
	Switchable    bool      `json:"switchable"`
	TicketImages  []string  `json:"ticket_images"`
	Prized        bool      `json:"prized"`
	NeedUpload    bool      `json:"need_upload"`
}

type OrderGroups []*OrderGroup

type JoinGroupOrderForm struct {
	Volume int64 `binding:"required"`
}

// todo
//func FromDbGroup(group *db.OrderGroup, user *db.User) *OrderGroup {
//
//	if group == nil {
//		return nil
//	}
//
//	rsp := OrderGroup{
//		Id: group.Id,
//		//Plan:          FromDbPlan(dbPlan, ),
//		User: FromDbUser(user),
//		//Store:         FromDbStore(store, ),
//		Volume:        group.Volume,
//		VolumeOrdered: group.VolumeOrdered,
//		Floor:         group.Floor,
//		RewardRate:    group.RewardRate,
//		Remark:        group.Remark,
//		CreatedAt:     timed.SmartTime(group.CreatedAt.Unix()),
//		CreatedAtTs:   group.CreatedAt.Unix(),
//		MStatus:        enum.OrderGroupStatus(group.MStatus),
//		Tags:          nil,
//		Joinable:      helper.InSlice(group.MStatus, enum.JoinableOrderGroupStatus),
//		//JoinerCount:   joinerCount,
//	}
//
//	return &rsp
//}
