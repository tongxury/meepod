package keeperrest

import (
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	keeperservice "gitee.com/meepo/backend/shop/app/payment/service/keeper"
	"github.com/gin-gonic/gin"
)

func ListCoStorePayments(c *gin.Context) {

	month := c.Query("month")
	cat := c.Query("cat")

	//storeId := c.GetString("StoreId")
	userId := c.GetString("UserId")

	storeId := c.Query("storeId")
	coStoreId := c.Query("coStoreId")

	page := conv.Int64(c.DefaultQuery("page", "1"))
	size := conv.Int64(c.DefaultQuery("size", "15"))

	ctx := c.Request.Context()

	orders, total, err := new(keeperservice.CoStorePaymentService).ListPayments(ctx, month, cat, userId, storeId, coStoreId, page, size)

	if err != nil {
		slf.WithError(err).Errorw("ListPayments err")
		gind.Error(c, err)
		return
	}

	gind.Page(c, orders, page, size, total)
}

func AddCoStoreTopUp(c *gin.Context) {

	ctx := c.Request.Context()

	storeId := c.GetString("StoreId")
	userId := c.GetString("UserId")

	var params struct {
		StoreId string
		Amount  float64
	}

	if err := c.ShouldBindJSON(&params); err != nil {
		gind.BadRequest(c, err)
		return
	}

	err := new(keeperservice.CoStorePaymentService).Topup(ctx, storeId, params.StoreId, params.Amount)

	if err != nil {
		slf.WithError(err).Errorw("Topup add err", slf.UserId(userId))
		gind.Error(c, err)
		return
	}

	gind.OK(c)
}
