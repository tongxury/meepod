package rest

import (
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	keeperservice "gitee.com/meepo/backend/shop/app/shop/service/keeper"
	"github.com/gin-gonic/gin"
	"time"
)

func UpdateProxy(c *gin.Context) {
	storeId := c.GetString("StoreId")
	keeperId := c.GetString("UserId")

	action := c.Query("action")

	proxyId := c.Param("proxyId")

	ctx := c.Request.Context()

	var err error
	switch action {
	case "updateRewardRate":
		rewardRate := conv.Float64(c.Query("rewardRate"))
		err = new(keeperservice.ProxyService).UpdateRewardRate(ctx, keeperId, storeId, proxyId, rewardRate)
	case "delete":
		err = new(keeperservice.ProxyService).Delete(ctx, keeperId, storeId, proxyId)
	case "recover":
		err = new(keeperservice.ProxyService).Recover(ctx, keeperId, storeId, proxyId)
	default:
	}

	if err != nil {
		slf.WithError(err).Errorw("Update err")
		gind.Error(c, err)
		return
	}

	item, err := new(keeperservice.ProxyService).FindById(ctx, proxyId)
	if err != nil {
		slf.WithError(err).Errorw("FindById err")
		gind.Error(c, err)
		return
	}
	gind.OK(c, item)
}

func AddProxy(c *gin.Context) {

	// todo
	time.Sleep(1 * time.Second)

	var params struct {
		UserId     string
		RewardRate string `binding:"required"`
	}

	if err := c.ShouldBindJSON(&params); err != nil {
		gind.BadRequest(c, err)
		return
	}

	keeperId := c.GetString("UserId")
	storeId := c.GetString("StoreId")

	ctx := c.Request.Context()

	err := new(keeperservice.ProxyService).AddProxy(ctx, keeperId, storeId, params.UserId, conv.Float64(params.RewardRate))
	if err != nil {
		slf.WithError(err).Errorw("AddProxy err")
		gind.Error(c, err)
		return
	}

	gind.OK(c)

}

func ListProxies(c *gin.Context) {

	page := conv.Int64(c.DefaultQuery("page", "1"))
	size := conv.Int64(c.DefaultQuery("size", "15"))

	keeperId := c.GetString("UserId")
	storeId := c.GetString("StoreId")

	ctx := c.Request.Context()

	items, total, err := new(keeperservice.ProxyService).ListProxies(ctx, keeperId, storeId, page, size)
	if err != nil {
		slf.WithError(err).Errorw("ListStores err")
		gind.Error(c, err)
		return
	}
	gind.Page(c, items, page, size, total)
}
