package service

import (
	"context"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/types"
	"github.com/go-pg/pg/v10"
)

type FeedbackService struct {
}

func (t *FeedbackService) Resolve(ctx context.Context, storeId, id string) error {
	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {
		_, err := new(db.Feedback).UpdateStatus(ctx, tx, id, storeId, enum.FeedbackStatus_Resolved.Value)
		return err
	})
	return xerror.Wrap(err)
}

func (t *FeedbackService) Add(ctx context.Context, storeId, userId, text string) error {

	fk := db.Feedback{
		StoreId: storeId,
		UserId:  userId,
		Text:    text,
		Status:  enum.FeedbackStatus_Submitted.Value,
	}

	_, err := fk.Add(ctx)
	if err != nil {
		return xerror.Wrap(err)
	}

	return nil
}

func (t *FeedbackService) List(ctx context.Context, storeId, userId string, page, size int64) (types.Feedbacks, int64, error) {

	dbItems, total, err := new(db.Feedback).List(ctx, db.ListFeedbackParams{
		StoreId: storeId, UserId: userId,
		Page: page, Size: size,
	})
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	users, err := t.Assemble(ctx, dbItems)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	return users, total, nil

}

func (t *FeedbackService) FindById(ctx context.Context, storeId, id string) (*types.Feedback, error) {

	dbItems, _, err := new(db.Feedback).List(ctx, db.ListFeedbackParams{Id: id, StoreId: storeId})
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	items, err := t.Assemble(ctx, dbItems)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return items[0], nil
}

//func (t *FeedbackService) FindByUserId(ctx context.Context, storeId, userId string) (*types.Proxy, error) {
//
//	dbProxy, err := new(db.Proxy).FindByUserId(ctx, storeId, userId)
//	if err != nil {
//		return nil, xerror.Wrap(err)
//	}
//
//	if dbProxy == nil {
//		return nil, nil
//	}
//
//	proxies, err := t.Assemble(ctx, db.Proxies{dbProxy})
//	if err != nil {
//		return nil, xerror.Wrap(err)
//	}
//
//	return proxies[0], nil
//}

func (t *FeedbackService) Assemble(ctx context.Context, items db.Feedbacks) (types.Feedbacks, error) {

	_, userIds, storeIds := items.Ids()

	dbStores, _, err := new(db.Store).List(ctx, db.ListStoresParams{
		Ids: storeIds,
	})
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	dbStoresMap := dbStores.AsMap()
	_, ownerIds := dbStores.Ids()
	userIds = append(userIds, ownerIds...)

	dbUsers, err := new(db.User).ListByIds(ctx, userIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	usersMap := dbUsers.AsMap()

	var tmp types.Feedbacks
	for _, x := range items {

		store := dbStoresMap[x.StoreId]
		storeOwner := usersMap[store.OwnerId]

		y := types.FromDbFeedback(x, usersMap[x.UserId], store, storeOwner)

		tmp = append(tmp, y)
	}

	return tmp, nil

}
