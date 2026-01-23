package gind

import (
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"github.com/gin-gonic/gin"
)

func GetIp(c *gin.Context) string {
	return helper.OrString(c.GetHeader("X-Real-IP"), c.GetHeader("X-Forwarded-For"), c.ClientIP())
}
