package rest

import (
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	"gitee.com/meepo/backend/shop/app/shop/service"
	"gitee.com/meepo/backend/shop/core/types"
	"github.com/gin-gonic/gin"
)

func ListItems(c *gin.Context) {

	storeId := c.GetString("StoreId")
	ctx := c.Request.Context()

	var items types.ItemStates
	var err error

	if storeId == "" {
		items, err = new(service.ItemService).ListAllItems(ctx)
	} else {
		items, err = new(service.ItemService).ListStoreItems(ctx, storeId)

	}

	if err != nil {
		slf.WithError(err).Errorw("ListStoreItems err")
		gind.Error(c, err)
		return
	}

	gind.OK(c, items, "")
}

func GetCurrentIssue(c *gin.Context) {

	var params struct {
		ItemId string `form:"itemId" binding:"required"`
	}

	if err := c.ShouldBindQuery(&params); err != nil {
		gind.AbortWithCode(c, 10400, err)
		return
	}

	ctx := c.Request.Context()

	current, err := new(service.IssueService).FindCurrentIssue(ctx, params.ItemId)
	if err != nil {
		slf.WithError(err).Errorw("Current err")
		gind.Error(c, err)
		return
	}

	gind.OK(c, current, "")

}
