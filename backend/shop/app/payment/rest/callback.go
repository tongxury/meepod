package rest

import (
	"fmt"

	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	"gitee.com/meepo/backend/shop/app/payment/service"
	"github.com/gin-gonic/gin"
)

func Callback(c *gin.Context) {

	appAuthCode := c.Query("app_auth_code")
	//appId := c.Query("app_id")
	source := c.Query("source") // alipay_app_auth

	ctx := c.Request.Context()

	// Callback URL example (domain removed)

	var err error
	switch source {
	case "alipay_app_auth":
		err = new(service.AlipayService).SaveAuthToken(ctx, appAuthCode)
	default:
		err = fmt.Errorf("unknown source: %s", source)
	}

	if err != nil {
		slf.WithError(err).Errorw("Callback err")
		gind.Error(c, err)
		return
	}

	gind.OK(c)
}
