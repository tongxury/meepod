package rest

import (
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	keeperservice "gitee.com/meepo/backend/shop/app/shop/service/keeper"
	"gitee.com/meepo/backend/shop/core/types"
	"github.com/gin-gonic/gin"
)

func UpdateCoStore(c *gin.Context) {
	storeId := c.GetString("StoreId")
	keeperId := c.GetString("UserId")

	action := c.Query("action")

	coStoreId := c.Param("coStoreId")

	ctx := c.Request.Context()

	var err error
	switch action {
	case "updateItems":
		var params struct {
			Items map[string]float64 `binding:"required"`
		}

		if err := c.ShouldBindJSON(&params); err != nil {
			gind.BadRequest(c, err)
			return
		}

		err = new(keeperservice.CoStoreService).UpdateItems(ctx, keeperId, storeId, coStoreId, params.Items)
	case "pause":
		err = new(keeperservice.CoStoreService).Pause(ctx, keeperId, storeId, coStoreId)
	case "resume":
		err = new(keeperservice.CoStoreService).Resume(ctx, keeperId, storeId, coStoreId)
	case "endApply":
		err = new(keeperservice.CoStoreService).ApplyForEnd(ctx, keeperId, storeId, coStoreId)
	case "end":

		imageProof := c.Query("imageProof")
		if imageProof == "" {
			gind.BadRequestf(c, "imageProof is required")
			return
		}

		err = new(keeperservice.CoStoreService).End(ctx, keeperId, storeId, coStoreId, imageProof)
	case "recover":
		err = new(keeperservice.CoStoreService).Recover(ctx, keeperId, storeId, coStoreId)
	default:
	}

	if err != nil {
		slf.WithError(err).Errorw("Update err")
		gind.Error(c, err)
		return
	}

	//item, err := new(keeperservice.CoStoreService).FindByCoStoreId(ctx, storeId, coStoreId)
	//if err != nil {
	//	slf.WithError(err).Errorw("FindById err")
	//	gind.Error(c, err)
	//	return
	//}
	gind.OKM(c, "操作成功")
}

func AddCoStore(c *gin.Context) {

	var params struct {
		CoStoreId string
		ItemIds   []string
	}

	if err := c.ShouldBindJSON(&params); err != nil {
		gind.BadRequest(c, err)
		return
	}

	keeperId := c.GetString("UserId")
	storeId := c.GetString("StoreId")

	ctx := c.Request.Context()

	err := new(keeperservice.CoStoreService).AddCoStore(ctx, keeperId, storeId, params.CoStoreId, params.ItemIds)
	if err != nil {
		slf.WithError(err).Errorw("AddProxy err")
		gind.Error(c, err)
		return
	}

	gind.OK(c)

}

func ListCoStores(c *gin.Context) {

	page := conv.Int64(c.DefaultQuery("page", "1"))
	size := conv.Int64(c.DefaultQuery("size", "15"))

	keeperId := c.GetString("UserId")
	storeId := c.GetString("StoreId")

	enableOnly := c.Query("enableOnly") == "1"
	itemId := c.Query("itemId")

	ctx := c.Request.Context()

	category := c.Query("category")

	var items types.CoStores
	var total int64
	var err error

	switch category {

	case "in":
		items, total, err = new(keeperservice.CoStoreService).ListInCoStores(ctx, enableOnly, keeperId, storeId, page, size)
	//case "out":
	default:
		items, total, err = new(keeperservice.CoStoreService).ListOutCoStores(ctx, enableOnly, itemId, keeperId, storeId, page, size)
	}

	if err != nil {
		slf.WithError(err).Errorw("ListCoStores err")
		gind.Error(c, err)
		return
	}

	gind.Page(c, items, page, size, total)
}
