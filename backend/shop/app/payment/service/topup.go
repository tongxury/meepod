package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/kit/components/sdk/helper/mathd"
	"gitee.com/meepo/backend/kit/components/sdk/helper/timed"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/kit/components/third/xinsh"
	"gitee.com/meepo/backend/shop/app/payment/db"
	coredb "gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/types"
	"github.com/go-pg/pg/v10"
)

type TopupService struct {
}

var DuplicateError = errors.New("duplicate")

func (t *TopupService) ConfirmBuyingTopup(ctx context.Context, storeId, topupId, userId, orderId string, amount float64, extra ...map[string]string) (bool, error) {

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		updated, err := new(db.Topup).UpdateToPayed(ctx, tx, topupId, userId)
		if err != nil {
			return err
		}

		if len(extra) > 0 {
			for k, v := range extra[0] {
				if _, err := new(db.Topup).UpdateExtra(ctx, tx, topupId, k, v); err != nil {
					return err
				}
			}
		}

		if !updated {
			return DuplicateError
		}
		// 增加账本
		err = new(db.Account).Incr(ctx, tx, userId, storeId, amount)
		if err != nil {
			return err
		}

		// 消费记录
		payment := db.Payment{
			UserId:      userId,
			StoreId:     storeId,
			BizId:       orderId,
			BizCategory: enum.BizCategory_Order.Value,
			Category:    enum.PaymentCategory_BuyTicket.Value,
			Amount:      amount,
			Status:      enum.PaymentStatus_Payed.Value,
		}
		err = payment.Insert(ctx, tx)
		if err != nil {
			return err
		}

		// 修改订单状态
		_, err = new(coredb.Order).UpdateToPayed(ctx, tx, orderId, false)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		if errors.Is(err, DuplicateError) {
			return false, nil
		}

		return false, xerror.Wrap(err)
	}

	return true, nil
}

func (t *TopupService) ConfirmWalletTopup(ctx context.Context, storeId, topupId, userId string, amount float64, extra ...map[string]string) (bool, error) {

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		updated, err := new(db.Topup).UpdateToPayed(ctx, tx, topupId, userId)
		if err != nil {
			return err
		}

		if len(extra) > 0 {
			for k, v := range extra[0] {
				if _, err := new(db.Topup).UpdateExtra(ctx, tx, topupId, k, v); err != nil {
					return err
				}
			}
		}

		if !updated {
			return DuplicateError
		}

		// 增加账本
		err = new(db.Account).Incr(ctx, tx, userId, storeId, amount)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		if errors.Is(err, DuplicateError) {
			return false, nil
		}

		return false, xerror.Wrap(err)
	}

	return true, nil
}

func (t *TopupService) Cancel(ctx context.Context, userId, storeId, orderId string) error {

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		_, err := new(db.Topup).UpdateToCanceled(ctx, tx, orderId, userId)
		if err != nil {
			return err
		}

		//if canceled {
		//	err := new(db.Payment).Revert(ctx, tx, storeId, orderId, enum.PaymentStatus_Pending.Value)
		//	if err != nil {
		//		return xerror.Wrap(err)
		//	}
		//}

		return nil
	})

	return xerror.Wrap(err)
}

func (t *TopupService) ListTopupOrders(ctx context.Context, userId, storeId string, page, size int64) (types.PaymentTopups, int64, error) {

	orders, total, err := new(db.Topup).List(ctx, db.ListTopupsParams{
		UserId: userId, StoreId: storeId, Page: page, Size: size,
	})
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	rsp, err := t.Assemble(ctx, orders)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	return rsp, total, nil
}

func (t *TopupService) AddRewardTopUp(ctx context.Context, userId, storeId, bizId, bizCategory string, amount float64) (*types.PaymentTopup, error) {

	dbTopup := &db.Topup{
		UserId:      userId,
		StoreId:     storeId,
		Amount:      amount,
		Status:      enum.TopupStatus_Payed.Value,
		Category:    enum.TopupCategory_Reward.Value,
		BizId:       bizId,
		BizCategory: bizCategory,
	}

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {
		inserted, err := dbTopup.Insert(ctx, tx)
		if err != nil {
			return err
		}

		if inserted {
			// 增加账本
			err = new(db.Account).Incr(ctx, tx, userId, storeId, amount)
			if err != nil {
				return err
			}

		}

		return nil
	})

	if err != nil {
		return nil, xerror.Wrap(err)
	}

	topups, err := t.Assemble(ctx, db.Topups{dbTopup})
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return topups[0], nil
}

func (t *TopupService) AddProxyRewardTopUp(ctx context.Context, storeId, userId, bizId string, amount float64) (*types.PaymentTopup, error) {

	//orderCategory := enum.OrderCategory_TopUp.Value

	//exists, err := new(db.Topup).ExistsByStatus(ctx, userId, storeId, enum.TopupStatus_Submitted.Value)
	//if err != nil {
	//	return nil, xerror.Wrap(err)
	//}
	//
	//if exists {
	//	return nil, errorx.UserMessage("订单已存在")
	//}

	dbTopup := &db.Topup{
		UserId:      userId,
		StoreId:     storeId,
		Amount:      amount,
		Status:      enum.TopupStatus_Payed.Value,
		Category:    enum.TopupCategory_ProxyReward.Value,
		BizId:       bizId,
		BizCategory: enum.BizCategory_ProxyReward.Value,
	}

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {
		inserted, err := dbTopup.Insert(ctx, tx)
		if err != nil {
			return err
		}

		if inserted {
			// 增加账本
			err = new(db.Account).Incr(ctx, tx, userId, storeId, amount)
			if err != nil {
				return err
			}

		}

		return nil
	})

	if err != nil {
		return nil, xerror.Wrap(err)
	}

	topups, err := t.Assemble(ctx, db.Topups{dbTopup})
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return topups[0], nil
}

func (t *TopupService) AddWalletTopUp(ctx context.Context, storeId, userId string, amount float64, clientIp string) (*types.PaymentTopup, error) {
	return t.AddTopUp(ctx, enum.TopupCategory_Wallet.Value, storeId, userId, "", "", amount, clientIp)
}

func (t *TopupService) AddBuyingTopUp(ctx context.Context, storeId, userId, orderId, category string, amount float64, clientIp string) (*types.PaymentTopup, error) {
	return t.AddTopUp(ctx, enum.TopupCategory_Buying.Value, storeId, userId, orderId, category, amount, clientIp)
}

func (t *TopupService) AddTopUp(ctx context.Context, topupCategory, storeId, userId, orderId, category string, amount float64, clientIp string) (*types.PaymentTopup, error) {

	dbTopup := &db.Topup{
		UserId:   userId,
		StoreId:  storeId,
		Amount:   amount,
		Status:   enum.TopupStatus_Payed.Value,
		Category: topupCategory,
		Extra: &db.TopupExtra{
			OrderId:  orderId,
			Category: category,
		},
	}

	slf.Debugw("AddTopUp ", slf.Reflect("dbTopup", dbTopup))

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {
		inserted, err := dbTopup.Insert(ctx, tx)
		if err != nil {
			return err
		}

		if inserted {
			// 增加账本
			err = new(db.Account).Incr(ctx, tx, userId, storeId, amount)
			if err != nil {
				return err
			}

			if topupCategory == enum.TopupCategory_Buying.Value {
				// 消费记录
				payment := db.Payment{
					UserId:      userId,
					StoreId:     storeId,
					BizId:       orderId,
					BizCategory: enum.BizCategory_Order.Value,
					Category:    enum.PaymentCategory_BuyTicket.Value,
					Amount:      amount,
					Status:      enum.PaymentStatus_Payed.Value,
				}
				err = payment.Insert(ctx, tx)
				if err != nil {
					return err
				}

				// 修改订单状态
				_, err = new(coredb.Order).UpdateToPayed(ctx, tx, orderId, false)
				if err != nil {
					return err
				}
			}
		}

		// Mock: Skip external payment provider and mark as success immediately
		// payInfo, err := t.getXinshPayUrl(ctx, dbTopup.Id, dbTopup.StoreId, dbTopup.UserId, enum.TopupCategory_Buying.Value, orderId, enum.TopupCategory_Buying.Name, dbTopup.Amount, clientIp)
		// if err != nil {
		// 	return err
		// }

		// _, err = dbTopup.SetExtra(ctx, tx, dbTopup.Id, db.TopupExtra{
		// 	PayUrl: payInfo.PayUrl,
		// 	//PayMethod:    method,
		// 	OrderId:      orderId,
		// 	Category:     topupCategory,
		// 	TradeNo:      payInfo.OrderNo,
		// 	MerchantNo:   payInfo.MerchantNo,
		// 	SellerId:     "",
		// 	BuyerId:      "",
		// 	BuyerLogonId: "",
		// })
		// if err != nil {
		// 	return err
		// }

		// dbTopup.Extra.QrCode = payInfo.PayUrl

		return nil
	})

	if err != nil {
		return nil, xerror.Wrap(err)
	}

	topups, err := t.Assemble(ctx, db.Topups{dbTopup})
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return topups[0], nil
}

func (t *TopupService) FindById(ctx context.Context, id string) (*types.PaymentTopup, error) {

	dbOrder, err := new(db.Topup).FindById(ctx, id)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	orders, err := t.Assemble(ctx, db.Topups{dbOrder})
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return orders[0], nil
}

func (t *TopupService) Assemble(ctx context.Context, orders db.Topups) (types.PaymentTopups, error) {

	_, storeIds, userIds := orders.Ids()

	dbStores, err := new(coredb.Store).ListByIds(ctx, storeIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	dbStoresMap := dbStores.AsMap()

	_, ownerIds := dbStores.Ids()
	userIds = append(userIds, ownerIds...)
	dbUsers, err := new(coredb.User).ListByIds(ctx, userIds)
	if err != nil {
		return nil, xerror.Wrap(err)
	}
	dbUsersMap := dbUsers.AsMap()

	var rsp types.PaymentTopups

	for _, x := range orders {

		user := types.FromDbUser(dbUsersMap[x.UserId])

		dnStore := dbStoresMap[x.StoreId]
		store := types.FromDbStore(dnStore, dbUsersMap[dnStore.OwnerId])

		y := types.PaymentTopup{
			Id:          x.Id,
			User:        user,
			Store:       store,
			Amount:      x.Amount,
			Category:    enum.TopupCategory(x.Category),
			CreatedAt:   timed.SmartTime(x.CreatedAt.Unix()),
			CreatedAtTs: x.CreatedAt.Unix(),
			Status:      enum.TopupStatus(x.Status),
			Payable:     helper.InSlice(x.Status, enum.PayableTopupStatus),
			Payed:       x.Status == enum.PaymentStatus_Payed.Value,
			Cancelable:  helper.InSlice(x.Status, enum.CancelableTopupStatus),
			TimeLeft:    mathd.Max(int64(0), x.CreatedAt.Add(enum.PaymentTimeout).Unix()-time.Now().Unix()),
		}

		y.Payable = helper.InSlice(x.Status, enum.PayableTopupStatus) && y.TimeLeft > 0
		if y.Payable {
			y.PayUrl = x.Extra.PayUrl
			y.PayMethod = x.Extra.PayMethod
			y.QrCode = x.Extra.QrCode
		}

		rsp = append(rsp, &y)
	}

	return rsp, nil
}

func (t *TopupService) getXinshPayUrl(ctx context.Context, topupId, storeId, userId, category, orderId, subject string, amount float64, clientIp string) (*xinsh.PayInfo, error) {

	store, err := new(db.Store).RequireByStoreId(ctx, storeId)
	if err != nil {
		return nil, err
	}

	passbackParams := map[string]any{
		"userId":   userId,
		"storeId":  storeId,
		"topupId":  topupId,
		"category": category,
		"amount":   amount,
		"subject":  "商品购买",
		"orderId":  orderId,
	}

	var params []string
	for k, v := range passbackParams {
		params = append(params, fmt.Sprintf("%s=%v", k, v))
	}

	notifyUrl := fmt.Sprintf("%s/api/payment/v1/xinsh-topup-callback?%s", comp.Flags().GetStr("domain"), strings.Join(params, "&"))

	tradeParams := xinsh.TradeParams{
		MerchantNo: store.Xinsh.MerchantNo,
		OrderId:    topupId,
		Amount:     amount,
		//Type:         "",
		Subject:      "商品购买",
		UserClientIp: clientIp,
		TimeExpire:   enum.PaymentTimeoutMinute,
		NotifyUrl:    notifyUrl,
	}

	payInfo, req, err := comp.SDK().Xinsh().GenerateTradeQrCode(ctx, tradeParams)
	if err != nil {
		slf.WithError(err).Errorw("Xinsh.GenerateTradeQrCode ", slf.Reflect("req", req))
		return nil, xerror.Wrap(err)
	}

	slf.Debugw("GenerateTradeQrCode ", slf.Reflect("tradeParams", tradeParams), slf.Reflect("req", req), slf.Reflect("payInfo", payInfo))

	return payInfo, nil
}

//func (t *TopupService) getPayUrl(ctx context.Context, topupId, storeId, userId, category, orderId, subject string, amount float64) (string, *alipay.TradeParams, error) {
//	token, err := new(AlipayService).GetAuthToken(ctx, storeId)
//	if err != nil {
//		return "", nil, xerror.Wrap(err)
//	}
//
//	passbackParams := map[string]any{
//		"userId":   userId,
//		"storeId":  storeId,
//		"topupId":  topupId,
//		"category": category,
//		"amount":   amount,
//		"subject":  subject,
//		"orderId":  orderId,
//	}
//
//	//providerId := comp.Flags().GetStr("ali.provider.pid")
//
//	params := alipay.TradeParams{
//		NotifyUrl:      fmt.Sprintf("%s/api/payment/v1/topup-callback", comp.Flags().GetStr("domain")),
//		PassbackParams: passbackParams,
//		AppAuthToken:   token,
//		OrderId:        topupId,
//		TotalAmount:    amount,
//		Subject:        subject,
//		TimeExpire:     time.Now().Add(enum.PaymentTimeout).In(timed.LocAsiaShanghai).Format(time.DateTime),
//		ProviderPID:    comp.Flags().AlipayOptions.Income.Pid,
//	}
//
//	url, err := comp.SDK().Alipay().GenerateTradeQrCode(ctx, params)
//	if err != nil {
//		return "", nil, xerror.Wrap(err)
//	}
//
//	slf.Debugw("GenerateTradeQrCode ", slf.Reflect("params", params), slf.String("url", url))
//
//	return url, &params, nil
//}
