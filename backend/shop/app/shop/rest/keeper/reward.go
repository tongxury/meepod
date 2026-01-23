package rest

import (
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	keeperservice "gitee.com/meepo/backend/shop/app/shop/service/keeper"
	"github.com/gin-gonic/gin"
)

func ListRewards(c *gin.Context) {
	storeId := c.GetString("StoreId")
	page := conv.Int64(c.DefaultQuery("page", "1"))
	size := conv.Int64(c.DefaultQuery("size", "15"))

	ctx := c.Request.Context()

	rewards, total, err := new(keeperservice.RewardService).ListRewards(ctx, storeId, page, size)
	if err != nil {
		slf.WithError(err).Errorw("ListRewards err")
		gind.Error(c, err)
		return
	}

	gind.Page(c, rewards, page, size, total)
}

func UpdateReward(c *gin.Context) {

	action := c.Query("action")

	storeId := c.GetString("StoreId")
	rewardId := c.Param("id")
	if rewardId == "" {
		gind.BadRequestf(c, "id is required")
		return
	}

	ctx := c.Request.Context()

	var err error
	switch action {
	case "reward":
		err = new(keeperservice.RewardService).Reward(ctx, rewardId, storeId)
	case "reject":
		reason := c.Query("reason")
		err = new(keeperservice.RewardService).Reject(ctx, rewardId, storeId, reason)
	default:

	}

	if err != nil {
		slf.WithError(err).Errorw("Update err", slf.String("action", action))
		gind.Error(c, err)
		return
	}

	newValue, err := new(keeperservice.RewardService).RequireByIdAndStoreId(ctx, rewardId, storeId)
	if err != nil {
		slf.WithError(err).Errorw("RequireByIdAndStoreId err", slf.String("action", action))
		gind.Error(c, err)
		return
	}

	gind.OK(c, newValue)
}
