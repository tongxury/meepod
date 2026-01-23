package adminrest

import (
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	adminserivce "gitee.com/meepo/backend/shop/app/shop/service/admin"
	"github.com/gin-gonic/gin"
)

func ListOrders(c *gin.Context) {

	id := c.Query("id")
	storeId := c.Query("storeId")
	status := c.Query("status")

	page := conv.Int64(c.DefaultQuery("page", "1"))
	size := conv.Int64(c.DefaultQuery("size", "15"))

	ctx := c.Request.Context()

	stores, total, err := new(adminserivce.OrderService).ListOrders(ctx, id, storeId, status, page, size)
	if err != nil {
		gind.Error(c, err)
		slf.WithError(err).Errorw("ListOrders err")
		return
	}

	gind.Page(c, stores, page, size, total)
}
