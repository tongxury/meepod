package keeperrest

import (
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	keeperservice "gitee.com/meepo/backend/shop/app/payment/service/keeper"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {

	action := c.Query("action")

	storeId := c.GetString("StoreId")
	id := c.Param("accountId")
	if id == "" {
		gind.BadRequestf(c, "accountId is required")
		return
	}

	ctx := c.Request.Context()

	var err error
	switch action {
	case "topUp":

		var params struct {
			Amount float64
			Remark string
		}

		if err := c.ShouldBindJSON(&params); err != nil {
			gind.BadRequest(c, err)
			return
		}

		if params.Amount <= 0 {
			gind.BadRequest(c, err)
			return
		}

		//err = new(service.AccountService).TopUp(ctx, storeId, id, params.Amount, params.Remark)
	case "decrease":
		var params struct {
			Amount float64
			Remark string
		}

		if err := c.ShouldBindJSON(&params); err != nil {
			gind.BadRequest(c, err)
			return
		}

		if params.Amount <= 0 {
			gind.BadRequestf(c, "")
			return
		}

		err = new(keeperservice.AccountService).Decr(ctx, storeId, id, params.Amount, params.Remark)
	}

	if err != nil {
		slf.WithError(err).Errorw("Update err", slf.String("action", action))
		gind.Error(c, err)
		return
	}

	newAccount, err := new(keeperservice.AccountService).RequireByIdAndStoreId(ctx, id, storeId)
	if err != nil {
		slf.WithError(err).Errorw("RequireByIdAndStoreId err", slf.String("action", action))
		gind.Error(c, err)
		return
	}

	gind.OK(c, newAccount)
}

func GetSummary(c *gin.Context) {
	storeId := c.GetString("StoreId")
	ctx := c.Request.Context()

	summary, err := new(keeperservice.AccountService).GetStoreSummary(ctx, storeId)
	if err != nil {
		slf.WithError(err).Errorw("GetStoreSummary err")
		gind.Error(c, err)
		return
	}

	gind.OK(c, summary)
}

func ListAccounts(c *gin.Context) {

	storeId := c.GetString("StoreId")
	page := conv.Int64(c.DefaultQuery("page", "1"))
	size := conv.Int64(c.DefaultQuery("size", "15"))

	ctx := c.Request.Context()

	accounts, total, err := new(keeperservice.AccountService).ListAccounts(ctx, storeId, page, size)
	if err != nil {
		slf.WithError(err).Errorw("ListAccounts err")
		gind.Error(c, err)
		return
	}

	gind.Page(c, accounts, page, size, total)
}
