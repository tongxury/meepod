package keeperrest

import (
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	"gitee.com/meepo/backend/shop/app/payment/service"
	keeperservice "gitee.com/meepo/backend/shop/app/payment/service/keeper"
	"github.com/gin-gonic/gin"
)

func UpdateWithdraw(c *gin.Context) {

	action := c.Query("action")
	id := c.Param("id")
	if id == "" {
		gind.BadRequestf(c, "id is required")
		return
	}
	ctx := c.Request.Context()

	userId := c.GetString("UserId")
	storeId := c.GetString("StoreId")

	var err error
	switch action {
	case "accept":

		image := c.Query("image")
		if image == "" {
			gind.BadRequestf(c, "image proof is required")
			return
		}

		err = new(keeperservice.WithdrawService).Accept(ctx, userId, storeId, id, image)
	case "reject":
		var params struct {
			Reason string
		}

		if err := c.ShouldBindJSON(&params); err != nil {
			gind.BadRequest(c, err)
			return
		}

		err = new(keeperservice.WithdrawService).Reject(ctx, userId, storeId, id, params.Reason)
	}

	if err != nil {
		slf.WithError(err).Errorw("Update err")
		gind.Error(c, err)
		return
	}

	newOrder, err := new(service.WithdrawService).FindById(ctx, id)
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

	orders, total, err := new(keeperservice.WithdrawService).ListWithdraws(ctx, userId, storeId, page, size)

	if err != nil {
		slf.WithError(err).Errorw("ListWithdraws err")
		gind.Error(c, err)
		return
	}

	gind.Page(c, orders, page, size, total)
}
