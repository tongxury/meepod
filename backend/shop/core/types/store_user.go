package types

import (
	"gitee.com/meepo/backend/shop/core/db"
)

type StoreUser struct {
	User   *User  `json:"user"`
	Remark string `json:"remark,omitempty"`
}

type StoreUsers []*StoreUser

func FromDbStoreUser(storeUser *db.StoreUser, user *db.User) *StoreUser {

	if storeUser == nil {
		return nil
	}

	rsp := StoreUser{
		User:   FromDbUser(user),
		Remark: storeUser.Extra.Remark,
	}
	return &rsp
}
