package rest

import (
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	keeperservice "gitee.com/meepo/backend/shop/app/shop/service/keeper"
	"github.com/gin-gonic/gin"
)

func ListProxyUsers(c *gin.Context) {

	proxyId := c.Param("proxyId")

	storeId := c.GetString("StoreId")
	//keeperId := c.GetString("UserId")

	page := conv.Int64(c.DefaultQuery("page", "1"))
	size := conv.Int64(c.DefaultQuery("size", "15"))

	ctx := c.Request.Context()

	users, total, err := new(keeperservice.ProxyUserService).ListProxyUsers(ctx, proxyId, storeId, page, size)
	if err != nil {
		slf.WithError(err).Errorw("ListProxyUsers err")
		gind.Error(c, err)
		return
	}

	gind.Page(c, users, page, size, int64(total))
}
func AddProxyUser(c *gin.Context) {

	storeId := c.GetString("StoreId")
	//keeperId := c.GetString("UserId")

	proxyId := c.Param("proxyId")

	ctx := c.Request.Context()

	userId := c.Query("userId")
	err := new(keeperservice.ProxyUserService).AddUser(ctx, storeId, proxyId, userId)
	if err != nil {
		slf.WithError(err).Errorw("AddUser err")
		gind.Error(c, err)
		return
	}
	gind.OKM(c, "添加成功")
}
