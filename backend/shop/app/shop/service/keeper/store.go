package keeperservice

import (
	"context"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/app/shop/service"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/types"
	"github.com/go-pg/pg/v10"
)

type StoreService struct {
	service.StoreService
}

func (t *StoreService) ListStores(ctx context.Context, keeperId, id, keyword string) (types.Stores, error) {

	dbStores, _, err := new(db.Store).List(ctx, db.ListStoresParams{
		Id:      id,
		Keyword: keyword, Page: 1, Size: 20,
	})
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	var filterStores db.Stores
	for _, store := range dbStores {
		if store.OwnerId == keeperId {
			continue
		}
		filterStores = append(filterStores, store)
	}

	stores, err := t.Assemble(ctx, filterStores)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return stores, nil

}

func (t *StoreService) UpdateNotice(ctx context.Context, keeperId, storeId string, notice string) error {

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		_, err := new(db.Store).UpdateSettings(ctx, tx, storeId, "notice", notice)
		if err != nil {
			return err
		}

		return nil
	})
	return xerror.Wrap(err)
}

func (t *StoreService) UpdateItemSettings(ctx context.Context, keeperId, storeId string, itemIds []string) error {

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		_, err := new(db.Store).UpdateSettings(ctx, tx, storeId, "item_ids", itemIds)
		if err != nil {
			return err
		}

		return nil
	})
	return xerror.Wrap(err)
}

//func (t *StoreService) GetItemSettings(ctx context.Context, keeperId, storeId string) (model.Items, []string, error) {
//
//	items, err := new(ItemService).ListMetaByIds(ctx, nil, true)
//	if err != nil {
//		return nil, nil, xerror.Wrap(err)
//	}
//
//	store, err := new(db.Store).RequireById(ctx, storeId)
//	if err != nil {
//		return nil, nil, xerror.Wrap(err)
//	}
//
//	itemIds := store.Settings.ItemIds
//
//	return items, itemIds, nil
//}

func (t *StoreService) Update(ctx context.Context, id, field string, value string) error {

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		_, err := new(db.Store).Update(ctx, tx, id, field, value)
		return err
	})

	return xerror.Wrap(err)
}
