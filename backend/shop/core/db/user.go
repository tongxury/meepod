package db

import (
	"context"
	"fmt"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"github.com/go-pg/pg/v10"
	"time"
)

type User struct {
	tableName struct{} `pg:"t_users,alias:u"`
	Id        string
	Phone     string
	//Password  string
	CreatedAt time.Time `pg:"default:now()"`
	Status    string
	Extra     UserExtra
}

type UserExtra struct {
	Icon            string `json:"icon"`
	Nickname        string `json:"nickname"`
	WechatPayQrcode string `json:"wechat_pay_qrcode"`
	AliPayQrcode    string `json:"ali_pay_qrcode"`
}

//
//func (t *User) Icon() string {
//
//	icon, err := FindInExtra(t.Extra, "icon")
//	if err != nil {
//		return ""
//	}
//
//	return icon
//}
//
//func (t *User) Nickname() string {
//
//	val, err := FindInExtra(t.Extra, "nickname")
//	if err != nil {
//		return ""
//	}
//
//	return val
//}
//func (t *User) Wechat() string {
//
//	val, err := FindInExtra(t.Extra, "wechat")
//	if err != nil {
//		return ""
//	}
//
//	return val
//}
//
//func (t *User) Alipay() string {
//
//	val, err := FindInExtra(t.Extra, "alipay")
//	if err != nil {
//		return ""
//	}
//
//	return val
//}

type Users []*User

func (ts Users) AsMap() map[string]*User {
	rsp := make(map[string]*User, len(ts))

	for _, t := range ts {
		rsp[t.Id] = t
	}

	return rsp
}

type ListUsersParams struct {
	//StoreId    string
	Keyword    string
	Id         string
	Phone      string
	PhoneLike  string
	Page, Size int64
}

func (t *User) List(ctx context.Context, params ListUsersParams) (Users, int64, error) {

	var tmp Users
	q := comp.SDK().Postgres().Model(&tmp).Context(ctx)

	if params.Keyword != "" {
		q = q.Where("id like ? ", "%"+params.Keyword+"%")
	}

	if params.Id != "" {
		if conv.Int(params.Id) > 0 {
			q = q.Where("id = ?", params.Id)
		}
	}

	if params.PhoneLike != "" {
		q = q.Where("phone like ?", "%"+params.PhoneLike+"%")
	}
	if params.Phone != "" {
		q = q.Where("phone = ?", params.Phone)
	}

	if params.Size > 0 {
		page := params.Page
		if page <= 0 {
			page = 1
		}

		q = q.Limit(int(params.Size)).Offset(int((page - 1) * params.Size))
	}
	count, err := q.OrderExpr("id desc").SelectAndCount()

	if err != nil {
		return nil, 0, err
	}

	return tmp, int64(count), nil
}

func (t *User) CreateNX(ctx context.Context) (*User, error) {

	_, err := comp.SDK().Postgres().Model(t).Context(ctx).
		OnConflict("(phone) do nothing").
		Where("phone = ?", t.Phone).
		SelectOrInsert(t)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (t *User) FindPyPhone(ctx context.Context, phone string) (*User, error) {
	var tmp Users
	err := comp.SDK().Postgres().Model(&tmp).Context(ctx).
		Where("phone = ?", phone).
		Select()

	if err != nil {
		return nil, err
	}

	if len(tmp) == 0 {
		return nil, nil
	}

	return tmp[0], nil
}

func (t *User) FindById(ctx context.Context, id string) (*User, error) {

	var tmp Users
	err := comp.SDK().Postgres().Model(&tmp).Context(ctx).
		Where("id = ?", id).
		Select()

	if err != nil {
		return nil, err
	}

	if len(tmp) == 0 {
		return nil, nil
	}

	return tmp[0], nil
}

func (t *User) ListByIds(ctx context.Context, ids []string) (Users, error) {

	if len(ids) == 0 {
		return nil, nil
	}

	var rsp Users

	q := comp.SDK().Postgres().Model(&rsp).Context(ctx).
		WhereIn("id in (?)", ids)

	err := q.Select()

	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func (t *User) Update(ctx context.Context, tx *pg.Tx, id, field, value string) (bool, error) {

	update, err := tx.Model((*User)(nil)).Context(ctx).
		Where("id = ?", id).
		Set("extra =  jsonb_set(extra, ?, ?)", fmt.Sprintf("{%s}", field), fmt.Sprintf("\"%s\"", value)).
		Update()

	if err != nil {
		return false, err
	}

	return update.RowsAffected() > 0, nil
}
