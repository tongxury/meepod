package rest

import (
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	"gitee.com/meepo/backend/shop/app/shop/service"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/types"
	"github.com/gin-gonic/gin"
)

func AddOrderGroup(c *gin.Context) {

	var params types.CreateGroupOrderForm
	if err := c.ShouldBindJSON(&params); err != nil {
		gind.BadRequest(c, err)
		return
	}

	storeId := c.GetString("StoreId")
	userId := c.GetString("UserId")

	ctx := c.Request.Context()

	orderId, err := new(service.OrderGroupService).AddGroup(ctx, userId, storeId, params.PlanId,
		params.TotalVolume, params.Floor, params.Remark, params.RewardRate)

	if err != nil {
		slf.WithError(err).Errorw("AddGroup err")
		gind.Error(c, err)
		return
	}

	gind.OK(c, orderId)
}

func AddGroupShare(c *gin.Context) {

	ctx := c.Request.Context()

	storeId := c.GetString("StoreId")
	userId := c.GetString("UserId")

	groupId := c.Param("id")
	if groupId == "" {
		gind.BadRequestf(c, "id is required")
		return
	}

	var params types.JoinGroupOrderForm
	if err := c.ShouldBindJSON(&params); err != nil {
		gind.BadRequest(c, err)
		return
	}
	if params.Volume == 0 {
		params.Volume = 1
	}

	orderId, err := new(service.OrderGroupService).JoinGroup(ctx, userId, storeId, groupId, params.Volume)

	if err != nil {
		slf.WithError(err).Errorw("JoinGroup err")
		gind.Error(c, err)
		return
	}

	gind.OK(c, orderId)

}

func ListOrderGroupShares(c *gin.Context) {

	groupId := c.Param("id")
	if groupId == "" {
		gind.BadRequestf(c, "id is required")
		return
	}

	storeId := c.GetString("StoreId")
	userId := c.GetString("UserId")
	page := conv.Int64(c.DefaultQuery("page", "1"))
	//size := conv.Int64(c.DefaultQuery("size", "15"))

	if page > 1 {
		gind.Page(c, nil, 1, 1, 1)
		return
	}

	ctx := c.Request.Context()

	joiners, err := new(service.OrderGroupShareService).ListGroupShares(ctx, groupId, userId, storeId)

	if err != nil {
		slf.WithError(err).Errorw("ListGroupJoiners err")
		gind.Error(c, err)
		return
	}

	gind.Page(c, joiners, 1, 1, 1)
}

//func ListOrderGroupOrders(c *gin.Context) {
//
//	groupId := c.Param("id")
//	if groupId == "" {
//		gind.BadRequestf(c, "id is required")
//		return
//	}
//
//	storeId := c.GetString("StoreId")
//	userId := c.GetString("UserId")
//	page := conv.Int64(c.DefaultQuery("page", "1"))
//	size := conv.Int64(c.DefaultQuery("size", "15"))
//
//	ctx := c.Request.Context()
//
//	orders, err := new(service.OrderGroupShareService).ListGroupJoiners(ctx, groupId, userId, storeId)
//
//	if err != nil {
//		slf.WithError(err).Errorw("ListGroupJoiners err")
//		gind.Error(c, err)
//		return
//	}
//
//	gind.Page(c, orders, page, size, 0)
//}

func GetOrderGroup(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		gind.BadRequestf(c, "id is required")
		return
	}

	ctx := c.Request.Context()

	bizCategory := c.Query("biz_category")
	if bizCategory == enum.BizCategory_GroupShare.Value {

		share, err := new(db.OrderGroupShare).RequireById(ctx, id)
		if err != nil {
			gind.Error(c, err)
			return
		}
		id = share.GroupId
	}

	userId := c.GetString("UserId")

	group, err := new(service.OrderGroupService).GetGroup(ctx, userId, id)
	if err != nil {
		slf.WithError(err).Errorw("GetGroup err")
		gind.Error(c, err)
		return
	}

	gind.OK(c, group)
}

func ListOrderGroups(c *gin.Context) {

	category := c.Query("category")

	storeId := c.GetString("StoreId")
	userId := c.GetString("UserId")
	page := conv.Int64(c.DefaultQuery("page", "1"))
	size := conv.Int64(c.DefaultQuery("size", "15"))

	ctx := c.Request.Context()

	var orders types.OrderGroups
	var total int64
	var err error
	switch category {
	case "createdByMe":
		orders, total, err = new(service.OrderGroupService).ListCreatedOrderGroups(ctx, userId, storeId, page, size)
	case "my":
		orders, total, err = new(service.OrderGroupService).ListMyOrderGroups(ctx, userId, storeId, page, size)
	case "joinable":
		orders, total, err = new(service.OrderGroupService).ListJoinableGroups(ctx, userId, storeId, page, size)
	}

	if err != nil {
		slf.WithError(err).Errorw("ListOrderGroups err")
		gind.Error(c, err)
		return
	}

	gind.Page(c, orders, page, size, total)

}
