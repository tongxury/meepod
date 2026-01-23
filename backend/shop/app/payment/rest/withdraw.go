package rest

import (
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	"gitee.com/meepo/backend/shop/app/payment/service"
	"github.com/gin-gonic/gin"
)

func UpdateWithdraw(c *gin.Context) {

	action := c.Query("action")
	orderId := c.Param("id")
	if orderId == "" {
		gind.BadRequestf(c, "id is required")
		return
	}
	ctx := c.Request.Context()

	userId := c.GetString("UserId")
	storeId := c.GetString("StoreId")

	var err error
	switch action {
	case "cancel":
		err = new(service.WithdrawService).Cancel(ctx, userId, storeId, orderId)
	}

	if err != nil {
		slf.WithError(err).Errorw("Update err")
		gind.Error(c, err)
		return
	}

	newOrder, err := new(service.WithdrawService).FindById(ctx, orderId)
	if err != nil {
		slf.WithError(err).Errorw("RequireByStoreIdAndId err")
		gind.Error(c, err)
		return
	}

	gind.OK(c, newOrder)
}

func ListWithdraws(c *gin.Context) {

	storeId := c.GetString("StoreId")
	userId := c.GetString("UserId")

	page := conv.Int64(c.DefaultQuery("page", "1"))
	size := conv.Int64(c.DefaultQuery("size", "15"))

	ctx := c.Request.Context()

	orders, total, err := new(service.WithdrawService).ListWithdraws(ctx, userId, storeId, page, size)

	if err != nil {
		slf.WithError(err).Errorw("ListWithdraws err")
		gind.Error(c, err)
		return
	}

	gind.Page(c, orders, page, size, total)
}

func AddWithdraw(c *gin.Context) {

	ctx := c.Request.Context()

	storeId := c.GetString("StoreId")
	userId := c.GetString("UserId")

	var params struct {
		Amount float64
	}

	if err := c.ShouldBindJSON(&params); err != nil {
		gind.BadRequest(c, err)
		return
	}

	if params.Amount <= 0 {
		gind.BadRequestf(c, "amount should be > 0")
		return
	}

	data, err := new(service.WithdrawService).AddWithdraw(ctx, storeId, userId, params.Amount)

	if err != nil {
		slf.WithError(err).Errorw("AddWithdraw add err", slf.UserId(userId))
		gind.Error(c, err)
		return
	}

	gind.OK(c, data)
}
