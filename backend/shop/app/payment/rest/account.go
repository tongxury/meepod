package rest

import (
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	"gitee.com/meepo/backend/shop/app/payment/service"
	"github.com/gin-gonic/gin"
)

func GetAccount(c *gin.Context) {

	storeId := c.GetString("StoreId")
	caller := c.GetString("UserId")

	userId := c.Param("id")
	if userId == "" {
		gind.BadRequestf(c, "id is required")
		return
	}

	if userId == "me" {
		userId = caller
	}

	ctx := c.Request.Context()

	profile, err := new(service.AccountService).RequireByUserIdAndStoreId(ctx, userId, storeId)
	if err != nil {
		slf.WithError(err).Errorw("RequireByIdAndStoreId err")
		gind.Error(c, err)
		return
	}

	gind.OK(c, profile)

}
