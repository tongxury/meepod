package gind

import (
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"gitee.com/meepo/backend/kit/services/util/gind/errorx"
	"github.com/gin-gonic/gin"
)

func OK(c *gin.Context, data ...any) {

	if len(data) > 0 {
		c.JSONP(200, gin.H{"code": 0, "data": data[0]})
	} else {
		c.JSONP(200, gin.H{"code": 0})
	}
}
func OKM(c *gin.Context, message string, data ...any) {

	if len(data) > 0 {
		c.JSONP(200, gin.H{"code": 0, "message": message, "data": data[0]})
	} else {
		c.JSONP(200, gin.H{"code": 0, "message": message})
	}
}

func Page(c *gin.Context, data any, page, size, total int64) {

	noMore := helper.Choose(total%size == 0, page >= total/size, page >= total/size+1)

	c.JSONP(200, gin.H{"code": 0, "data": gin.H{"list": data, "no_more": noMore, "current": page, "size": size, "total": total}})
	//c.JSONP(200, gin.H{"list": data, "no_more": noMore, "current": page, "size": size, "total": total})
}

func AbortWithCode(c *gin.Context, code int, err error) {
	c.AbortWithStatusJSON(200, gin.H{"code": code, "error": err.Error()})
}

func AbortWithCodef(c *gin.Context, code int, format string, a ...any) {
	c.AbortWithStatusJSON(200, gin.H{"code": code, "error": fmt.Sprintf(format, a...)})
}

func Error(c *gin.Context, err error) {

	if errorx.IsMyErr(err) {
		e := err.(errorx.Error)
		c.AbortWithStatusJSON(200, gin.H{"code": e.Code, "error": e.Error(), "message": e.Message})
	} else {

		c.AbortWithStatusJSON(200, gin.H{"code": 10500, "message": "服务器繁忙，请稍后产重试"})
	}
}

func BadRequest(c *gin.Context, err error) {
	c.AbortWithStatusJSON(200, gin.H{"code": 10400, "error": helper.Choose(err == nil, "", err.Error()), "message": "参数错误"})
}

func BadRequestf(c *gin.Context, format string, a ...any) {
	c.AbortWithStatusJSON(200, gin.H{"code": 10400, "error": fmt.Sprintf(format, a...), "message": "参数错误"})
}

func BadRequestM(c *gin.Context, message string) {
	c.AbortWithStatusJSON(200, gin.H{"code": 10400, "message": message})
}
