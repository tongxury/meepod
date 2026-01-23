package service

import (
	"context"
	"fmt"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/kit/services/util/gind/errorx"
	paymentdb "gitee.com/meepo/backend/shop/app/payment/db"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/types"
	"github.com/go-pg/pg/v10"
	"github.com/lithammer/shortuuid/v4"
	redisV9 "github.com/redis/go-redis/v9"
	"time"
)

type StoreService struct {
}

func (t *StoreService) ListStores(ctx context.Context, id, name, phone string, page, size int64) (types.Stores, int64, error) {

	dbStores, total, err := new(db.Store).List(ctx, db.ListStoresParams{
		Id: id, NameLike: name, Phone: phone, Page: page, Size: size,
	})
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	stores, err := t.Assemble(ctx, dbStores)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	return stores, total, nil

}

func (t *StoreService) UpdateStatus(ctx context.Context, ids []string, status string) error {

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		_, err := new(db.Store).UpdateStatus(ctx, tx, ids, status)
		if err != nil {
			return err
		}

		return nil
	})

	return xerror.Wrap(err)
}

func (t *StoreService) Confirm(ctx context.Context, ids []string) error {

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		_, err := new(db.Store).UpdateStatus(ctx, tx, ids, enum.StoreStatus_Confirmed.Value)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return xerror.Wrap(err)
	}

	for _, id := range ids {
		t.sendEvent(ctx, "store.created.event", id)
	}
	return nil
}

func (t *StoreService) UpdateMember(ctx context.Context, id, level string, until int64) error {
	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		for k, v := range map[string]any{
			"level": level,
			"until": until,
		} {
			_, err := new(db.Store).UpdateMember(ctx, tx, id, k, v)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return xerror.Wrap(err)
	}

	return nil
}

func (t *StoreService) UpdateStoreExtra(ctx context.Context, id string, params types.StoreParams) error {

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		//for k, v := range map[string]string{
		//	"name":          params.StoreName,
		//	"id_card_front": params.IdCardFront,
		//	"id_card_back":  params.IdCardBack,
		//	"store_front":   params.StoreFront,
		//	"loc":           params.Loc,
		//	"address":       params.Address,
		//} {
		//	_, err := new(db.Store).Update(ctx, tx, id, k, v)
		//	if err != nil {
		//		return err
		//	}
		//}
		_, err := new(db.Store).Update(ctx, tx, id, "extra", t.toExtra(params))
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return xerror.Wrap(err)
	}

	return nil
}

func (t *StoreService) toExtra(params types.StoreParams) db.StoreExtra {
	return db.StoreExtra{
		Icon: enum.DefaultStoreIcon,

		StoreFront:      params.StoreFront,
		StoreInSide:     params.StoreInSide,
		Loc:             params.Loc,
		Address:         params.Address,
		Username:        params.Username,
		Email:           params.Email,
		IdCardNo:        params.IdCardNo,
		IdCardFrom:      params.IdCardFrom[:10],
		IdCardTo:        params.IdCardTo[:10],
		IdCardFront:     params.IdCardFront,
		IdCardBack:      params.IdCardBack,
		IdCardHandled:   params.IdCardHandled,
		SalesCard:       params.SalesCard,
		BankAccountName: params.BankAccountName,
		BankCardFront:   params.BankCardFront,
		BankCardBack:    params.BankCardBack,
		BankName:        params.BankName,
		BankPhone:       params.BankPhone,
		BankAccount:     params.BankAccount,
		BankBranch:      params.BankBranch,
	}
}

func (t *StoreService) AddStore(ctx context.Context, params types.StoreParams) error {

	owner := db.User{
		Phone: params.Phone,
	}

	user, err := owner.CreateNX(ctx)
	if err != nil {
		return xerror.Wrap(err)
	}

	oldStore, err := new(db.Store).FindByOwnerId(ctx, user.Id)
	if err != nil {
		return xerror.Wrap(err)
	}

	if oldStore != nil {
		return errorx.UserMessage("重复创建")
	}

	items, err := new(db.Item).ListByIds(ctx, nil, true)
	if err != nil {
		return xerror.Wrap(err)
	}

	store := db.Store{
		Id:      shortuuid.New()[:8],
		Name:    params.StoreName,
		OwnerId: user.Id,
		Status:  enum.StoreStatus_Pending.Value,
		Member: db.Member{
			Level: enum.MemberLevel_Normal.Value,
			Until: time.Now().AddDate(1, 0, 0).Unix(),
		},
		Settings: db.Settings{
			ItemIds: items.Ids(),
		},
		Extra: t.toExtra(params),
	}

	err = comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		for {
			upsert, err := store.Insert(ctx, tx)
			if err != nil {
				return err
			}

			if upsert {
				return nil
			}

			store.Id = shortuuid.New()[:8]
		}

	})

	if err != nil {
		return xerror.Wrap(err)
	}

	return nil

}

func (t *StoreService) sendEvent(ctx context.Context, event string, storeId string) {

	err := comp.SDK().Redis().XAdd(ctx, &redisV9.XAddArgs{
		Stream: event,
		Values: map[string]interface{}{
			"storeId": storeId,
		},
	}).Err()
	if err != nil {
		slf.WithError(err).Errorw("XAdd err")
	}
}

func (t *StoreService) RequireConfirmedStoreById(ctx context.Context, storeId string) (*types.Store, error) {

	store, err := t.FindConfirmedStoreById(ctx, storeId)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	if store == nil {
		return nil, fmt.Errorf("no confirmed store found by id: %s", storeId)
	}

	return store, nil
}

func (t *StoreService) FindConfirmedStoreById(ctx context.Context, storeId string) (*types.Store, error) {

	dbStore, err := new(db.Store).FindByIdAndStatus(ctx, storeId, enum.StoreStatus_Confirmed.Value)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	if dbStore == nil {
		return nil, nil
	}

	stores, err := t.Assemble(ctx, db.Stores{dbStore})
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return stores[0], nil
}

func (t *StoreService) FindStoreById(ctx context.Context, storeId string) (*types.Store, error) {

	dbStore, err := new(db.Store).FindById(ctx, storeId)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	if dbStore == nil {
		return nil, nil
	}

	stores, err := t.Assemble(ctx, db.Stores{dbStore})
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return stores[0], nil
}

func (t *StoreService) Assemble(ctx context.Context, stores db.Stores) (types.Stores, error) {

	var extraMaps db.ExtraMaps

	storeIds, userIds := stores.Ids()

	dbItems, err := new(db.Item).ListByIds(ctx, nil, true)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	extraMaps.Items = dbItems.AsMap()

	users, err := new(db.User).ListByIds(ctx, userIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	usersMap := users.AsMap()

	wallets, _, err := new(paymentdb.StoreWallet).List(ctx, paymentdb.ListStoreWalletsParams{
		StoreIds: storeIds,
	})
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	walletsMap := wallets.AsMap()

	var rsp types.Stores

	for _, x := range stores {

		y := types.FromDbStore(x, usersMap[x.OwnerId], extraMaps)

		if val, found := walletsMap[x.Id]; found {
			y.Balance = val.Balance
		}

		rsp = append(rsp, y)
	}

	return rsp, nil
}
