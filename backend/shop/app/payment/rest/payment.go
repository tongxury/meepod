package rest

import (
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	"gitee.com/meepo/backend/shop/app/payment/service"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/types"
	"github.com/gin-gonic/gin"
)

func ListPayMethods(c *gin.Context) {

	storeId := c.GetString("StoreId")
	userId := c.GetString("UserId")

	amount := conv.Float64(c.Query("amount"))
	ctx := c.Request.Context()

	methods, err := new(service.PaymentService).ListPayMethods(ctx, storeId, userId, amount)
	if err != nil {
		slf.WithError(err).Errorw("ListPayMethods err")
		gind.Error(c, err)
		return
	}

	gind.OK(c, methods)
}

func ListPayments(c *gin.Context) {

	storeId := c.GetString("StoreId")
	userId := c.GetString("UserId")

	page := conv.Int64(c.DefaultQuery("page", "1"))
	size := conv.Int64(c.DefaultQuery("size", "15"))

	ctx := c.Request.Context()

	orders, total, err := new(service.PaymentService).ListPayments(ctx, userId, storeId, page, size)

	if err != nil {
		slf.WithError(err).Errorw("ListPayments err")
		gind.Error(c, err)
		return
	}

	gind.Page(c, orders, page, size, total)
}

//func AddPayment(c *gin.Context) {
//
//	var params types.PaymentForm
//
//	if err := c.ShouldBindJSON(&params); err != nil {
//		gind.BadRequest(c, err)
//		return
//	}
//
//	if params.Category == "" {
//		params.Category = enum.BizCategory_Order.Value
//	}
//
//	method := c.Query("method")
//	storeId := c.GetString("StoreId")
//
//	ctx := c.Request.Context()
//
//	userId := c.GetString("UserId")
//
//	ip := gind.GetIp(c)
//
//	var data any
//	var err error
//
//	switch method {
//	case enum.PayMethod_Alipay.Value, enum.PayMethod_Wechat.Value:
//		data, err = new(service.PaymentService).PayByQrCode(ctx, storeId, userId, params.OrderId, params.Category, ip)
//	default:
//		_, err = new(service.PaymentService).PayByAccount(ctx, storeId, userId, params.OrderId, params.Category)
//	}
//
//	if err != nil {
//		slf.WithError(err).Errorw("Pay err", slf.Reflect("params", params))
//		gind.Error(c, err)
//		return
//	}
//
//	gind.OK(c, data)
//
//}

func AddPaymentV2(c *gin.Context) {

	var params types.PaymentForm

	if err := c.ShouldBindJSON(&params); err != nil {
		gind.BadRequest(c, err)
		return
	}

	if params.Category == "" {
		params.Category = enum.BizCategory_Order.Value
	}

	ctx := c.Request.Context()
	storeId := c.GetString("StoreId")
	userId := c.GetString("UserId")

	ip := gind.GetIp(c)

	payed, err := new(service.PaymentService).PayByAccount(ctx, storeId, userId, params.OrderId, params.Category)
	if err != nil {
		slf.WithError(err).Errorw("PayByAccount err", slf.Reflect("params", params))
		gind.Error(c, err)
		return
	}

	if payed {
		gind.OK(c)
		return
	}

	data, err := new(service.PaymentService).PayByQrCode(ctx, storeId, userId, params.OrderId, params.Category, ip)

	if err != nil {
		slf.WithError(err).Errorw("Pay err", slf.Reflect("params", params))
		gind.Error(c, err)
		return
	}

	gind.OK(c, data)

}
