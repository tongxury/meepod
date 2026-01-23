package rest

import (
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	"gitee.com/meepo/backend/shop/app/shop/service"
	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {

	userId := c.GetString("UserId")

	field := c.Query("field")
	value := c.Query("value")

	id := c.Param("id")
	if id != "me" {
		gind.BadRequestf(c, "no permission")
		return
	}

	ctx := c.Request.Context()

	err := new(service.UserService).UpdateUser(ctx, userId, field, value)
	if err != nil {
		slf.WithError(err).Errorw("UpdateUser err")
		gind.Error(c, err)
		return
	}

	newUser, err := new(service.UserService).RequireById(ctx, userId)
	if err != nil {
		slf.WithError(err).Errorw("RequireById err")
		gind.Error(c, err)
		return
	}
	gind.OK(c, newUser)
}

func GetUser(c *gin.Context) {
	//storeId := c.GetString("StoreId")
	caller := c.GetString("UserId")

	userId := c.Param("id")
	if userId == "" {
		gind.BadRequestf(c, "id is required")
		return
	}

	if userId != "me" {
		gind.BadRequestf(c, "no permission")

	}

	userId = caller

	ctx := c.Request.Context()

	user, err := new(service.UserService).RequireById(ctx, userId)
	if err != nil {
		slf.WithError(err).Errorw("RequireById err")
		gind.Error(c, err)
		return
	}

	gind.OK(c, user)
}

func GetUserProfile(c *gin.Context) {

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

	profile, err := new(service.UserService).GetUserProfile(ctx, userId, storeId)
	if err != nil {
		slf.WithError(err).Errorw("GetUserProfile err")
		gind.Error(c, err)
		return
	}

	gind.OK(c, profile)

}
