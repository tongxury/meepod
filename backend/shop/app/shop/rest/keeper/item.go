package rest

import (
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	keeperservice "gitee.com/meepo/backend/shop/app/shop/service/keeper"
	"github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context) {

	//storeId := c.GetString("StoreId")
	//userId := c.GetString("UserId")

	ctx := c.Request.Context()

	items, err := new(keeperservice.ItemService).ListMetaByIds(ctx, nil, true)
	if err != nil {
		slf.WithError(err).Errorw("ListMetaByIds err")
		gind.Error(c, err)
		return
	}

	gind.OK(c, items)
}
