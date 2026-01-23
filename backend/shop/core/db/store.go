package db

import (
	"context"
	"fmt"
	"time"

	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/kit/services/util/pgd"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-pg/pg/v10"
)

type Store struct {
	tableName struct{} `pg:"t_stores"`
	Id        string
	Name      string
	OwnerId   string
	CreatedAt time.Time
	Status    string
	Extra     StoreExtra
	Member    Member
	Settings  Settings
	//Icon      string
}

type Settings struct {
	ItemIds []string `json:"item_ids"`
	Notice  string   `json:"notice"`
}

type Member struct {
	Level string `json:"level"`
	Until int64  `json:"until"`
}

type StoreExtra struct {
	Icon        string `json:"icon"`
	WechatImage string `json:"wechat_image"`

	StoreFront      string `json:"store_front"`
	StoreInSide     string `json:"store_in_side"`
	Loc             string `json:"loc"`
	Address         string `json:"address"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	IdCardNo        string `json:"id_card_no"`
	IdCardFrom      string `json:"id_card_from"`
	IdCardFront     string `json:"id_card_front"`
	IdCardBack      string `json:"id_card_back"`
	IdCardHandled   string `json:"id_card_handled"`
	IdCardTo        string `json:"id_card_to"`
	SalesCard       string `json:"sales_card"`
	BankAccountName string `json:"bank_account_name"`
	BankCardFront   string `json:"bank_card_front"`
	BankCardBack    string `json:"bank_card_back"`
	BankName        string `json:"bank_name"`
	BankPhone       string `json:"bank_phone"`
	BankAccount     string `json:"bank_account"`
	BankBranch      string `json:"bank_branch"`
}

//
//func (t *Store) Icon() string {
//
//	icon, err := FindInExtra(t.Extra, "icon")
//	if err != nil {
//		return ""
//	}
//
//	return icon
//}

type Stores []*Store

func (ts Stores) Ids() ([]string, []string) {

	tmp1 := mapset.NewSet[string]()
	tmp2 := mapset.NewSet[string]()

	for _, t := range ts {
		tmp1.Add(t.Id)
		tmp2.Add(t.OwnerId)
	}

	return tmp1.ToSlice(), tmp2.ToSlice()

}

func (ts Stores) AsMap() map[string]*Store {
	rsp := make(map[string]*Store, len(ts))

	for _, t := range ts {
		rsp[t.Id] = t
	}

	return rsp
}

func (t *Store) FindByOwnerIdAndStatus(ctx context.Context, ownerId, status string) (*Store, error) {
	var tmp Stores
	err := comp.SDK().Postgres().Model(&tmp).Context(ctx).
		Where("owner_id = ?", ownerId).
		Where("status = ?", status).
		Select()

	if err != nil {
		return nil, err
	}

	if len(tmp) == 0 {
		return nil, nil
	}

	return tmp[0], nil
}

func (t *Store) FindByOwnerId(ctx context.Context, ownerId string) (*Store, error) {
	var tmp Stores
	err := comp.SDK().Postgres().Model(&tmp).Context(ctx).
		Where("owner_id = ?", ownerId).
		Select()

	if err != nil {
		return nil, err
	}

	if len(tmp) == 0 {
		return nil, nil
	}

	return tmp[0], nil
}

func (t *Store) FindByIdAndStatus(ctx context.Context, id, status string) (*Store, error) {

	var tmp Stores
	err := comp.SDK().Postgres().Model(&tmp).Context(ctx).
		Where("id = ?", id).
		Where("status = ?", status).
		Select()

	if err != nil {
		return nil, err
	}

	if len(tmp) == 0 {
		return nil, nil
	}

	return tmp[0], nil
}

func (t *Store) RequireById(ctx context.Context, id string) (*Store, error) {

	store, err := t.FindById(ctx, id)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	if store == nil {
		return nil, fmt.Errorf("no store found by id: %s", id)
	}

	return store, nil
}

func (t *Store) FindById(ctx context.Context, id string) (*Store, error) {

	var tmp Stores
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

func (t *Store) UpdateMember(ctx context.Context, tx *pg.Tx, id string, field string, value any) (bool, error) {

	q := tx.Model(t).Context(ctx).Where("id = ?", id)
	q = pgd.SetJSONField(q, "member", field, value)

	u, err := q.Update()

	if err != nil {
		return false, err
	}

	return u.RowsAffected() > 0, nil
}

func (t *Store) UpdateSettings(ctx context.Context, tx *pg.Tx, id string, field string, value any) (bool, error) {

	q := tx.Model(t).Context(ctx).Where("id = ?", id)
	q = pgd.SetJSONField(q, "settings", field, value)

	u, err := q.Update()

	if err != nil {
		return false, err
	}

	return u.RowsAffected() > 0, nil
}

func (t *Store) Update(ctx context.Context, tx *pg.Tx, id string, field, value any) (bool, error) {

	q := tx.Model(t).Context(ctx).Where("id = ?", id)

	switch field {
	case "name":
		q = q.Set("name = ?", value)
	case "extra":
		q = q.Set("extra = ?", value)
	default:
		switch t := value.(type) {
		case string:
			value = fmt.Sprintf("\"%v\"", t)
		}

		q = q.Set("extra = jsonb_set(extra, ?, ?)", fmt.Sprintf("{%s}", field), value)
	}

	u, err := q.Update()

	if err != nil {
		return false, err
	}

	return u.RowsAffected() > 0, nil
}

func (t *Store) Insert(ctx context.Context, tx *pg.Tx) (bool, error) {

	insert, err := tx.Model(t).Context(ctx).
		OnConflict("(id) DO NOTHING").
		Insert()

	if err != nil {
		return false, err
	}

	return insert.RowsAffected() > 0, err
}

func (t *Store) ListByIds(ctx context.Context, ids []string) (Stores, error) {

	if len(ids) == 0 {
		return nil, nil
	}

	var rsp Stores

	q := comp.SDK().Postgres().Model(&rsp).Context(ctx).
		WhereIn("id in (?)", ids)

	err := q.Select()

	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func (t *Store) UpdateStatus(ctx context.Context, tx *pg.Tx, ids []string, status string) (bool, error) {

	update, err := tx.Model((*Store)(nil)).Context(ctx).
		WhereIn("id in (?)", ids).
		Set("status = ?", status).
		Update()

	if err != nil {
		return false, err
	}

	return update.RowsAffected() == len(ids), nil
}

func (t *Store) UpdateExtra(ctx context.Context, tx *pg.Tx, id, field, value string) (bool, error) {

	update, err := tx.Model((*Store)(nil)).Context(ctx).
		Where("id = ?", id).
		Set("extra =  jsonb_set(extra, ?, ?)", fmt.Sprintf("{%s}", field), fmt.Sprintf("\"%s\"", value)).
		Update()

	if err != nil {
		return false, err
	}

	return update.RowsAffected() > 0, nil
}

type ListStoresParams struct {
	Ids      []string
	Keyword  string
	Id       string
	NameLike string
	OwnerId  string
	Phone    string
	Page     int64
	Size     int64
}

func (t *Store) List(ctx context.Context, params ListStoresParams) (Stores, int64, error) {

	var tmp Stores
	q := comp.SDK().Postgres().Model(&tmp).Context(ctx).
		Join("left join t_users as u on store.owner_id = u.id")

	if params.Keyword != "" {
		q = q.Where("store.id like ? or store.name like ?", "%"+params.Keyword+"%", "%"+params.Keyword+"%")
	}

	if params.Id != "" {
		q = q.Where("store.id = ?", params.Id)
	}

	if len(params.Ids) != 0 {
		q = q.WhereIn("store.id in (?)", params.Ids)
	}

	if params.NameLike != "" {
		q = q.Where("store.name like ?", "%"+params.NameLike+"%")
	}

	if params.OwnerId != "" {
		q = q.Where("store.owner_id = ?", params.OwnerId)
	}

	if params.Phone != "" {
		q = q.Where("u.phone = ?", params.Phone)
	}

	if params.Size > 0 {
		page := params.Page
		if page <= 0 {
			page = 1
		}

		q = q.Limit(int(params.Size)).Offset(int((page - 1) * params.Size))
	}
	count, err := q.OrderExpr("store.created_at desc").SelectAndCount()

	if err != nil {
		return nil, 0, err
	}

	return tmp, int64(count), nil
}
