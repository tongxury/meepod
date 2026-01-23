package job

import (
	"context"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/shop/app/payment/db"
	"gitee.com/meepo/backend/shop/core/enum"
)

func CleanTopups() {

	// 清理过期的未支付的topup

	ctx := context.Background()

	err := new(db.Topup).CleanTimeoutTopups(ctx, enum.PaymentTimeout)
	if err != nil {
		slf.WithError(err).Errorw("CleanTimeoutTopups err")
		return
	}

}
