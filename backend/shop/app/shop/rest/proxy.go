package rest

import (
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	"gitee.com/meepo/backend/shop/app/shop/service"
	"github.com/gin-gonic/gin"
)

func GetProxy(c *gin.Context) {

	ctx := c.Request.Context()

	id := c.Param("id")

	if id != "me" {
		gind.BadRequestf(c, "no permission")
		return
	}

	userId := c.GetString("UserId")
	storeId := c.GetString("StoreId")

	proxy, err := new(service.ProxyService).FindByUserId(ctx, storeId, userId)
	if err != nil {
		slf.WithError(err).Errorw("FindByUserId err")
		gind.Error(c, err)
		return
	}

	gind.OK(c, proxy)

}
