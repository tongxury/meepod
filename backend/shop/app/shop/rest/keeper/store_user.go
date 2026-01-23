package rest

import (
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	keeperservice "gitee.com/meepo/backend/shop/app/shop/service/keeper"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	phone := c.Query("phone")

	keeperId := c.GetString("UserId")
	storeId := c.GetString("StoreId")

	page := conv.Int64(c.DefaultQuery("page", "1"))
	size := conv.Int64(c.DefaultQuery("size", "15"))

	if phone == "" {
		gind.OK(c)
		return
	}

	ctx := c.Request.Context()

	storeUsers, total, err := new(keeperservice.StoreUserService).ListStoreUsers(ctx, keeperId, storeId, phone, page, size)
	if err != nil {
		slf.WithError(err).Errorw("FindUserByPhone err")
		gind.Error(c, err)
		return
	}

	gind.Page(c, storeUsers, page, size, total)
}

func GetStoreUser(c *gin.Context) {

	userId := c.Param("userId")

	storeId := c.GetString("StoreId")

	//page := conv.Int64(c.DefaultQuery("page", "1"))
	//size := conv.Int64(c.DefaultQuery("size", "15"))

	if userId == "" {
		gind.BadRequestf(c, "userId is required")
		return
	}

	ctx := c.Request.Context()

	user, err := new(keeperservice.StoreUserService).FindStoreUserById(ctx, storeId, userId)
	if err != nil {
		slf.WithError(err).Errorw("FindStoreUserById err")
		gind.Error(c, err)
		return
	}
	gind.OK(c, user)
}

func UpdateStoreUser(c *gin.Context) {

	storeId := c.GetString("StoreId")

	field := c.Query("field")
	value := c.Query("value")

	userId := c.Param("userId")
	if userId == "" {
		gind.BadRequestf(c, "userId is required")
		return
	}

	ctx := c.Request.Context()

	err := new(keeperservice.StoreUserService).UpdateUser(ctx, storeId, userId, field, value)
	if err != nil {
		slf.WithError(err).Errorw("UpdateUser err")
		gind.Error(c, err)
		return
	}

	newUser, err := new(keeperservice.StoreUserService).FindStoreUserById(ctx, storeId, userId)
	if err != nil {
		slf.WithError(err).Errorw("FindStoreUserById err")
		gind.Error(c, err)
		return
	}
	gind.OK(c, newUser)
}
