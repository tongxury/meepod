package gind

import "github.com/gin-gonic/gin"

func RequireHeader(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		val := c.GetHeader(name)
		if val == "" {
			AbortWithCodef(c, 10400, name+" is required in header")
			return
		}
		c.Set(name, val)
		c.Next()
	}
}

func SetFromHeader(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(name, c.GetHeader(name))
		c.Next()
	}
}
