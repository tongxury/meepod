package rest

import (
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	"gitee.com/meepo/backend/shop/app/shop/service"
	"gitee.com/meepo/backend/shop/core/types"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {

	action := c.Query("action")
	orderId := c.Param("id")
	if orderId == "" {
		gind.BadRequestf(c, "orderId is required")
		return
	}
	ctx := c.Request.Context()

	userId := c.GetString("UserId")
	storeId := c.GetString("StoreId")

	var err error
	switch action {
	case "cancel":
		err = new(service.OrderService).CancelOrder(ctx, userId, orderId)
	}

	if err != nil {
		slf.WithError(err).Errorw("Update err")
		gind.Error(c, err)
		return
	}

	newOrder, err := new(service.OrderService).RequireByStoreIdAndId(ctx, storeId, orderId)
	if err != nil {
		slf.WithError(err).Errorw("RequireByStoreIdAndId err")
		gind.Error(c, err)
		return
	}

	gind.OK(c, newOrder)

}

func GetOrder(c *gin.Context) {

	orderId := c.Param("id")
	if orderId == "" {
		gind.BadRequestf(c, "orderId is required")
		return
	}

	ctx := c.Request.Context()

	storeId := c.GetHeader("StoreId")
	userId := c.GetString("UserId")

	order, err := new(service.OrderService).GetOwnOrder(ctx, userId, storeId, orderId)
	if err != nil {
		slf.WithError(err).Errorw("GetOrder err")
		gind.Error(c, err)
		return
	}

	gind.OK(c, order)

}

func ListOrders(c *gin.Context) {

	category := c.Query("category")

	storeId := c.GetString("StoreId")
	userId := c.GetString("UserId")
	page := conv.Int64(c.DefaultQuery("page", "1"))
	size := conv.Int64(c.DefaultQuery("size", "15"))

	ctx := c.Request.Context()

	var orders types.Orders
	var total int64
	var err error

	switch category {
	case "submitted":
		orders, total, err = new(service.OrderService).ListSubmittedOrders(ctx, userId, storeId, page, size)
	case "toPay":
		orders, total, err = new(service.OrderService).ListToPayOrders(ctx, userId, storeId, page, size)
	case "order":
		orders, total, err = new(service.OrderService).ListMyOrders(ctx, userId, storeId, page, size)
	case "followable":
		orders, total, err = new(service.OrderService).ListFollowableOrders(ctx, userId, storeId, page, size)

	}

	if err != nil {
		slf.WithError(err).Errorw("ListOrders err")
		gind.Error(c, err)
		return
	}

	gind.Page(c, orders, page, size, total)

}

func AddOrder(c *gin.Context) {

	ctx := c.Request.Context()

	storeId := c.GetString("StoreId")
	userId := c.GetString("UserId")

	action := c.Query("action")

	var orderId string
	var err error

	// todo 检查期数 目前前端做限制

	switch action {
	case "follow":
		var params types.FollowOrderForm
		if err := c.ShouldBindJSON(&params); err != nil {
			gind.BadRequest(c, err)
			return
		}

		orderId, err = new(service.OrderService).Follow(ctx, userId, storeId, params.FollowOrderId)
	//case "createGroup":
	//	var params model.CreateGroupOrderForm
	//	if err := c.ShouldBindJSON(&params); err != nil {
	//		gind.BadRequest(c, err)
	//		return
	//	}
	//
	//	orderId, err = new(service.OrderService).CreateGroupAndFirstGroupOrder(ctx, userId, storeId, params.PlanId,
	//		params.TotalVolume, params.Volume, params.Floor, params.Remark, params.RewardRate)
	default:
		var params types.OrderForm
		if err := c.ShouldBindJSON(&params); err != nil {
			gind.BadRequest(c, err)
			return
		}
		orderId, err = new(service.OrderService).AddOrder(ctx, userId, storeId, params.PlanId, params.NeedUpload)
	}

	if err != nil {
		slf.WithError(err).Errorw("AddOrder err")
		gind.Error(c, err)
		return
	}

	gind.OK(c, orderId)

}
