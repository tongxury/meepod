package rest

import (
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/helper/timed"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	"gitee.com/meepo/backend/shop/app/shop/service"
	keeperservice "gitee.com/meepo/backend/shop/app/shop/service/keeper"
	"gitee.com/meepo/backend/shop/core/enum"
	"github.com/gin-gonic/gin"
	"time"
)

func UpdateOrder(c *gin.Context) {

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
	case "pay":
	//err = new(keeperservice.OrderService).Pay(ctx, storeId, orderId)
	case "switch":
		toStoreId := c.Query("toStoreId")
		if toStoreId == "" {
			gind.BadRequestf(c, "toStoreId is required")
			return
		}
		err = new(keeperservice.OrderService).Switch(ctx, storeId, orderId, toStoreId)
	case "accept":
		err = new(keeperservice.OrderService).Accept(ctx, storeId, orderId)
	case "reject":
		reasonId := c.Query("reasonId")
		if orderId == "" {
			gind.BadRequestf(c, "reasonId is required")
			return
		}
		err = new(keeperservice.OrderService).Reject(ctx, storeId, orderId, reasonId)
	case "ticket":
		var params struct {
			Images []string
		}

		if err := c.ShouldBindJSON(&params); err != nil {
			gind.BadRequestf(c, err.Error())
			return
		}

		err = new(keeperservice.OrderService).Ticket(ctx, storeId, orderId, params.Images)
	}

	if err != nil {
		slf.WithError(err).Errorw("UpdateOrder err")
		gind.Error(c, err)
		return
	}

	newOrder, err := new(keeperservice.OrderService).RequireByStoreIdAndId(ctx, storeId, orderId)
	if err != nil {
		slf.WithError(err).Errorw("RequireByStoreIdAndId err")
		gind.Error(c, err)
		return
	}

	gind.OK(c, newOrder)
}

func GetOrder(c *gin.Context) {

	orderId := c.Param("orderId")
	if orderId == "" {
		gind.BadRequestf(c, "orderId is required")
		return
	}

	ctx := c.Request.Context()

	keeperId := c.GetString("UserId")
	storeId := c.GetString("StoreId")

	order, err := new(keeperservice.OrderService).GetOrder(ctx, keeperId, storeId, orderId)
	if err != nil {
		slf.WithError(err).Errorw("GetOrder err")
		gind.Error(c, err)
		return
	}

	gind.OK(c, order)

}

func ListOrders(c *gin.Context) {

	dateRange := c.Query("dateRange")

	var from, to *time.Time

	now := time.Now()

	switch dateRange {
	case "today":
		tmp := timed.DateStart(now)
		from = &tmp
	case "yesterday":
		yesterday := now.AddDate(0, 0, -1)
		tmpStart := timed.DateStart(yesterday)
		from = &tmpStart
		tmpEnd := timed.DateEnd(yesterday)
		to = &tmpEnd
	case "last30days":
		tmp := now.AddDate(0, 0, -30)
		from = &tmp
	case "last60days":
		tmp := now.AddDate(0, 0, -60)
		from = &tmp
	default:
	}

	category := c.Query("status")

	var status []string
	switch category {
	case "toHandle":
		status = enum.ToHandleByKeeperOrderStatus
	case "toPay":
		status = enum.PayableStatus
	case "toAccept":
		status = enum.AcceptableStatus
	case "ticketed":
		status = enum.TicketedStatus
	case "timeout":
		status = []string{enum.OrderStatus_Timeout.Value}
	case "rejected":
		status = []string{enum.OrderStatus_Rejected.Value}
	default:
	}

	itemId := c.Query("item")

	storeId := c.GetString("StoreId")

	page := conv.Int64(c.DefaultQuery("page", "1"))
	size := conv.Int64(c.DefaultQuery("size", "20"))

	ctx := c.Request.Context()

	orders, total, err := new(keeperservice.OrderService).ListOrders(ctx, storeId, itemId, status, from, to, page, size)

	if err != nil {
		slf.WithError(err).Errorw("ListOrders err")
		gind.Error(c, err)
		return
	}

	gind.Page(c, orders, page, size, total)

}

func ListOrderFilters(c *gin.Context) {

	storeId := c.GetString("StoreId")

	ctx := c.Request.Context()

	items, err := new(service.ItemService).ListStoreItems(ctx, storeId)
	if err != nil {
		slf.WithError(err).Errorw("ListMetaByIds err")
	}

	var itemOptions []gin.H
	itemOptions = append(itemOptions, gin.H{"name": "全部彩种", "value": ""})
	for _, x := range items {
		itemOptions = append(itemOptions, gin.H{"name": x.Name, "value": x.Id})
	}

	gind.OK(c, gin.H{
		"order": gin.H{
			"dateRange": []gin.H{
				{"name": "全部时间", "value": ""},
				{"name": "今天", "value": "today"},
				{"name": "昨天", "value": "yesterday"},
				{"name": "过去30天", "value": "last30days"},
				{"name": "过去60天", "value": "last60days"},
			},
			"category": []gin.H{
				{"name": "全部状态", "value": ""},
				{"name": "待处理", "value": "toHandle"},
				{"name": "待支付", "value": "toPay"},
				{"name": "待接单", "value": "toAccept"},
				{"name": "出票成功", "value": "ticketed"},
				{"name": "已拒单", "value": "rejected"},
			},
			"item": itemOptions,
		},
	})

}
