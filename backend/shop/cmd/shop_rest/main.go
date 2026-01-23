package main

import (
	"log"

	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/services/app"
	"gitee.com/meepo/backend/kit/services/auth"
	"gitee.com/meepo/backend/shop/app/base"
	"gitee.com/meepo/backend/shop/app/shop/rest"
	keeperrest "gitee.com/meepo/backend/shop/app/shop/rest/keeper"
	"gitee.com/meepo/backend/shop/app/shop/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	args := comp.Flags().Logger().Redis().Postgres().AliOSS().AliSMS().UCloudSMS().Server().
		Parse()

	comp.SDK().Preparing().Logger(args.Log).Postgres(args.RepoPostgres).Redis(args.RepoRedis).
		AliOSS(args.AliOSSOptions).AliSMS(args.AliSMSOptions).UCloudSMS(args.UCloudSMSOptions)

	engine := gin.New()

	ginCors := cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})

	engine.Use(ginCors)

	token := base.RequireAuthToken
	rStoreId := base.RequireStoreId
	storeId := base.StoreId
	confirmedStoreId := base.RequiredConfirmedStoreId
	lc := base.UserLock
	uriLc := base.URILock
	owner := base.StoreOwner

	api := engine.Group("/api/")
	{
		api.POST("/v1/auth-codes", auth.CodeRoute(args.UCloudSMSOptions.Sign, args.UCloudSMSOptions.TemplateCode, args.UCloudSMSOptions.TestPrefix))
		us := new(service.UserService)
		api.GET("/v1/auth-tokens", auth.TokenRoute(us.FindOrCreateUserIdByPhone))
		api.GET("/v1/auth-status", auth.CheckTokenRoute())
		api.GET("/v1/settings", rest.GetSettings)
		//api.GET("/v1/updates", rest.GetUpdates)
		api.POST("/v2/images", token, lc, app.UploadBase64Route(args.AliOSSOptions.Bucket))
	}

	// keeper
	{
		api.GET("/v1/stores/:storeId/order-filters", token, owner, keeperrest.ListOrderFilters)
		api.GET("/v1/stores/:storeId/items", token, owner, keeperrest.GetItems)
		//api.PUT("/v1/stores/:storeId/items", token, owner, keeperrest.UpdateItems)
		api.GET("/v1/stores/:storeId/profiles", token, owner, keeperrest.GetStore)
		api.PUT("/v1/stores/:storeId", token, lc, owner, keeperrest.UpdateStore)
		//api.PUT("/v1/stores/:storeId/stores", token, lc, owner, keeperrest.ListStores)

		api.GET("/v1/stores/:storeId/orders", token, owner, keeperrest.ListOrders)
		api.GET("/v1/stores/:storeId/orders/:orderId", token, owner, keeperrest.GetOrder)
		api.PUT("/v1/stores/:storeId/orders/:orderId", token, lc, owner, keeperrest.UpdateOrder)

		api.GET("/v1/stores/:storeId/order-groups", token, owner, keeperrest.ListOrderGroups)
		api.GET("/v1/stores/:storeId/order-groups/:orderId/shares", token, owner, keeperrest.ListOrderGroupShares)
		api.GET("/v1/stores/:storeId/order-groups/:orderId", token, owner, keeperrest.GetOrderGroup)
		api.PUT("/v1/stores/:storeId/order-groups/:orderId", token, lc, owner, keeperrest.UpdateOrderGroup)

		api.GET("/v1/stores/:storeId/rewards", token, owner, keeperrest.ListRewards)
		api.PUT("/v1/stores/:storeId/rewards/:id", token, lc, owner, keeperrest.UpdateReward)

		api.GET("/v1/stores/:storeId/counters", token, owner, keeperrest.ListCounters)
		//api.GET("/v1/stores/:storeId/stores", token, owner, keeperrest.ListStores)

		api.GET("/v1/stores/:storeId/proxies", token, owner, keeperrest.ListProxies)
		api.POST("/v1/stores/:storeId/proxies", token, lc, owner, keeperrest.AddProxy)
		api.PUT("/v1/stores/:storeId/proxies/:proxyId", token, lc, owner, keeperrest.UpdateProxy)
		api.POST("/v1/stores/:storeId/proxies/:proxyId/users", token, lc, owner, keeperrest.AddProxyUser)
		api.GET("/v1/stores/:storeId/proxies/:proxyId/users", token, owner, keeperrest.ListProxyUsers)

		api.GET("/v1/stores/:storeId/proxy-rewards", token, owner, keeperrest.ListProxyRewards)
		api.PUT("/v1/stores/:storeId/proxy-rewards/:rewardId", token, lc, owner, keeperrest.UpdateProxyReward)

		api.GET("/v1/stores/:storeId/co-stores", token, owner, keeperrest.ListCoStores)
		api.POST("/v1/stores/:storeId/co-stores", token, lc, owner, keeperrest.AddCoStore)
		//api.GET("/v1/stores/:storeId/co-stores/:coStoreId", token, owner, keeperrest.GetCoStore)
		api.PUT("/v1/stores/:storeId/co-stores/:coStoreId", token, lc, owner, keeperrest.UpdateCoStore)

		api.GET("/v1/stores/:storeId/users", token, owner, keeperrest.GetUsers)
		api.GET("/v1/stores/:storeId/users/:userId", token, owner, keeperrest.GetStoreUser)
		api.PUT("/v1/stores/:storeId/users/:userId", token, owner, keeperrest.UpdateStoreUser)

		api.GET("/v1/stores/:storeId/feedbacks", token, owner, keeperrest.ListFeedbacks)

	}

	// user
	{
		api.GET("/v1/users/:id", token, rest.GetUser)
		api.PUT("/v1/users/:id", token, rest.UpdateUser)
		api.GET("/v1/users/:id/profiles", token, rStoreId, rest.GetUserProfile)

		api.GET("/v1/match-filters", token, rest.ListMatchFilters)
		api.GET("/v1/matches", token, rest.ListMatches)
		api.GET("/v1/stores/:storeId", token, rest.GetStore)
		api.POST("/v1/stores/:storeId/users", token, rest.AddStoreUser)
		api.GET("/v1/store-items", token, storeId, rest.ListItems)
		api.GET("/v1/current-issues", token, rStoreId, rest.GetCurrentIssue)

		api.POST("/v1/plans", token, lc, rStoreId, rest.AddPlan)
		api.DELETE("/v1/plans/:id", token, rStoreId, rest.DeletePlan)
		api.GET("/v1/plans", token, rStoreId, rest.ListPlans)

		api.POST("/v1/orders", token, lc, confirmedStoreId, rest.AddOrder)
		api.GET("/v1/orders", token, rStoreId, rest.ListOrders)
		api.GET("/v1/orders/:id", token, rStoreId, rest.GetOrder)
		api.PUT("/v1/orders/:id", token, lc, rStoreId, rest.Update)

		api.POST("/v1/order-groups", token, rStoreId, rest.AddOrderGroup)
		api.GET("/v1/order-groups", token, rStoreId, rest.ListOrderGroups)
		api.GET("/v1/order-groups/:id", token, rStoreId, rest.GetOrderGroup)
		api.GET("/v1/order-groups/:id/shares", token, rStoreId, rest.ListOrderGroupShares)
		api.POST("/v1/order-groups/:id/shares", token, uriLc, lc, rStoreId, rest.AddGroupShare)

		api.GET("/v1/proxies/:id", token, rStoreId, rest.GetProxy)
		api.GET("/v1/proxies/:id/users", token, rStoreId, rest.ListProxyUsers)

		api.GET("/v1/feedbacks", token, rStoreId, rest.ListFeedbacks)
		api.POST("/v1/feedbacks", token, lc, rStoreId, rest.AddFeedback)
		api.PUT("/v1/feedbacks/:id", token, lc, rStoreId, rest.UpdateFeedback)

	}

	err := engine.Run("0.0.0.0:" + args.ServerOptions.Port)
	if err != nil {
		log.Fatal(err)
	}

}
