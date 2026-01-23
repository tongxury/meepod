package rest

import (
	"gitee.com/meepo/backend/kit/components/sdk/conv"
	"gitee.com/meepo/backend/kit/components/sdk/slf"
	"gitee.com/meepo/backend/kit/services/util/gind"
	"gitee.com/meepo/backend/shop/app/shop/service"
	"gitee.com/meepo/backend/shop/core/types"
	"github.com/gin-gonic/gin"
)

func ListPlans(c *gin.Context) {

	category := c.Query("category")

	storeId := c.GetString("StoreId")
	userId := c.GetString("UserId")
	page := conv.Int64(c.DefaultQuery("page", "1"))
	size := conv.Int64(c.DefaultQuery("size", "15"))

	ctx := c.Request.Context()

	var plans types.Plans
	var total int64
	var err error

	switch category {
	case "saved":
		plans, total, err = new(service.PlanService).ListSavedPlans(ctx, userId, storeId, page, size)
	}

	if err != nil {
		slf.WithError(err).Errorw("ListPlans err")
		gind.Error(c, err)
		return
	}

	gind.Page(c, plans, page, size, total)

}

func DeletePlan(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		gind.BadRequestf(c, "id is required")
		return
	}

	ctx := c.Request.Context()
	userId := c.GetString("UserId")

	err := new(service.PlanService).Delete(ctx, userId, id)
	if err != nil {
		slf.WithError(err).Errorw("Delete err", slf.String("id", id))
		gind.Error(c, err)
		return
	}

	gind.OK(c)
}

func AddPlan(c *gin.Context) {

	var params types.PlanForm
	if err := c.ShouldBindJSON(&params); err != nil {
		gind.BadRequest(c, err)
		return
	}

	ctx := c.Request.Context()

	userId := c.GetString("UserId")
	storeId := c.GetString("StoreId")

	planId, err := new(service.PlanService).AddPlan(ctx, userId, storeId, params)
	if err != nil {
		slf.WithError(err).Errorw("Add err", slf.Reflect("plan", params))
		gind.Error(c, err)
		return
	}

	gind.OK(c, planId)
}
