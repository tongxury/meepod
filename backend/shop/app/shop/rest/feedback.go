package rest

import (
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	"gitee.com/meepo/backend/shop/app/shop/service"
	"github.com/gin-gonic/gin"
)

func UpdateFeedback(c *gin.Context) {

	action := c.Query("action")
	id := c.Param("id")
	if id == "" {
		gind.BadRequestf(c, "id is required")
		return
	}
	ctx := c.Request.Context()

	//userId := c.GetString("UserId")
	storeId := c.GetString("StoreId")

	var err error
	switch action {
	case "resolved":
		err = new(service.FeedbackService).Resolve(ctx, storeId, id)
	}

	if err != nil {
		slf.WithError(err).Errorw("Update err")
		gind.Error(c, err)
		return
	}

	newOrder, err := new(service.FeedbackService).FindById(ctx, storeId, id)
	if err != nil {
		slf.WithError(err).Errorw("FindById err")
		gind.Error(c, err)
		return
	}

	gind.OK(c, newOrder)

}

func AddFeedback(c *gin.Context) {

	ctx := c.Request.Context()

	var params struct {
		Text string `binding:"required"`
	}

	if err := c.ShouldBindJSON(&params); err != nil {
		gind.BadRequest(c, err)
		return
	}

	userId := c.GetString("UserId")
	storeId := c.GetString("StoreId")

	err := new(service.FeedbackService).Add(ctx, storeId, userId, params.Text)
	if err != nil {
		slf.WithError(err).Errorw("Add err")
		gind.Error(c, err)
		return
	}

	gind.OK(c)

}

func ListFeedbacks(c *gin.Context) {

	ctx := c.Request.Context()

	page := conv.Int64(c.DefaultQuery("page", "1"))
	size := conv.Int64(c.DefaultQuery("size", "15"))

	userId := c.GetString("UserId")
	storeId := c.GetString("StoreId")

	items, total, err := new(service.FeedbackService).List(ctx, storeId, userId, page, size)
	if err != nil {
		slf.WithError(err).Errorw("List err")
		gind.Error(c, err)
		return
	}

	gind.Page(c, items, page, size, total)

}
