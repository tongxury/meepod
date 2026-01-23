package rest

import (
	"io"
	"net/url"
	"strings"

	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	"gitee.com/meepo/backend/shop/app/payment/service"
	"gitee.com/meepo/backend/shop/core/enum"
	"github.com/gin-gonic/gin"
)

func DebugConfirmTopup(c *gin.Context) {
	ctx := c.Request.Context()

	storeId := c.Query("storeId")
	topupId := c.Query("topupId")
	orderId := c.Query("orderId")
	userId := c.Query("userId")
	amount := conv.Float64(c.Query("amount"))

	confirmed, err := new(service.TopupService).ConfirmBuyingTopup(ctx, storeId, topupId, userId, orderId, amount)
	if err != nil {
		gind.Error(c, err)
		return
	}

	gind.OK(c, confirmed)
}

func ConfirmXinshTopup(c *gin.Context) {

	slf.Debugw("ConfirmXinshTopup", slf.Reflect("req", c.Request))

	gind.OK(c)
}

func ConfirmAlipayTopup(c *gin.Context) {

	params, raw, err := func(c *gin.Context) (url.Values, string, error) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			return nil, "", err
		}

		path := string(body)

		path, _ = url.QueryUnescape(path)

		params := url.Values{}
		for _, x := range strings.Split(path, "&") {
			kv := strings.Split(x, "=")
			params.Add(kv[0], kv[1])
		}

		return params, string(body), nil
	}(c)

	if err != nil {
		slf.WithError(err).Errorw("ParseBody err", slf.String("body", raw))
		gind.BadRequest(c, err)
		return
	}

	// Test data example (sensitive data removed)

	err = comp.SDK().Alipay().CheckSign(params, params.Get("sign"))
	if err != nil {
		slf.WithError(err).Errorw("[MAIN] validSign err", slf.Reflect("params", params), slf.String("raw", raw))
		gind.Error(c, err)
		return
	}

	var passbackParams map[string]any
	err = conv.J2M(params.Get("passback_params"), &passbackParams)
	if err != nil {
		gind.BadRequest(c, err)
		return
	}

	userId := conv.String(passbackParams["userId"])
	storeId := conv.String(passbackParams["storeId"])
	orderId := conv.String(passbackParams["orderId"])
	category := conv.String(passbackParams["category"])
	topupId := conv.String(passbackParams["topupId"])
	amount := conv.Float64(passbackParams["amount"])
	// Response fields (details removed for security)

	ctx := c.Request.Context()

	var confirmed bool

	options := map[string]string{
		"trade_no":       params.Get("trade_no"),
		"seller_id":      params.Get("seller_id"),
		"buyer_id":       params.Get("buyer_id"),
		"buyer_logon_id": params.Get("buyer_logon_id"),
	}

	switch category {
	case enum.TopupCategory_Wallet.Value:
		confirmed, err = new(service.TopupService).ConfirmWalletTopup(ctx, storeId, topupId, userId, amount, options)
	case enum.TopupCategory_Buying.Value:
		confirmed, err = new(service.TopupService).ConfirmBuyingTopup(ctx, storeId, topupId, userId, orderId, amount, options)
	}

	if err != nil {
		slf.WithError(err).Errorw("[MAIN] Confirm err", slf.Reflect("raw", raw), slf.Reflect("params", passbackParams))
		gind.Error(c, err)
		return
	}

	if confirmed {
		slf.Infow("alipay confirmed ", slf.Reflect("raw", raw), slf.Reflect("passbackParams", passbackParams))
	}

	gind.OK(c)
}

func GetTopup(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		gind.BadRequestf(c, "id is required")
		return
	}
	ctx := c.Request.Context()

	//userId := c.GetString("UserId")
	//storeId := c.GetString("StoreId")

	newOrder, err := new(service.TopupService).FindById(ctx, id)
	if err != nil {
		slf.WithError(err).Errorw("RequireByStoreIdAndId err")
		gind.Error(c, err)
		return
	}

	gind.OK(c, newOrder)
}

func UpdateTopup(c *gin.Context) {

	action := c.Query("action")
	orderId := c.Param("id")
	if orderId == "" {
		gind.BadRequestf(c, "id is required")
		return
	}
	ctx := c.Request.Context()

	userId := c.GetString("UserId")
	storeId := c.GetString("StoreId")

	var err error
	switch action {
	case "cancel":
		err = new(service.TopupService).Cancel(ctx, userId, storeId, orderId)
	}

	if err != nil {
		slf.WithError(err).Errorw("Update err")
		gind.Error(c, err)
		return
	}

	newOrder, err := new(service.TopupService).FindById(ctx, orderId)
	if err != nil {
		slf.WithError(err).Errorw("RequireByStoreIdAndId err")
		gind.Error(c, err)
		return
	}

	gind.OK(c, newOrder)
}

func ListTopups(c *gin.Context) {

	storeId := c.GetString("StoreId")
	userId := c.GetString("UserId")

	page := conv.Int64(c.DefaultQuery("page", "1"))
	size := conv.Int64(c.DefaultQuery("size", "15"))

	ctx := c.Request.Context()

	orders, total, err := new(service.TopupService).ListTopupOrders(ctx, userId, storeId, page, size)

	if err != nil {
		slf.WithError(err).Errorw("ListTopupOrders err")
		gind.Error(c, err)
		return
	}

	gind.Page(c, orders, page, size, total)
}

func AddTopup(c *gin.Context) {

	ctx := c.Request.Context()

	storeId := c.GetString("StoreId")
	userId := c.GetString("UserId")

	var params struct {
		Amount float64 `binding:"required"`
		//Method string
	}

	if err := c.ShouldBindJSON(&params); err != nil {
		gind.BadRequest(c, err)
		return
	}

	//if params.Method == "" {
	//	params.Method = enum.PayMethod_Alipay.Value
	//}

	ip := gind.GetIp(c)

	data, err := new(service.TopupService).AddWalletTopUp(ctx, storeId, userId, params.Amount, ip)

	if err != nil {
		slf.WithError(err).Errorw("AddTopUp add err", slf.UserId(userId))
		gind.Error(c, err)
		return
	}

	gind.OK(c, data)
}
