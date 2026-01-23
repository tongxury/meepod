package db

import (
	"context"
	"gitee.com/meepo/backend/kit/components/comp"
	mapset "github.com/deckarep/golang-set/v2"
	"time"
)

type ProxyUser struct {
	tableName   struct{} `pg:"t_proxy_users"`
	Id          string
	ProxyId     string
	ProxyUserId string
	StoreId     string
	UserId      string
	CreatedAt   time.Time
	Extra       string
}

type ProxyUsers []*ProxyUser

func (ts ProxyUsers) Ids() []string {

	tmp1 := mapset.NewSet[string]()
	for _, t := range ts {
		tmp1.Add(t.UserId)
	}

	return tmp1.ToSlice()
}

func (t *ProxyUser) CreateNX(ctx context.Context) (*ProxyUser, error) {

	_, err := comp.SDK().Postgres().Model(t).Context(ctx).
		OnConflict("(proxy_id, user_id) do nothing").
		Insert()

	if err != nil {
		return nil, err
	}

	return t, nil
}

type ListProxyUsersParams struct {
	StoreId     string
	ProxyId     string
	ProxyUserId string
	Page, Size  int64
}

func (t *ProxyUser) List(ctx context.Context, params ListProxyUsersParams) (ProxyUsers, int, error) {

	var tmp ProxyUsers
	q := comp.SDK().Postgres().Model(&tmp).Context(ctx)

	if params.StoreId != "" {
		q = q.Where("store_id = ?", params.StoreId)
	}

	if params.ProxyId != "" {
		q = q.Where("proxy_id = ?", params.ProxyId)
	}

	if params.ProxyUserId != "" {
		q = q.Where("proxy_user_id = ?", params.ProxyUserId)
	}

	if params.Size > 0 {
		page := params.Page
		if page <= 0 {
			page = 1
		}

		q = q.Limit(int(params.Size)).Offset(int((page - 1) * params.Size))
	}
	q = q.OrderExpr("created_at desc")

	count, err := q.SelectAndCount()
	if err != nil {
		return nil, 0, err
	}

	return tmp, count, nil
}
