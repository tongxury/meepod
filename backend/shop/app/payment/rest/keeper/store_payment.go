package keeperrest

import (
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	keeperservice "gitee.com/meepo/backend/shop/app/payment/service/keeper"
	"github.com/gin-gonic/gin"
)

func ListPayments(c *gin.Context) {

	month := c.Query("month")

	storeId := c.GetString("StoreId")
	userId := c.GetString("UserId")

	page := conv.Int64(c.DefaultQuery("page", "1"))
	size := conv.Int64(c.DefaultQuery("size", "15"))

	ctx := c.Request.Context()

	orders, total, err := new(keeperservice.StorePaymentService).ListPayments(ctx, month, userId, storeId, page, size)

	if err != nil {
		slf.WithError(err).Errorw("ListPayments err")
		gind.Error(c, err)
		return
	}

	gind.Page(c, orders, page, size, total)
}
