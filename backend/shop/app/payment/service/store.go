package service

import (
	"context"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/kit/components/sdk/helper/timed"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/kit/components/third/xinsh"
	"gitee.com/meepo/backend/shop/app/payment/db"
	coredb "gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/types"
	"strings"
	"time"
)

type StoreService struct {
}

func (t *StoreService) AddStore(ctx context.Context, appId, storeId string) (*types.PaymentStore, error) {

	dbStore := &db.Store{
		StoreId: storeId,
		//AppId:   appId,
		Alipay: &db.Alipay{
			AppId: appId,
		},
		Xinsh: &db.Xinsh{},
	}

	_, err := dbStore.Insert(ctx)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	stores, err := t.Assemble(ctx, db.Stores{dbStore})
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return stores[0], nil
}

func (t *StoreService) ApplyXinsh(ctx context.Context, storeId string) error {

	dbStore, err := new(coredb.Store).RequireById(ctx, storeId)
	if err != nil {
		return xerror.Wrap(err)
	}

	pStore, err := new(db.Store).RequireByStoreId(ctx, storeId)
	if err != nil {
		return xerror.Wrap(err)
	}

	if helper.InSlice(pStore.Xinsh.Status, []string{
		enum.PaymentStoreStatus_Pending.Value,
		enum.PaymentStoreStatus_Confirmed.Value,
	}) {
		return nil
	}

	isUpdate := pStore.Xinsh.Status == enum.PaymentStoreStatus_Rejected.Value

	user, err := new(coredb.User).FindById(ctx, dbStore.OwnerId)
	if err != nil {
		return err
	}

	locParts := strings.Split(dbStore.Extra.Loc, "-")
	p, c, d := locParts[0], locParts[1], locParts[2]

	//bankLocParts := strings.Split(dbStore.Extra.BankLoc, "-")
	//bp, bc := bankLocParts[0], bankLocParts[1]

	bankBranchParts := strings.Split(dbStore.Extra.BankBranch, "-")
	bankId, branchNo := bankBranchParts[0], bankBranchParts[1]

	bank, err := new(coredb.Bank).RequireByBranchNo(ctx, branchNo)
	if err != nil {
		return err
	}

	resp, err := comp.SDK().Xinsh().SaveMerchant(ctx, xinsh.MerchantParams{
		RequestId:       conv.String(time.Now().UnixMilli()),
		MerchantNm:      dbStore.Extra.Username,
		Province:        p,
		City:            c,
		District:        d,
		Username:        dbStore.Extra.Username,
		Address:         dbStore.Extra.Address,
		Email:           dbStore.Extra.Email,
		Phone:           user.Phone,
		IdCardNo:        dbStore.Extra.IdCardNo,
		IdCardFrom:      dbStore.Extra.IdCardFrom,
		IdCardTo:        dbStore.Extra.IdCardTo,
		IdCardFront:     dbStore.Extra.IdCardFront,
		IdCardBack:      dbStore.Extra.IdCardBack,
		IdCardHandled:   dbStore.Extra.IdCardHandled,
		StoreFront:      dbStore.Extra.StoreFront,
		StoreInSide:     dbStore.Extra.StoreInSide,
		BankAccount:     dbStore.Extra.BankAccount,
		BankAccountName: dbStore.Extra.BankAccountName,
		AccountType:     "2",
		AccountNature:   "1",
		CnapsCode:       branchNo,
		BankId:          bankId,
		BankCardFront:   dbStore.Extra.BankCardFront,
		BankName:        dbStore.Extra.BankName,
		BankPhone:       dbStore.Extra.BankPhone,
		BankProvince:    bank.ProvinceId,
		BankCity:        bank.CityId,
	}, isUpdate)

	if err != nil {
		return xerror.Wrap(err)
	}

	err = new(db.Store).UpdateXinsh(ctx, storeId, db.Xinsh{
		MerchantNo: resp.MerchantNo,
		RequestId:  resp.RequestId,
	})
	if err != nil {
		slf.WithError(err).Errorw("[MAIN] UpdateXinSh err", slf.Reflect("resp", resp))
	}

	return nil
}

func (t *StoreService) AddStoreV2(ctx context.Context, storeId string) (*types.PaymentStore, error) {

	dbStore := &db.Store{
		StoreId: storeId,
	}

	_, err := dbStore.Insert(ctx)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	stores, err := t.Assemble(ctx, db.Stores{dbStore})
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return stores[0], nil
}

func (t *StoreService) ListStores(ctx context.Context, storeId string, page, size int64) (types.PaymentStores, int64, error) {

	dbStores, total, err := new(db.Store).List(ctx, db.ListStoresParams{
		StoreId: storeId,
		Page:    page, Size: size,
	})
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	payments, err := t.Assemble(ctx, dbStores)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	return payments, total, nil
}

func (t *StoreService) Assemble(ctx context.Context, stores db.Stores) (types.PaymentStores, error) {

	storeIds := stores.Ids()
	dbStores, err := new(coredb.Store).ListByIds(ctx, storeIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	dbStoresMap := dbStores.AsMap()

	_, ownerIds := dbStores.Ids()

	dbUsers, err := new(coredb.User).ListByIds(ctx, ownerIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	dbUsersMap := dbUsers.AsMap()

	var rsp types.PaymentStores

	for _, x := range stores {

		store := dbStoresMap[x.StoreId]

		y := types.PaymentStore{
			Store:     types.FromDbStore(store, dbUsersMap[store.OwnerId]),
			CreatedAt: timed.SmartTime(x.CreatedAt.Unix()),
			Status:    enum.PaymentStoreStatus(x.Status),
			Xinsh:     x.Xinsh,
			Aliyun:    x.Alipay,
		}

		rsp = append(rsp, &y)
	}

	return rsp, nil

}
