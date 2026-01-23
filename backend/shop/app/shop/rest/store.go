package rest

import (
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	"gitee.com/meepo/backend/shop/app/shop/service"
	"github.com/gin-gonic/gin"
)

func GetStore(c *gin.Context) {

	ctx := c.Request.Context()

	storeId := c.Param("storeId")

	if storeId == "" {
		gind.BadRequestf(c, "storeId is required")
		return
	}

	store, err := new(service.StoreService).FindStoreById(ctx, storeId)
	if err != nil {
		slf.WithError(err).Errorw("GetStore err")
		gind.Error(c, err)
		return
	}

	//if store == nil {
	//	gind.OKM(c, "未知店铺")
	//	return
	//}

	gind.OK(c, store)

}

func AddStoreUser(c *gin.Context) {
	ctx := c.Request.Context()

	storeId := c.Param("storeId")

	if storeId == "" {
		gind.BadRequestf(c, "storeId is required")
		return
	}

	userId := c.GetString("UserId")

	proxyId := c.Query("proxyId")

	store, err := new(service.StoreService).FindStoreById(ctx, storeId)
	if err != nil {
		slf.WithError(err).Errorw("GetStore err")
		gind.Error(c, err)
		return
	}

	if store == nil {
		gind.OKM(c, "未知店铺")
		return
	}

	err = new(service.StoreUserService).AddStoreUser(ctx, storeId, userId)
	if err != nil {
		slf.WithError(err).Errorw("AddStoreUser err")
		gind.Error(c, err)
		return
	}

	if proxyId != "" {
		err := new(service.ProxyUserService).AddUser(ctx, storeId, proxyId, userId)
		if err != nil {
			slf.WithError(err).Errorw("AddProxyUser err")
			gind.Error(c, err)
			return
		}
	}

	gind.OK(c)

}
