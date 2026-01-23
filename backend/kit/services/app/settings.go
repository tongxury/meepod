package app

import (
	"context"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	"github.com/gin-gonic/gin"
)

type GetSettingsFunc func(ctx context.Context) (any, error)

func SettingsRoute(f GetSettingsFunc) func(c *gin.Context) {
	return func(c *gin.Context) {

		ctx := c.Request.Context()
		settings, err := f(ctx)
		if err != nil {
			slf.WithError(err).Errorw("get settings err")
			gind.Error(c, err)
			return
		}

		gind.OK(c, settings)
	}
}
