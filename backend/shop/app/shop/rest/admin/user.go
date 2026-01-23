package adminrest

import (
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	"gitee.com/meepo/backend/shop/app/shop/service"
	"github.com/gin-gonic/gin"
)

func ListUsers(c *gin.Context) {

	id := c.Query("id")
	phone := c.Query("phone")

	page := conv.Int64(c.DefaultQuery("page", "1"))
	size := conv.Int64(c.DefaultQuery("size", "15"))

	ctx := c.Request.Context()

	stores, total, err := new(service.UserService).ListUsers(ctx, id, phone, "", page, size)
	if err != nil {
		gind.Error(c, err)
		slf.WithError(err).Errorw("ListUsers err")
		return
	}

	gind.Page(c, stores, page, size, total)
}
