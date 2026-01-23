package base

import (
	"fmt"
	"strings"
	"time"

	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	"gitee.com/meepo/backend/kit/services/util/gind/errorx"
	coredb "gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var URILock = func(c *gin.Context) {

	ctx := c.Request.Context()

	key := c.Request.RequestURI

	set, err := comp.SDK().Redis().SetNX(ctx, key, "", 5*time.Second).Result()
	if err != nil {
		gind.Error(c, err)
		return
	}

	if !set {
		gind.BadRequestf(c, "too frequently")
		return
	}

	c.Next()
}

var UserLock = func(c *gin.Context) {
	userId := c.GetString("UserId")

	//if userId == "" {
	//	gind.BadRequestf(c, "userId is not found")
	//	return
	//}

	ctx := c.Request.Context()

	key := c.Request.Method + "_" + c.Request.RequestURI + "_" + userId

	set, err := comp.SDK().Redis().SetNX(ctx, key, "", 2*time.Second).Result()
	if err != nil {
		gind.Error(c, err)
		return
	}

	if !set {
		gind.BadRequest(c, errorx.UserMessage("操作频繁，请稍后"))
		return
	}

	c.Next()
}

var RequireStoreId = func(c *gin.Context) {
	val := c.GetHeader("StoreId")
	if val == "" {
		gind.AbortWithCodef(c, 10404, "StoreId is required in header")
		return
	}
	c.Set("StoreId", val)
	c.Next()
}

var StoreId = func(c *gin.Context) {
	val := c.GetHeader("StoreId")
	c.Set("StoreId", val)
	c.Next()
}

var RequiredConfirmedStoreId = func(c *gin.Context) {
	val := c.GetHeader("StoreId")
	if val == "" {
		gind.AbortWithCodef(c, 10404, "StoreId is required in header")
		return
	}

	ctx := c.Request.Context()

	store, err := new(coredb.Store).FindByIdAndStatus(ctx, val, enum.StoreStatus_Confirmed.Value)
	if err != nil {
		slf.WithError(err).Errorw("FindByIdAndStatus err")
		gind.Error(c, err)
		return
	}

	if store == nil {
		gind.AbortWithCodef(c, 10400, "store is invalid")
		return
	}

	c.Set("StoreId", val)
	c.Next()
}

var RequireAuthToken = func(c *gin.Context) {

	authToken := c.GetHeader("Authorization")
	if authToken == "" {
		gind.AbortWithCodef(c, 10401, "token is required")
		return
	}

	authToken = strings.Split(authToken, " ")[1]

	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return enum.SecretSignKey, nil
	})

	if err != nil {
		slf.WithError(err).Errorw("jwt.Parse err")
		gind.AbortWithCodef(c, 10401, "token is required")
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		c.Set("UserId", claims["user_id"])
	} else {
		slf.WithError(err).Errorw("token invalid err")
		gind.AbortWithCodef(c, 10401, "token is required")
		return
	}

	c.Next()
}

var StoreOwner = func(c *gin.Context) {

	storeId := c.Param("storeId")
	if storeId != "me" {
		gind.AbortWithCodef(c, 10404, "没有权限")
		return
	}

	userId := c.GetString("UserId")

	ctx := c.Request.Context()

	store, err := new(coredb.Store).FindByOwnerId(ctx, userId)
	if err != nil {
		slf.WithError(err).Errorw("FindStore err")
		gind.Error(c, err)
		return
	}

	if store == nil {
		gind.AbortWithCodef(c, 10404, "未开通店铺")
		return
	}

	c.Set("StoreId", store.Id)
	c.Next()
}
