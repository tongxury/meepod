package rest

import (
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	"gitee.com/meepo/backend/shop/app/shop/service"
	"github.com/gin-gonic/gin"
)

func ListFeedbacks(c *gin.Context) {

	ctx := c.Request.Context()

	page := conv.Int64(c.DefaultQuery("page", "1"))
	size := conv.Int64(c.DefaultQuery("size", "15"))

	//userId := c.GetString("UserId")
	storeId := c.GetString("StoreId")

	items, total, err := new(service.FeedbackService).List(ctx, storeId, "", page, size)
	if err != nil {
		slf.WithError(err).Errorw("List err")
		gind.Error(c, err)
		return
	}

	gind.Page(c, items, page, size, total)

}
