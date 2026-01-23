package keeperservice

import (
	"context"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/kit/services/util/gind/errorx"
	paymentdb "gitee.com/meepo/backend/shop/app/payment/db"
	keeperservice "gitee.com/meepo/backend/shop/app/payment/service/keeper"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/types"
	"github.com/go-pg/pg/v10"
)

type CoStoreService struct {
}

func (t *CoStoreService) UpdateItems(ctx context.Context, keeperId, storeId, coStoreId string, items map[string]float64) error {

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		_, err := new(db.CoStore).UpdateItems(ctx, tx, storeId, coStoreId, items)
		if err != nil {
			return err
		}

		return err
	})

	return xerror.Wrap(err)
}

func (t *CoStoreService) Pause(ctx context.Context, keeperId, storeId, coStoreId string) error {

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		_, err := new(db.CoStore).UpdateStatus(ctx, tx, storeId, coStoreId,
			enum.CoStoreStatus_Paused.Value, enum.PausableCoStoreStatus)
		if err != nil {
			return err
		}

		return err
	})

	return xerror.Wrap(err)
}

func (t *CoStoreService) Resume(ctx context.Context, keeperId, storeId, coStoreId string) error {

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		_, err := new(db.CoStore).UpdateStatus(ctx, tx, storeId, coStoreId,
			enum.CoStoreStatus_Confirmed.Value, enum.ResumableCoStoreStatus)
		if err != nil {
			return err
		}

		return err
	})

	return xerror.Wrap(err)
}

func (t *CoStoreService) Recover(ctx context.Context, keeperId, storeId, coStoreId string) error {

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		_, err := new(db.CoStore).UpdateStatus(ctx, tx, storeId, coStoreId,
			enum.ProxyStatus_Confirmed.Value, enum.RecoverableCoStoreStatus)
		if err != nil {
			return err
		}

		return err
	})

	return xerror.Wrap(err)
}

func (t *CoStoreService) ApplyForEnd(ctx context.Context, keeperId, storeId, coStoreId string) error {

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		_, err := new(db.CoStore).UpdateStatus(ctx, tx, storeId, coStoreId,
			enum.CoStoreStatus_EndPending.Value, enum.EndApplyableCoStoreStatus)
		if err != nil {
			return err
		}

		return err
	})

	return xerror.Wrap(err)
}

func (t *CoStoreService) End(ctx context.Context, keeperId, storeId, coStoreId, imageProof string) error {

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		_, err := new(db.CoStore).UpdateStatus(ctx, tx, coStoreId, storeId,
			enum.CoStoreStatus_End.Value, enum.EndableCoStoreStatus)
		if err != nil {
			return err
		}

		//if updated {
		//	_, err := new(db.CoStore).UpdateExtra(ctx, tx, coStoreId, storeId, "end_proof", imageProof)
		//	if err != nil {
		//		return err
		//	}
		//}

		return err
	})

	if err != nil {
		return xerror.Wrap(err)
	}

	// todo event
	err = new(keeperservice.CoStorePaymentService).Return(ctx, coStoreId, storeId, imageProof)
	if err != nil {
		slf.WithError(err).Errorw("Return err")
	}

	_, err = new(db.CoStore).UpdateToSynced(ctx, coStoreId, storeId, "sync_return")
	if err != nil {
		slf.WithError(err).Errorw("UpdateToSynced err")
	}

	return xerror.Wrap(err)
}

func (t *CoStoreService) ListOutCoStores(ctx context.Context, enableOnly bool, itemId, keeperId, storeId string, page, size int64) (types.CoStores, int64, error) {

	params := db.ListCoStoresParams{
		StoreId: storeId,
		Page:    page, Size: size,
	}
	if enableOnly {
		params.MStatus = enum.EnableCoStoreStatus
	}

	dbItems, total, err := new(db.CoStore).List(ctx, params)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	if itemId != "" {
		var filterDbItems db.CoStores

		for _, x := range dbItems {
			if _, found := x.Items[itemId]; found {
				filterDbItems = append(filterDbItems, x)
			}
		}
		dbItems = filterDbItems
	}

	items, err := t.Assemble(ctx, dbItems)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	for _, item := range items {
		item.TopUpable = false
	}

	return items, total, nil

}

func (t *CoStoreService) ListInCoStores(ctx context.Context, enableOnly bool, keeperId, storeId string, page, size int64) (types.CoStores, int64, error) {

	params := db.ListCoStoresParams{
		CoStoreId: storeId,
		Page:      page, Size: size,
	}
	if enableOnly {
		params.MStatus = enum.EnableCoStoreStatus
	}

	dbItems, total, err := new(db.CoStore).List(ctx, params)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	items, err := t.Assemble(ctx, dbItems)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	for _, item := range items {
		item.Updatable = false
		item.Pausable = false
		item.Resumable = false
		item.EndApplyable = false
		item.Recoverable = false
	}

	return items, total, nil

}

func (t *CoStoreService) Assemble(ctx context.Context, xs db.CoStores) (types.CoStores, error) {

	_, storeIds, coStoreIds := xs.Ids()
	allStoreIds := append(storeIds, coStoreIds...)

	dbStores, err := new(db.Store).ListByIds(ctx, allStoreIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	dbStoresMap := dbStores.AsMap()

	_, ownerIds := dbStores.Ids()

	dbUsers, err := new(db.User).ListByIds(ctx, ownerIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	usersMap := dbUsers.AsMap()

	var extraMaps db.ExtraMaps
	dbItems, err := new(db.Item).ListByIds(ctx, nil, true)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	extraMaps.Items = dbItems.AsMap()

	wallets, _, err := new(paymentdb.CoStoreWallet).List(ctx, paymentdb.ListCoStoreWalletsParams{
		StoreIds: storeIds, CoStoreIds: coStoreIds,
	})
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	walletsMap := wallets.AsMap()

	var tmp types.CoStores

	for _, x := range xs {

		store := dbStoresMap[x.StoreId]
		storeOwner := usersMap[store.OwnerId]
		coStore := dbStoresMap[x.CoStoreId]
		coStoreOwner := usersMap[coStore.OwnerId]

		y := types.FromDbCoStore(x, store, storeOwner, coStore, coStoreOwner, extraMaps)

		if val, found := walletsMap[x.StoreId+"-"+x.CoStoreId]; found {
			y.Balance = val.Balance
		}

		tmp = append(tmp, y)
	}

	return tmp, nil

}

func (t *CoStoreService) AddCoStore(ctx context.Context, keeperId, storeId, coStoreId string, itemIds []string) error {

	if coStoreId == storeId {
		return errorx.UserMessage("不能添加自己为合作店铺")
	}

	store, err := new(StoreService).RequireConfirmedStoreById(ctx, storeId)
	if err != nil {
		return xerror.Wrap(err)
	}

	oldCoStores, count, err := new(db.CoStore).List(ctx, db.ListCoStoresParams{
		StoreId: storeId,
	})
	if err != nil {
		return xerror.Wrap(err)
	}

	if _, _, coStoreIds := oldCoStores.Ids(); helper.InSlice(coStoreId, coStoreIds) {
		return errorx.UserMessage("重复创建")
	}

	if store.MemberLevel.Value == enum.MemberLevel_Normal.Value {
		if count >= 1 {
			return errorx.UserMessage("请升级会员等级")
		}
	}

	items := make(map[string]float64, len(itemIds))
	for _, id := range itemIds {
		items[id] = 1
	}

	dbCoStore := db.CoStore{
		StoreId:   storeId,
		CoStoreId: coStoreId,
		Items:     items,
		Status:    enum.ProxyStatus_Confirmed.Value,
	}

	_, err = dbCoStore.CreateNX(ctx)
	if err != nil {
		return xerror.Wrap(err)
	}
	return nil
}

func (t *CoStoreService) FindByCoStoreId(ctx context.Context, storeId, coStoreId string) (*types.CoStore, error) {

	dbItems, _, err := new(db.CoStore).List(ctx, db.ListCoStoresParams{StoreId: storeId, CoStoreId: coStoreId})
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	items, err := t.Assemble(ctx, dbItems)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return items[0], nil
}
