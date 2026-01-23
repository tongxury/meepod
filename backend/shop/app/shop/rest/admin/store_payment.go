package adminrest

import (
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	adminserivce "gitee.com/meepo/backend/shop/app/shop/service/admin"
	"github.com/gin-gonic/gin"
)

//func UpdateStore(c *gin.Context) {
//
//	action := c.Query("action")
//	ids := c.Query("ids")
//
//	idList := strings.Split(ids, ",")
//	if len(idList) == 0 {
//		gind.BadRequestf(c, "ids is required")
//		return
//	}
//
//	ctx := c.Request.Context()
//
//	var err error
//	switch action {
//	case "updateMember":
//		var params struct {
//			Level int64
//			Until int64
//		}
//
//		if err := c.ShouldBindJSON(&params); err != nil {
//			gind.BadRequest(c, err)
//			return
//		}
//
//		err = new(service.StoreService).UpdateMember(ctx, idList[0], params.Level, params.Until)
//
//	case "update":
//		var params types.StoreParams
//
//		if err := c.ShouldBindJSON(&params); err != nil {
//			gind.BadRequest(c, err)
//			return
//		}
//
//		err = new(service.StoreService).UpdateStore(ctx, idList[0], params)
//	case "delete":
//		err = new(service.StoreService).UpdateStatus(ctx, idList, enum.StoreStatus_Deleted.Value)
//	case "confirm":
//		err = new(service.StoreService).UpdateStatus(ctx, idList, enum.StoreStatus_Confirmed.Value)
//	}
//
//	if err != nil {
//		slf.WithError(err).Errorw("UpdateStatus err")
//		gind.Error(c, err)
//		return
//	}
//
//	gind.OK(c)
//}

func AddStorePayment(c *gin.Context) {

	storeId := c.Param("storeId")

	var params struct {
		Amount float64
	}

	if err := c.ShouldBindJSON(&params); err != nil {
		gind.BadRequest(c, err)
		return
	}

	ctx := c.Request.Context()

	err := new(adminserivce.StorePaymentService).Topup(ctx, storeId, params.Amount)
	if err != nil {
		slf.WithError(err).Errorw("Topup err")
		gind.Error(c, err)
		return
	}

	gind.OK(c)
}

//func ListStores(c *gin.Context) {
//
//	id := c.Query("id")
//	name := c.Query("name")
//	phone := c.Query("owner")
//
//	page := conv.Int64(c.DefaultQuery("page", "1"))
//	size := conv.Int64(c.DefaultQuery("size", "15"))
//
//	ctx := c.Request.Context()
//
//	stores, total, err := new(service.StoreService).ListStores(ctx, id, name, phone, page, size)
//	if err != nil {
//		gind.Error(c, err)
//		slf.WithError(err).Errorw("ListStores err")
//		return
//	}
//
//	gind.Page(c, stores, page, size, total)
//}
