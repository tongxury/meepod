package db

import (
	"context"
	"fmt"
	"gitee.com/meepo/backend/kit/components/comp"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-pg/pg/v10"
	"time"
)

type StoreUser struct {
	tableName struct{} `pg:"t_store_users"`
	Id        string
	StoreId   string
	UserId    string
	CreatedAt time.Time
	Status    string
	Extra     StoreUserExtra
}

type StoreUserExtra struct {
	Remark string `json:"remark"`
}

type StoreUsers []*StoreUser

func (ts StoreUsers) Ids() ([]string, []string) {

	tmp1 := mapset.NewSet[string]()
	tmp2 := mapset.NewSet[string]()

	for _, t := range ts {
		tmp1.Add(t.StoreId)
		tmp2.Add(t.UserId)
	}

	return tmp1.ToSlice(), tmp2.ToSlice()

}

func (ts StoreUsers) AsUserIdMap() map[string]*StoreUser {
	rsp := make(map[string]*StoreUser, len(ts))

	for _, t := range ts {
		rsp[t.UserId] = t
	}

	return rsp
}

func (t *StoreUser) Insert(ctx context.Context, tx *pg.Tx) (bool, error) {

	insert, err := tx.Model(t).Context(ctx).
		OnConflict("(store_id, user_id) DO NOTHING").
		Insert()

	if err != nil {
		return false, err
	}

	return insert.RowsAffected() > 0, err
}

type ListStoreUsersParams struct {
	StoreId string
	UserIds []string
	Phone   string
	Page    int64
	Size    int64
}

func (t *StoreUser) List(ctx context.Context, params ListStoreUsersParams) (StoreUsers, int64, error) {

	var tmp StoreUsers
	q := comp.SDK().Postgres().Model(&tmp).Context(ctx).
		Join("left join t_users u on store_user.user_id = u.id")

	if params.StoreId != "" {
		q = q.Where("store_user.store_id = ?", params.StoreId)
	}

	if params.Phone != "" {
		q = q.Where("u.phone = ?", params.Phone)
	}

	if len(params.UserIds) > 0 {
		q = q.WhereIn("store_user.user_id in (?)", params.UserIds)
	}

	if params.Size > 0 {
		page := params.Page
		if page <= 0 {
			page = 1
		}

		q = q.Limit(int(params.Size)).Offset(int((page - 1) * params.Size))
	}
	count, err := q.OrderExpr("store_user.created_at desc").SelectAndCount()

	if err != nil {
		return nil, 0, err
	}

	return tmp, int64(count), nil
}

func (t *StoreUser) Update(ctx context.Context, tx *pg.Tx, storeId, userId, field, value string) (bool, error) {

	update, err := tx.Model((*StoreUser)(nil)).Context(ctx).
		Where("store_id = ?", storeId).
		Where("user_id = ?", userId).
		Set("extra =  jsonb_set(extra, ?, ?)", fmt.Sprintf("{%s}", field), fmt.Sprintf("\"%s\"", value)).
		Update()

	if err != nil {
		return false, err
	}

	return update.RowsAffected() > 0, nil
}
