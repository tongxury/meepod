package rest

import (
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	keeperservice "gitee.com/meepo/backend/shop/app/shop/service/keeper"
	"github.com/gin-gonic/gin"
)

func ListStores(c *gin.Context) {

	keyword := c.Query("keyword")
	id := c.Query("id")

	keeperId := c.GetString("UserId")

	ctx := c.Request.Context()

	stores, err := new(keeperservice.StoreService).ListStores(ctx, keeperId, id, keyword)
	if err != nil {
		slf.WithError(err).Errorw("ListStores err")
		gind.Error(c, err)
		return
	}
	gind.OK(c, stores)
}

func UpdateStore(c *gin.Context) {

	storeId := c.GetString("StoreId")
	keeperId := c.GetString("UserId")

	category := c.Query("category")

	ctx := c.Request.Context()

	var err error
	switch category {
	case "notice":
		var params struct {
			Notice string
		}

		if err := c.ShouldBindJSON(&params); err != nil {
			gind.BadRequest(c, err)
			return
		}

		err = new(keeperservice.StoreService).UpdateNotice(ctx, keeperId, storeId, params.Notice)

	case "items":
		var params struct {
			ItemIds []string
		}

		if err := c.ShouldBindJSON(&params); err != nil {
			gind.BadRequest(c, err)
			return
		}

		err = new(keeperservice.StoreService).UpdateItemSettings(ctx, keeperId, storeId, params.ItemIds)

	default:

		field := c.Query("field")
		value := c.Query("value")
		err = new(keeperservice.StoreService).Update(ctx, storeId, field, value)
	}

	if err != nil {
		slf.WithError(err).Errorw("Update err")
		gind.Error(c, err)
		return
	}

	store, err := new(keeperservice.StoreService).FindStoreById(ctx, storeId)
	if err != nil {
		slf.WithError(err).Errorw("GetStoreProfile err")
		gind.Error(c, err)
		return
	}
	gind.OK(c, store)
}

func GetStore(c *gin.Context) {

	storeId := c.GetString("StoreId")

	ctx := c.Request.Context()

	user, err := new(keeperservice.StoreService).FindStoreById(ctx, storeId)
	if err != nil {
		slf.WithError(err).Errorw("FindStoreById err")
		gind.Error(c, err)
		return
	}

	gind.OK(c, user)
}

//func GetStoreProfile(c *gin.Context) {
//
//	ctx := c.Request.Context()
//
//	storeId := c.GetString("StoreId")
//
//	store, err := new(keeperservice.StoreService).GetStoreProfile(ctx, storeId)
//	if err != nil {
//		slf.WithError(err).Errorw("FindStore err")
//		gind.Error(c, err)
//		return
//	}
//
//	gind.OK(c, store)
//}
