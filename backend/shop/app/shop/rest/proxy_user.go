package rest

import (
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	"gitee.com/meepo/backend/shop/app/shop/service"
	"github.com/gin-gonic/gin"
)

func ListProxyUsers(c *gin.Context) {

	proxyId := c.Param("id")

	if proxyId != "me" {
		gind.BadRequestf(c, "no permission")
		return
	}

	storeId := c.GetString("StoreId")
	userId := c.GetString("UserId")

	page := conv.Int64(c.DefaultQuery("page", "1"))
	size := conv.Int64(c.DefaultQuery("size", "15"))

	ctx := c.Request.Context()

	users, total, err := new(service.ProxyUserService).ListProxyUsers(ctx, userId, storeId, page, size)
	if err != nil {
		slf.WithError(err).Errorw("ListProxyUsers err")
		gind.Error(c, err)
		return
	}

	gind.Page(c, users, page, size, int64(total))
}
