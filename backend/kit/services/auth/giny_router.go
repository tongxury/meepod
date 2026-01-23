package auth

import (
	"fmt"
	"gitee.com/meepo/backend/kit/components/sdk/helper/phonenum"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	"github.com/gin-gonic/gin"
	"strings"
)

func CodeRoute(sign, templateCode, testPrefix string) func(c *gin.Context) {
	return func(c *gin.Context) {
		var params struct {
			Phone    string `form:"phone" binding:"required"`
			Category string
		}

		if err := c.ShouldBindQuery(&params); err != nil {
			gind.BadRequest(c, err)
			return
		}

		if !strings.HasPrefix(params.Phone, testPrefix) {

			if !phonenum.IsValid(params.Phone) {
				gind.BadRequestM(c, "手机号格式不正确")
				return
			}
		}

		ctx := c.Request.Context()

		code, isTest, err := new(Service).SendCode(ctx, params.Phone, params.Category, sign, templateCode, testPrefix)
		if err != nil {
			slf.WithError(err).Errorw("SendCode err", slf.Reflect("params", params))
			gind.Error(c, err)
			return
		}

		msg := "验证码已发送"
		if isTest {
			msg = fmt.Sprintf("%s", code)
		}

		gind.OKM(c, msg)
	}
}

func TokenRoute(checkUserFunc CheckUserFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		var params struct {
			Phone string `form:"phone" binding:"required"`
			Code  string `form:"code" binding:"required"`
		}

		if err := c.ShouldBindQuery(&params); err != nil {
			gind.BadRequest(c, err)
			return
		}

		ctx := c.Request.Context()

		token, err := new(Service).GenerateTokenByCode(ctx, params.Phone, params.Code, checkUserFunc)
		if err != nil {
			slf.WithError(err).Errorw("GenerateTokenByCode err", slf.Reflect("params", params))
			gind.Error(c, err)
			return
		}

		gind.OK(c, token)
	}
}

func TokenByPasswordRoute(checkUserFunc CheckUserPasswordFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		var params struct {
			Phone    string `form:"phone" binding:"required"`
			Password string `form:"password" binding:"required"`
		}

		if err := c.ShouldBindQuery(&params); err != nil {
			c.AbortWithStatusJSON(200, gin.H{"code": 10400, "error": err.Error()})
			return
		}

		ctx := c.Request.Context()

		token, err := new(Service).GenerateTokenByPassword(ctx, params.Phone, params.Password, checkUserFunc)
		if err != nil {
			slf.WithError(err).Errorw("GenerateTokenByPassword err", slf.Reflect("params", params))
			c.AbortWithStatusJSON(200, gin.H{"code": 10500, "error": err.Error()})
			return
		}

		c.JSONP(200, gin.H{"code": 0, "data": token})
	}
}

func CheckTokenRoute() func(c *gin.Context) {
	return func(c *gin.Context) {
		var params struct {
			Token string `form:"token" binding:"required"`
		}
		if err := c.ShouldBindQuery(&params); err != nil {
			c.AbortWithStatusJSON(200, gin.H{"code": 10400, "error": err.Error()})
			return
		}

		ctx := c.Request.Context()
		userId, err := new(Service).FindUserIdByToken(ctx, params.Token)
		if err != nil {
			slf.WithError(err).Errorw("FindUserIdByToken err", slf.Reflect("params", params))
			c.AbortWithStatusJSON(200, gin.H{"code": 10500, "error": err.Error()})
			return
		}

		if userId == "" {
			c.AbortWithStatusJSON(200, gin.H{"code": 10400, "error": "invalid token"})
			return
		}

		c.JSONP(200, gin.H{"code": 0})
	}
}
