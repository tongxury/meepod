package main

import (
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/shop/app/base"
	"gitee.com/meepo/backend/shop/app/payment/rest"
	keeperrest "gitee.com/meepo/backend/shop/app/payment/rest/keeper"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	args := comp.Flags().Logger().Redis().Postgres().Server().AliOSS().Alipay().Xinsh().
		CustomStr("domain", "", "").
		Parse()

	comp.SDK().Preparing().Logger(args.Log).Redis(args.RepoRedis).Postgres(args.RepoPostgres).AliOSS(args.AliOSSOptions).
		Alipay(args.AlipayOptions).Xinsh(args.XinshConfig)

	engine := gin.New()

	ginCors := cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})

	engine.Use(ginCors)

	lc := base.UserLock
	token := base.RequireAuthToken
	storeId := base.RequireStoreId
	owner := base.StoreOwner

	root := engine.Group("/api/payment")
	{
		//keeper
		root.GET("/v1/stores/:storeId/accounts", token, owner, keeperrest.ListAccounts)
		root.GET("/v1/stores/:storeId/account-summaries", token, owner, keeperrest.GetSummary)
		root.PUT("/v1/stores/:storeId/accounts/:accountId", token, lc, owner, keeperrest.Update)

		root.GET("/v1/stores/:storeId/withdraws", token, owner, keeperrest.ListWithdraws)
		root.PUT("/v1/stores/:storeId/withdraws/:id", token, lc, owner, keeperrest.UpdateWithdraw)

		root.POST("/v1/stores/:storeId/co-store-payments", token, lc, owner, keeperrest.AddCoStoreTopUp)
		root.GET("/v1/stores/:storeId/co-store-payments", token, owner, keeperrest.ListCoStorePayments)

		root.GET("/v1/stores/:storeId/store-payments", token, owner, keeperrest.ListPayments)

	}

	{
		//curl http://localhost:6066/payment/v1/callback?app_auth_code=P8cd71fd6ff55403fb7c0bfc2cdf9047&app_id=2021004101679599&source=alipay_app_auth

		root.GET("/v1/callback", lc, rest.Callback)
		//root.GET("/v1/callback-repair", lc, rest.Callback)
		root.POST("/v1/topup-callback", rest.ConfirmAlipayTopup)
		root.Any("/v1/xinsh-topup-callback", rest.ConfirmXinshTopup)

		root.GET("/v1/accounts/:id", token, storeId, rest.GetAccount)

		root.POST("/v1/topups", token, lc, storeId, rest.AddTopup)
		root.GET("/v1/topups", token, storeId, rest.ListTopups)
		root.PUT("/v1/topups/:id", token, storeId, lc, rest.UpdateTopup)
		root.GET("/v1/topups/:id", token, storeId, rest.GetTopup)

		root.GET("/v1/payment-methods", token, storeId, rest.ListPayMethods)
		root.POST("/v1/payments", token, storeId, lc, rest.AddPaymentV2)
		root.GET("/v1/payments", token, storeId, rest.ListPayments)

		root.POST("/v1/withdraws", token, storeId, lc, rest.AddWithdraw)
		root.GET("/v1/withdraws", token, storeId, rest.ListWithdraws)
		root.PUT("/v1/withdraws/:id", token, storeId, rest.UpdateWithdraw)

	}

	err := engine.Run("0.0.0.0:" + args.ServerOptions.Port)
	if err != nil {
		log.Fatal(err)
	}

}
