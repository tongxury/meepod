package adminrest

import (
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	adminserivce "gitee.com/meepo/backend/shop/app/shop/service/admin"
	"github.com/gin-gonic/gin"
)

func GetStats(c *gin.Context) {

	ctx := c.Request.Context()

	stats, err := new(adminserivce.StatsService).GetStats(ctx)
	if err != nil {
		slf.WithError(err).Errorw("GetStats err")
		gind.Error(c, err)
		return
	}

	gind.OK(c, stats)

}
