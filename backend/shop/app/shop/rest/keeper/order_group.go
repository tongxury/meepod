package rest

import (
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	keeperservice "gitee.com/meepo/backend/shop/app/shop/service/keeper"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"github.com/gin-gonic/gin"
)

func UpdateOrderGroup(c *gin.Context) {

	action := c.Query("action")

	ctx := c.Request.Context()

	orderId := c.Param("orderId")
	if orderId == "" {
		gind.BadRequestf(c, "orderId is required")
		return
	}

	storeId := c.GetString("StoreId")

	var err error
	switch action {

	case "accept":
		err = new(keeperservice.OrderGroupService).Accept(ctx, storeId, orderId)
	case "reject":
		reasonId := c.Query("reasonId")
		if orderId == "" {
			gind.BadRequestf(c, "reasonId is required")
			return
		}
		err = new(keeperservice.OrderGroupService).Reject(ctx, storeId, orderId, reasonId)
	case "switch":
		toStoreId := c.Query("toStoreId")
		if toStoreId == "" {
			gind.BadRequestf(c, "toStoreId is required")
			return
		}
		err = new(keeperservice.OrderGroupService).Switch(ctx, storeId, orderId, toStoreId)
	case "ticket":
		var params struct {
			Images []string
		}

		if err := c.ShouldBindJSON(&params); err != nil {
			gind.BadRequestf(c, err.Error())
			return
		}

		err = new(keeperservice.OrderGroupService).Ticket(ctx, storeId, orderId, params.Images)
	}

	if err != nil {
		slf.WithError(err).Errorw("UpdateOrder err")
		gind.Error(c, err)
		return
	}

	newOrder, err := new(keeperservice.OrderGroupService).RequireByStoreIdAndId(ctx, storeId, orderId)
	if err != nil {
		slf.WithError(err).Errorw("RequireByStoreIdAndId err")
		gind.Error(c, err)
		return
	}

	gind.OK(c, newOrder)
}

func GetOrderGroup(c *gin.Context) {

	ctx := c.Request.Context()

	orderId := c.Param("orderId")
	if orderId == "" {
		gind.BadRequestf(c, "orderId is required")
		return
	}

	bizCategory := c.Query("biz_category")
	if bizCategory == enum.BizCategory_GroupShare.Value {

		share, err := new(db.OrderGroupShare).RequireById(ctx, orderId)
		if err != nil {
			gind.Error(c, err)
			return
		}
		orderId = share.GroupId
	}

	userId := c.GetString("UserId")
	storeId := c.GetString("StoreId")

	group, err := new(keeperservice.OrderGroupService).GetGroup(ctx, userId, storeId, orderId)
	if err != nil {
		slf.WithError(err).Errorw("GetGroup err")
		gind.Error(c, err)
		return
	}

	gind.OK(c, group)
}

func ListOrderGroupShares(c *gin.Context) {

	orderId := c.Param("orderId")
	if orderId == "" {
		gind.BadRequestf(c, "orderId is required")
		return
	}

	storeId := c.GetString("StoreId")
	keeperId := c.GetString("UserId")
	page := conv.Int64(c.DefaultQuery("page", "1"))
	//size := conv.Int64(c.DefaultQuery("size", "15"))

	if page > 1 {
		gind.Page(c, nil, 1, 1, 1)
		return
	}

	ctx := c.Request.Context()

	joiners, err := new(keeperservice.OrderGroupShareService).ListGroupShares(ctx, orderId, keeperId, storeId)

	if err != nil {
		slf.WithError(err).Errorw("ListGroupJoiners err")
		gind.Error(c, err)
		return
	}

	gind.Page(c, joiners, 1, 1, 1)
}

func ListOrderGroups(c *gin.Context) {

	storeId := c.GetString("StoreId")
	//userId := c.GetString("UserId")
	page := conv.Int64(c.DefaultQuery("page", "1"))
	size := conv.Int64(c.DefaultQuery("size", "15"))

	ctx := c.Request.Context()

	orders, total, err := new(keeperservice.OrderGroupService).ListOrderGroups(ctx, storeId, page, size)

	if err != nil {
		slf.WithError(err).Errorw("ListOrderGroups err")
		gind.Error(c, err)
		return
	}

	gind.Page(c, orders, page, size, total)

}
