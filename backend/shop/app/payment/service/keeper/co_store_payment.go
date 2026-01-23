package keeperservice

import (
	"context"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/helper/mathd"
	"gitee.com/meepo/backend/kit/components/sdk/helper/timed"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/app/payment/db"
	coredb "gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/types"
	"github.com/go-pg/pg/v10"
)

type CoStorePaymentService struct {
}

func (t *CoStorePaymentService) Topup(ctx context.Context, storeId, toStoreId string, amount float64) error {

	payment := db.CoStorePayment{
		StoreId:   toStoreId,
		CoStoreId: storeId,
		Category:  enum.CoStorePaymentCategory_TopUp.Value,
		Amount:    amount,
		Status:    enum.CoStorePaymentStatus_Confirmed.Value,
	}

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {
		err := payment.Insert(ctx, tx)
		return err
	})

	return xerror.Wrap(err)
}

func (t *CoStorePaymentService) ListPayments(ctx context.Context, month, cat, keeperId, storeId, coStoreId string, page, size int64) (types.CoStorePayments, int64, error) {

	dbPayments, total, err := new(db.CoStorePayment).List(ctx, db.ListCoStorePaymentsParams{
		StoreId: storeId, CoStoreId: coStoreId, Month: month,
		Category: cat,
		Page:     page, Size: size,
	})
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	payments, err := t.Assemble(ctx, dbPayments)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	return payments, total, nil
}

func (t *CoStorePaymentService) Return(ctx context.Context, storeId, coStoreId, proofImage string) error {

	wallet, err := new(db.CoStoreWallet).FindByStoreIdAndCoStoreId(ctx, storeId, coStoreId)
	if err != nil {
		return xerror.Wrap(err)
	}
	if wallet != nil {
		if wallet.Balance > 0 {

			payment := db.CoStorePayment{
				StoreId:   storeId,
				CoStoreId: coStoreId,
				Category:  enum.CoStorePaymentCategory_Return.Value,
				Amount:    -wallet.Balance,
				Status:    enum.CoStorePaymentStatus_Confirmed.Value,
				Extra: db.CoStorePaymentExtra{
					ProofImage: proofImage,
				},
			}

			err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {
				err := payment.Insert(ctx, tx)
				return err
			})

			if err != nil {
				return xerror.Wrap(err)
			}

		}
	}

	return nil
}

func (t *CoStorePaymentService) PaySwitchOrder(ctx context.Context, storeId, toStoreId, bizCategory, bizId string, amount float64) error {

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		// 扣减在合作店铺的余额
		coStorePayment := &db.CoStorePayment{
			StoreId:     storeId,
			CoStoreId:   toStoreId,
			BizId:       bizId,
			BizCategory: bizCategory,
			Category:    enum.CoStorePaymentCategory_SwitchDecr.Value,
			Amount:      -amount,
			Status:      enum.CoStorePaymentStatus_Confirmed.Value,
		}

		err := coStorePayment.Insert(ctx, tx)
		if err != nil {
			return err
		}

		// 合作转单佣金
		coStoreReward := &db.CoStorePayment{
			StoreId:     storeId,
			CoStoreId:   toStoreId,
			BizId:       bizId,
			BizCategory: bizCategory,
			Category:    enum.CoStorePaymentCategory_SwitchFee.Value,
			Amount:      mathd.Min(amount*enum.CoStoreRewardRate, enum.CoStoreRewardMax),
			Status:      enum.CoStorePaymentStatus_Confirmed.Value,
		}

		err = coStoreReward.Insert(ctx, tx)
		if err != nil {
			return err
		}

		// 扣减平台转单服务费
		storePayment := db.StorePayment{
			StoreId:     toStoreId,
			BizId:       bizId,
			BizCategory: bizCategory,
			Category:    enum.StorePaymentCategory_SwichFee.Value,
			Amount:      -mathd.ToFixed2(amount * enum.CoStoreRewardRate),
			Status:      enum.StorePaymentStatus_Confirmed.Value,
			Extra: db.StorePaymentExtra{
				FromStoreId: storeId,
			},
		}

		err = storePayment.Insert(ctx, tx)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		//slf.WithError(err).Errorw("coStorePayment insert err")
		return xerror.Wrap(err)
	}

	return nil
}

func (t *CoStorePaymentService) Assemble(ctx context.Context, payments db.CoStorePayments) (types.CoStorePayments, error) {

	_, storeIds, coStoreIds := payments.Ids()
	storeIds = append(storeIds, coStoreIds...)
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

	var rsp types.CoStorePayments

	for _, x := range payments {

		store := dbStoresMap[x.StoreId]
		storeOwner := dbUsersMap[store.OwnerId]

		coStore := dbStoresMap[x.CoStoreId]
		coStoreOwner := dbUsersMap[coStore.OwnerId]

		y := types.CoStorePayment{
			Id:          x.Id,
			Store:       types.FromDbStore(store, storeOwner),
			CoStore:     types.FromDbStore(coStore, coStoreOwner),
			Category:    enum.CoStorePaymentCategory(x.Category),
			BizCategory: enum.BizCategory(x.BizCategory),
			BizId:       x.BizId,
			Amount:      x.Amount,
			CreatedAtTs: x.CreatedAt.Unix(),
			CreatedAt:   timed.SmartTime(x.CreatedAt.Unix()),
			Status:      enum.CoStorePaymentStatus(x.Status),
		}

		rsp = append(rsp, &y)
	}

	return rsp, nil

}
