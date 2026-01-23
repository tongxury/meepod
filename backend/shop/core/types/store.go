package types

import (
	"time"

	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/kit/components/sdk/helper/timed"
	"gitee.com/meepo/backend/kit/services/util/oss"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
)

type StoreParams struct {
	StoreName       string `json:"store_name" binding:"required"`
	Phone           string `json:"phone" binding:"required"`
	Username        string `json:"username" binding:"required"`
	Email           string `json:"email" binding:"required"`
	IdCardFront     string `json:"id_card_front" `
	IdCardBack      string `json:"id_card_back" `
	IdCardHandled   string `json:"id_card_handled" `
	IdCardNo        string `json:"id_card_no" `
	IdCardFrom      string `json:"id_card_from" `
	IdCardTo        string `json:"id_card_to" `
	StoreFront      string `json:"store_front" `
	StoreInSide     string `json:"store_in_side" `
	Loc             string `json:"loc" `
	Address         string `json:"address" `
	SalesCard       string `json:"sales_card" `
	BankAccountName string `json:"bank_account_name" `
	BankPhone       string `json:"bank_phone" `
	BankName        string `json:"bank_name" `
	BankCardFront   string `json:"bank_card_front" `
	BankCardBack    string `json:"bank_card_back" `
	BankAccount     string `json:"bank_account" `
	BankBranch      string `json:"bank_branch" `
	//BankLoc         string `json:"bank_loc" binding:"required"`
}

type Store struct {
	Id              string    `json:"id"`
	Name            string    `json:"name"`
	Owner           *User     `json:"owner"`
	CreatedAt       string    `json:"created_at"`
	CreatedAtTs     int64     `json:"created_at_ts"`
	Icon            string    `json:"icon"`
	WechatImage     string    `json:"wechat_image"`
	StoreFront      string    `json:"store_front"`
	Loc             string    `json:"loc"`
	Address         string    `json:"address"`
	Phone           string    `json:"phone"`
	Username        string    `json:"username"`
	Email           string    `json:"email"`
	IdCardNo        string    `json:"id_card_no"`
	IdCardFrom      string    `json:"id_card_from"`
	IdCardTo        string    `json:"id_card_to"`
	IdCardFront     string    `json:"id_card_front"`
	IdCardBack      string    `json:"id_card_back"`
	IdCardHandled   string    `json:"id_card_handled"`
	StoreInSide     string    `json:"store_in_side"`
	SalesCard       string    `json:"sales_card"`
	BankAccountName string    `json:"bank_account_name"`
	BankName        string    `json:"bank_name"`
	BankPhone       string    `json:"bank_phone"`
	BankCardFront   string    `json:"bank_card_front"`
	BankCardBack    string    `json:"bank_card_back"`
	BankAccount     string    `json:"bank_account"`
	BankBranch      string    `json:"bank_branch"`
	Status          enum.Enum `json:"status"`
	MemberLevel     enum.Enum `json:"member_level"`
	MemberUntil     string    `json:"member_until"`
	SelectedItemIds []string  `json:"selected_item_ids"`
	Items           Items     `json:"items"`
	Balance         float64   `json:"balance"`
	Notice          string    `json:"notice"`
}

type Stores []*Store

func FromDbStore(store *db.Store, storeOwner *db.User, extra ...db.ExtraMaps) *Store {
	if store == nil {
		return nil
	}

	y := &Store{
		Id:          store.Id,
		Name:        store.Name,
		Owner:       FromDbUser(storeOwner),
		CreatedAt:   timed.SmartTime(store.CreatedAt.Unix()),
		CreatedAtTs: store.CreatedAt.Unix(),
		Icon:        oss.Resource(helper.OrString(store.Extra.Icon, enum.DefaultStoreIcon)),
		WechatImage: oss.Resource(store.Extra.WechatImage),

		StoreFront:      store.Extra.StoreFront,
		Loc:             store.Extra.Loc,
		Address:         store.Extra.Address,
		Username:        store.Extra.Username,
		Email:           store.Extra.Email,
		IdCardNo:        store.Extra.IdCardNo,
		IdCardFrom:      store.Extra.IdCardFrom,
		IdCardTo:        store.Extra.IdCardTo,
		IdCardFront:     store.Extra.IdCardFront,
		IdCardBack:      store.Extra.IdCardBack,
		IdCardHandled:   store.Extra.IdCardHandled,
		StoreInSide:     store.Extra.StoreInSide,
		SalesCard:       store.Extra.SalesCard,
		BankAccountName: store.Extra.BankAccountName,
		BankName:        store.Extra.BankName,
		BankPhone:       store.Extra.BankPhone,
		BankCardFront:   store.Extra.BankCardFront,
		BankCardBack:    store.Extra.BankCardBack,
		BankAccount:     store.Extra.BankAccount,
		BankBranch:      store.Extra.BankBranch,
		Status:          enum.StoreStatus(store.Status),
		MemberLevel:     enum.MemberLevel(store.Member.Level),
		MemberUntil: helper.Choose(store.Member.Until == 0, "",
			time.Unix(store.Member.Until, 0).Format(time.DateOnly)),
		SelectedItemIds: store.Settings.ItemIds,
		Items:           nil,
		Balance:         0,
		Notice:          store.Settings.Notice,
	}

	if len(extra) > 0 {

		for _, id := range y.SelectedItemIds {
			y.Items = append(y.Items, FromDbItem(extra[0].Items[id]))

		}

	}
	return y
}
