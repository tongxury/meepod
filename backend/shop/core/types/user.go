package types

import (
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/kit/components/sdk/helper/timed"
	"gitee.com/meepo/backend/kit/services/util/oss"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
)

type UserProfile struct {
	User *User `json:"user"`
	//Account *Account `json:"account"`
}

type User struct {
	Id              string    `json:"id"`
	Phone           string    `json:"phone"`
	Nickname        string    `json:"nickname"`
	Wechat          string    `json:"wechat"`
	Alipay          string    `json:"alipay"`
	WechatPayQrcode string    `json:"wechat_pay_qrcode"`
	AlipayQrcode    string    `json:"ali_pay_qrcode"`
	Desc            string    `json:"desc"`
	Icon            string    `json:"icon"`
	CreatedAt       string    `json:"created_at"`
	CreatedAtTs     int64     `json:"created_at_ts"`
	Status          enum.Enum `json:"status"`
	Tags            Tags      `json:"tags"`
}

type Users []*User

type UserOptions struct {
	ShowConnect bool
}

func FromDbUser(user *db.User, options ...UserOptions) *User {

	if user == nil {
		return nil
	}

	rsp := User{
		Id: user.Id,
		//Phone: user.Phone[0:3] + "****",
		Phone:    user.Phone,
		Nickname: helper.OrString(user.Extra.Nickname, fmt.Sprintf("用户%s", user.Id)),
		//Wechat:          user.Extra.WechatPayQrCode,
		//Alipay:          user.Extra.Alipay,
		AlipayQrcode:    oss.Resource(user.Extra.AliPayQrcode),
		WechatPayQrcode: oss.Resource(user.Extra.WechatPayQrcode),
		Desc:            "",
		Icon:            oss.Resource(helper.OrString(user.Extra.Icon, enum.DefaultUserIcon)),
		CreatedAt:       timed.SmartTime(user.CreatedAt.Unix()),
		CreatedAtTs:     user.CreatedAt.Unix(),
		Tags:            nil,
	}
	//
	//if len(options) > 0 {
	//	if options[0].ShowConnect {
	//		rsp.Phone = user.Phone
	//		rsp.Wechat = user.Extra.Wechat
	//		rsp.Alipay = user.Extra.Alipay
	//	}
	//}

	return &rsp
}
