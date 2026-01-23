package job

import (
	"context"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/shop/app/payment/db"
	"gitee.com/meepo/backend/shop/app/payment/service"
	"gitee.com/meepo/backend/shop/core/enum"
)

func XinshPayResult() {

	ctx := context.Background()

	topups, _, err := new(db.Topup).List(ctx, db.ListTopupsParams{
		MStatus: []string{enum.TopupStatus_Submitted.Value},
		Size:    100,
	})
	if err != nil {
		slf.WithError(err).Errorw("List Topups err")
		return
	}

	for _, topup := range topups {

		result, err := comp.SDK().Xinsh().GetTradeResult(ctx, topup.Extra.MerchantNo, topup.Id)
		if err != nil {
			slf.WithError(err).Errorw("GetTradeResult err")
			continue
		}

		if result.Paying() {
			continue
		}

		if result.IsSuccess() {

			var err error

			switch topup.Category {
			case enum.TopupCategory_Wallet.Value:
				_, err = new(service.TopupService).ConfirmWalletTopup(ctx, topup.StoreId, topup.Id, topup.UserId, topup.Amount)
			case enum.TopupCategory_Buying.Value:
				_, err = new(service.TopupService).ConfirmBuyingTopup(ctx, topup.StoreId, topup.Id, topup.UserId, topup.Extra.OrderId, topup.Amount)
			}

			if err != nil {
				slf.WithError(err).Errorw("Confirm Topup err", slf.Reflect("topup", topup))
			}
		}

	}
}

func XinshStoreApplyState() {

	ctx := context.Background()

	stores, _, err := new(db.Store).List(ctx, db.ListStoresParams{
		XinshStatus: enum.PaymentStoreStatus_Pending.Value,
	})
	if err != nil {
		slf.WithError(err).Errorw("List err")
		return
	}

	if len(stores) == 0 {
		return
	}
	for _, x := range stores {
		state, err := comp.SDK().Xinsh().GetApplyState(ctx, x.Xinsh.RequestId)
		if err != nil {
			slf.WithError(err).Errorw("GetApplyState err", slf.Reflect("x", x))
			continue
		}

		//("0","待审核")
		//("5","开户请求失败")
		//("6","已提交")
		//("7","审核驳回")
		//("8","审核通过")
		//("9","图片审核驳回");

		stateText := map[string]string{
			"0": "待审核",
			"5": "开户请求失败",
			"6": "已提交",
			"7": "审核驳回",
			"8": "审核通过",
			"9": "图片审核驳回",
		}[state.OpStatus]

		err = new(db.Store).UpdateXinshField(ctx, x.StoreId, "state", stateText)
		if err != nil {
			slf.WithError(err).Errorw("UpdateXinshField err", slf.Reflect("x", x))
			return
		}

		if state.OpStatus == "8" {
			err = new(db.Store).UpdateXinshField(ctx, x.StoreId, "status", enum.PaymentStoreStatus_Confirmed.Value)
			if err != nil {
				slf.WithError(err).Errorw("UpdateXinshField err", slf.Reflect("x", x))
				return
			}
		} else if helper.InSlice(state.OpStatus, []string{"7", "9"}) {
			err = new(db.Store).UpdateXinshField(ctx, x.StoreId, "status", enum.PaymentStoreStatus_Rejected.Value)
			if err != nil {
				slf.WithError(err).Errorw("UpdateXinshField err", slf.Reflect("x", x))
				return
			}

			err = new(db.Store).UpdateXinshField(ctx, x.StoreId, "suggestion", state.Suggestion)
			if err != nil {
				slf.WithError(err).Errorw("UpdateXinshField err", slf.Reflect("x", x))
				return
			}
		}

	}

}
