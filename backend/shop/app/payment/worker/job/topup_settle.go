package job

import (
	"context"
	"fmt"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/helper/mathd"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/components/third/alipay"
	"gitee.com/meepo/backend/shop/app/payment/db"
	"gitee.com/meepo/backend/shop/app/payment/service"
	"gitee.com/meepo/backend/shop/core/enum"
)

func Settle() {

	ctx := context.Background()

	topups, _, err := new(db.Topup).List(ctx, db.ListTopupsParams{
		MStatus:     []string{enum.TopupStatus_Payed.Value},
		SyncSettled: 1,
		Page:        1, Size: 1,
	})
	if err != nil {
		slf.WithError(err).Errorw("List err")
		return
	}

	if len(topups) == 0 {
		return
	}

	topup := topups[0]

	token, err := new(service.AlipayService).GetAuthToken(ctx, topup.StoreId)
	if err != nil {
		slf.WithError(err).Errorw("")
		return
	}

	err = comp.SDK().Alipay().Bind(ctx, alipay.BindParams{
		Receivers: []alipay.Receiver{{
			Type:          "userId",
			Account:       comp.Flags().AlipayOptions.Income.Pid,
			Name:          comp.Flags().AlipayOptions.Income.Name,
			Memo:          comp.Flags().AlipayOptions.Income.Memo,
			LoginName:     comp.Flags().AlipayOptions.Income.LoginName,
			BindLoginName: comp.Flags().AlipayOptions.Income.BindLoginName,
		}},
		AppAuthToken: token,
	})
	if err != nil {
		slf.WithError(err).Errorw("")
		return
	}

	incomeRate := comp.Flags().AlipayOptions.Income.Rate

	amount := mathd.Max(mathd.ToFixed4(topup.Amount*incomeRate), 0.01)

	_, err = comp.SDK().Alipay().Settle(ctx, alipay.SettleParams{
		AppAuthToken: token,
		OrderId:      topup.Id,
		TradeNo:      topup.Extra.TradeNo,
		TotalAmount:  topup.Amount,
		RoyaltyParameters: []alipay.RoyaltyParameters{{
			RoyaltyType: "transfer",
			//TransOut:     topup.Extra.SellerId,
			//TransOutType: "userId",
			TransInType: "userId",
			TransIn:     comp.Flags().AlipayOptions.Income.Pid,
			Amount:      amount,
			Desc: fmt.Sprintf("分账给%s(%s)", comp.Flags().AlipayOptions.Income.LoginName,
				comp.Flags().AlipayOptions.Income.Pid),
			RoyaltyScene: "服务商佣金",
			TransInName:  comp.Flags().AlipayOptions.Income.Name,
		},
		},
	})
	if err != nil {
		slf.WithError(err).Errorw("Settle err")
		return
	}

	_, err = new(db.Topup).UpdateToSynced(ctx, []string{topup.Id}, "sync_settled")
	if err != nil {
		slf.WithError(err).Errorw("UpdateToSynced err")
		return
	}
}
