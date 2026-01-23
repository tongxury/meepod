package adminrest

import (
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	"gitee.com/meepo/backend/shop/app/payment/service"
	"github.com/gin-gonic/gin"
)

func AddPaymentStore(c *gin.Context) {

	var params struct {
		//AppId   string `binding:"required" json:"app_id"`
		StoreId string `binding:"required" json:"store_id"`
	}

	if err := c.ShouldBindJSON(&params); err != nil {
		gind.BadRequest(c, err)
		return
	}

	ctx := c.Request.Context()

	_, err := new(service.StoreService).AddStoreV2(ctx, params.StoreId)
	if err != nil {
		slf.WithError(err).Errorw("AddStore err")
		gind.Error(c, err)
		return
	}

	gind.OK(c)
}

func ListPaymentStores(c *gin.Context) {

	storeId := c.Query("store")

	page := conv.Int64(c.DefaultQuery("page", "1"))
	size := conv.Int64(c.DefaultQuery("size", "15"))

	ctx := c.Request.Context()

	stores, total, err := new(service.StoreService).ListStores(ctx, storeId, page, size)
	if err != nil {
		gind.Error(c, err)
		slf.WithError(err).Errorw("ListStores err")
		return
	}

	gind.Page(c, stores, page, size, total)
}

func UpdatePaymentStore(c *gin.Context) {

	storeId := c.Param("storeId")

	if storeId == "" {
		gind.BadRequestf(c, "storeId is required")
		return
	}

	action := c.Query("action")

	ctx := c.Request.Context()

	var err error

	switch action {
	case "applyXinsh":
		err = new(service.StoreService).ApplyXinsh(ctx, storeId)

	}

	if err != nil {
		gind.Error(c, err)
		slf.WithError(err).Errorw("Update Store err")
		return
	}

	gind.OK(c)
}
