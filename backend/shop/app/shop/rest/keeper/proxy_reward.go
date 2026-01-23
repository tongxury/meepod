package rest

import (
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	keeperservice "gitee.com/meepo/backend/shop/app/shop/service/keeper"
	"github.com/gin-gonic/gin"
)

func ListProxyRewards(c *gin.Context) {

	month := c.Query("month")
	cat := c.Query("cat")

	storeId := c.GetString("StoreId")
	//keeperId := c.GetString("UserId")

	page := conv.Int64(c.DefaultQuery("page", "1"))
	size := conv.Int64(c.DefaultQuery("size", "15"))

	ctx := c.Request.Context()

	users, total, err := new(keeperservice.ProxyRewardService).ListProxyRewards(ctx, storeId, month, cat, page, size)
	if err != nil {
		slf.WithError(err).Errorw("ListProxyUsers err")
		gind.Error(c, err)
		return
	}

	gind.Page(c, users, page, size, int64(total))
}

func UpdateProxyReward(c *gin.Context) {
	storeId := c.GetString("StoreId")
	keeperId := c.GetString("UserId")

	action := c.Query("action")

	rewardId := c.Param("rewardId")

	ctx := c.Request.Context()

	var err error
	switch action {
	case "pay":
		err = new(keeperservice.ProxyRewardService).Pay(ctx, keeperId, storeId, rewardId)
	default:
	}

	if err != nil {
		slf.WithError(err).Errorw("Update err")
		gind.Error(c, err)
		return
	}

	item, err := new(keeperservice.ProxyRewardService).FindById(ctx, rewardId)
	if err != nil {
		slf.WithError(err).Errorw("FindById err")
		gind.Error(c, err)
		return
	}
	gind.OK(c, item)
}
