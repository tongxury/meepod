package db

import (
	"context"
	"fmt"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/core/enum"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-pg/pg/v10"
	"time"
)

type OrderGroupShare struct {
	tableName struct{} `pg:"t_order_group_shares"`
	Id        string
	Volume    int64
	Amount    float64
	StoreId   string
	GroupId   string
	UserId    string
	CreatedAt time.Time
	Status    string
	Extra     string
}

type OrderGroupShares []*OrderGroupShare

func (ts OrderGroupShares) GroupByUserId() map[string]int64 {

	rsp := map[string]int64{}
	for _, x := range ts {
		rsp[x.UserId] += x.Volume
	}

	return rsp
}

func (ts OrderGroupShares) Ids() ([]string, []string, []string) {
	tmp1 := mapset.NewSet[string]()
	tmp2 := mapset.NewSet[string]()
	tmp3 := mapset.NewSet[string]()
	for _, t := range ts {
		tmp1.Add(t.Id)
		tmp2.Add(t.GroupId)
		tmp3.Add(t.UserId)
	}

	return tmp1.ToSlice(), tmp2.ToSlice(), tmp3.ToSlice()
}

func (t *OrderGroupShare) ListJoinedGroupIds(ctx context.Context, storeId, userId string) ([]string, error) {

	var result []struct {
		GroupId string
	}

	err := comp.SDK().Postgres().WithContext(ctx).Model(t).
		ColumnExpr("distinct(group_id) as group_id").
		Where("user_id = ?", userId).
		Where("store_id = ?", storeId).
		Select(&result)
	if err != nil {
		return nil, err
	}

	var ids []string
	for _, x := range result {
		ids = append(ids, x.GroupId)
	}

	return ids, nil
}

func (t *OrderGroupShare) FindJoinerCounts(ctx context.Context, groupIds []string) (map[string]int64, error) {

	var counts []struct {
		Count   int64
		GroupId string
	}

	err := comp.SDK().Postgres().WithContext(ctx).Model(t).
		ColumnExpr("count(1) as count").
		ColumnExpr("group_id").
		GroupExpr("group_id").
		WhereIn("group_id in (?)", groupIds).
		Select(&counts)
	if err != nil {
		return nil, err
	}

	rsp := make(map[string]int64, len(counts))
	for _, x := range counts {
		rsp[x.GroupId] = x.Count
	}

	return rsp, nil
}

func (t *OrderGroupShare) FindByGroupIdAndStatus(ctx context.Context, groupId string, status []string) (OrderGroupShares, error) {

	var rsp OrderGroupShares
	err := comp.SDK().Postgres().WithContext(ctx).Model(&rsp).
		Where("group_id = ?", groupId).
		WhereIn("status in (?)", status).
		Select()

	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func (t *OrderGroupShare) UpdateToPayed(ctx context.Context, tx *pg.Tx, id string) (bool, error) {

	order := OrderGroupShare{Id: id}

	q := tx.Model(&order).Context(ctx).
		Set("status = ?", enum.OrderGroupShareStatus_Payed.Value)

	updated, err := q.Where("id = ?", id).
		WhereIn("status in (?)", enum.PayableOrderGroupShareStatus).
		Update()

	if err != nil {
		return false, xerror.Wrap(err)
	}

	return updated.RowsAffected() > 0, nil

}

func (t *OrderGroupShare) RequireById(ctx context.Context, id string) (*OrderGroupShare, error) {

	var tmp OrderGroupShares
	err := comp.SDK().Postgres().WithContext(ctx).Model(&tmp).
		Where("id = ?", id).
		Select()

	if err != nil {
		return nil, err
	}

	if len(tmp) == 0 {
		return nil, fmt.Errorf("no order group share found by: %s", id)
	}

	return tmp[0], nil
}

func (t *OrderGroupShare) RequireByIdAndCreatorId(ctx context.Context, id, userId string) (*OrderGroupShare, error) {

	var tmp OrderGroupShares
	err := comp.SDK().Postgres().WithContext(ctx).Model(&tmp).
		Where("id = ?", id).
		Where("user_id = ?", userId).
		Select()

	if err != nil {
		return nil, err
	}

	if len(tmp) == 0 {
		return nil, fmt.Errorf("no order found by: %s, %s", id, userId)
	}

	return tmp[0], nil
}

func (t *OrderGroupShare) Rollback(ctx context.Context, id string) error {
	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		var tmp OrderGroupShares

		err := tx.Model(&tmp).Where("id = ?", id).Select()
		if err != nil {
			return err
		}

		// 更新剩余分数
		if len(tmp) > 0 {

			_, err := tx.Model(t).Where("id = ?", id).Delete()
			if err != nil {
				return err
			}

			_, err = tx.Model((*OrderGroup)(nil)).
				Set("volume_ordered = volume_ordered - ?", tmp[0].Volume).
				Set("shares =  jsonb_delete(shares, ?)", fmt.Sprintf("%s", tmp[0].Id)).
				Where("id = ?", tmp[0].GroupId).Update()

			if err != nil {
				return err
			}

			// 更新group 状态
			_, err = tx.Model((*OrderGroup)(nil)).
				Where("id = ?", tmp[0].GroupId).
				Set("status = case when volume = volume_ordered then ? else ? end",
					enum.OrderGroupStatus_Payed.Value,
					enum.OrderGroupStatus_Submitted.Value).
				Update()
			if err != nil {
				return err
			}

		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (t *OrderGroupShare) Insert(ctx context.Context) error {

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		// 插入合买订单
		insertResult, err := tx.Model(t).Context(ctx).Insert()
		if err != nil {
			return err
		}

		// 更新剩余分数
		if insertResult.RowsAffected() > 0 {
			_, err = tx.Model((*OrderGroup)(nil)).
				Set("volume_ordered = volume_ordered + ?", t.Volume).
				Set("shares =  jsonb_set(shares, ?, ?)", fmt.Sprintf("{%s}", t.Id), fmt.Sprintf("\"%s\"", t.UserId)).
				Where("id = ?", t.GroupId).Update()

			if err != nil {
				return err
			}
		}

		// 更新group 状态
		_, err = tx.Model((*OrderGroup)(nil)).
			Where("id = ?", t.GroupId).
			Set("status = case when volume = volume_ordered then ? else ? end",
				enum.OrderGroupStatus_Payed.Value,
				enum.OrderGroupStatus_Submitted.Value).
			Update()
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
