package rest

import (
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	paymentdb "gitee.com/meepo/backend/shop/app/payment/db"
	coredb "gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"github.com/gin-gonic/gin"
)

func ListCounters(c *gin.Context) {

	ctx := c.Request.Context()

	storeId := c.GetString("StoreId")
	toHandleOrderCount, err := new(coredb.Order).CountByStatus(ctx, storeId, enum.ToHandleByKeeperOrderStatus)
	if err != nil {
		slf.WithError(err).Errorw("CountByStatus err")
	}

	toHandleOrderGroupCount, err := new(coredb.OrderGroup).CountByStatus(ctx, storeId, enum.ToHandleByKeeperOrderGroupStatus)
	if err != nil {
		slf.WithError(err).Errorw("CountByStatus err")
	}

	toHandleWithdrawCount, err := new(paymentdb.Withdraw).CountByStatus(ctx, storeId, enum.ToHandleByKeeperWithdrawStatus)
	if err != nil {
		slf.WithError(err).Errorw("CountByStatus err")
	}
	toHandleRewardCount, err := new(coredb.Reward).CountByStatus(ctx, storeId, enum.ToHandleByKeeperRewardStatus)
	if err != nil {
		slf.WithError(err).Errorw("CountByStatus err")
	}

	feedbackCount, err := new(coredb.Feedback).CountByStatus(ctx, storeId, enum.ResolvableFeedbackStatus)
	if err != nil {
		slf.WithError(err).Errorw("CountByStatus err")
	}

	gind.OK(c,
		struct {
			Ticket   int64 `json:"ticket,omitempty"`
			Account  int64 `json:"account,omitempty"`
			Profile  int64 `json:"profile,omitempty"`
			Feedback int64 `json:"feedback,omitempty"`
		}{
			Ticket:   toHandleOrderCount + toHandleOrderGroupCount,
			Account:  toHandleWithdrawCount + toHandleRewardCount,
			Profile:  feedbackCount,
			Feedback: feedbackCount,
		})
}
