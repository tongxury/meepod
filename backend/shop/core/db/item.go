package db

import (
	"context"
	"gitee.com/meepo/backend/kit/components/comp"
	mapset "github.com/deckarep/golang-set/v2"
)

type Item struct {
	tableName struct{} `pg:"t_items"`
	Id        string
	Name      string
	Icon      string
	Status    string
}

type Items []*Item

func (ts Items) Ids() []string {
	tmp := mapset.NewSet[string]()

	for _, t := range ts {
		tmp.Add(t.Id)
	}

	return tmp.ToSlice()
}

func (ts Items) AsMap() map[string]*Item {
	rsp := make(map[string]*Item, len(ts))

	for _, t := range ts {
		rsp[t.Id] = t
	}

	return rsp
}

func (ts Items) Insert(ctx context.Context) error {

	_, err := comp.SDK().Postgres().Model(&ts).Context(ctx).Insert()
	if err != nil {
		return err
	}

	return nil
}

func (t *Item) ListByIds(ctx context.Context, ids []string, allIfEmptyIds bool) (Items, error) {

	var rsp Items

	q := comp.SDK().Postgres().Model(&rsp).Context(ctx)

	if len(ids) > 0 {
		q = q.WhereIn("id in (?)", ids)
	} else {
		if !allIfEmptyIds {
			return nil, nil
		}
	}

	err := q.OrderExpr("sort").Select()

	if err != nil {
		return nil, err
	}

	return rsp, nil
}
