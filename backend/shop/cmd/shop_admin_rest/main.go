package main

import (
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/services/app"
	adminrest "gitee.com/meepo/backend/shop/app/shop/rest/admin"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	args := comp.Flags().Logger().Redis().Postgres().Server().AliOSS().Xinsh().Parse()

	comp.SDK().Preparing().Logger(args.Log).Redis(args.RepoRedis).AliOSS(args.AliOSSOptions).
		Postgres(args.RepoPostgres).Loc().Xinsh(args.XinshConfig)

	engine := gin.New()

	root := engine.Group("/api/")

	{
		root.GET("/v1/stats", adminrest.GetStats)
		root.PUT("/v1/stores", adminrest.UpdateStore)
		root.POST("/v1/stores", adminrest.AddStore)
		root.GET("/v1/stores", adminrest.ListStores)
		root.POST("/v1/stores/:storeId/payments", adminrest.AddStorePayment)
		root.POST("/v1/payment-stores", adminrest.AddPaymentStore)
		root.GET("/v1/payment-stores", adminrest.ListPaymentStores)
		root.PATCH("/v1/payment-stores/:storeId", adminrest.UpdatePaymentStore)
		root.GET("/v1/users", adminrest.ListUsers)
		root.GET("/v1/orders", adminrest.ListOrders)
		root.GET("/v1/locs", adminrest.ListLocs)
		root.GET("/v1/banks", adminrest.ListBanks)
		root.POST("/v1/images", app.UploaderRoute(args.AliOSSOptions.Bucket))
	}

	err := engine.Run("0.0.0.0:" + args.ServerOptions.Port)
	if err != nil {
		log.Fatal(err)
	}

}
