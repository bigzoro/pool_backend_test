package routers

import (
	"github.com/gin-gonic/gin"
	"pool/api"
)

func InitPoolRouter(Router *gin.RouterGroup) {
	poolRouter := Router.Group("pool")
	{
		// 首页信息
		poolRouter.GET("index_data", api.IndexData)
		poolRouter.GET("hashrate", api.HashRate)
		poolRouter.GET("get_pools", api.GetPools)

		// 地址
		poolRouter.GET("get_address", api.GetAddressByUserId)
		poolRouter.POST("add_address", api.AddUserAddress)
		poolRouter.GET("delete_address", api.DeleteAddress)

		// 池子
		poolRouter.POST("purchase_pool", api.PurchasePool)
		poolRouter.POST("getPurchaseByUserId", api.GetPurchaseByUserId)

		// 购买
		poolRouter.GET("get_all_purchases", api.GetAllPurchases)

		// 方案
		poolRouter.POST("add_plan", api.AddPlan)
		poolRouter.POST("get_user_plan", api.GetUserPlan)
		poolRouter.POST("get_plan_detail", api.GetUserPlanDetails)
		poolRouter.POST("delete_plan", api.DeletePlan)

		// 区块
		poolRouter.POST("get_block_page", api.GetBlocksByPage)
		poolRouter.GET("block_info", api.BlockInfo)

		// 谷歌
		poolRouter.POST("get_google_auth_qr", api.GetGoogleAuthQR)
		poolRouter.POST("verify_google_auth_code", api.VerifyGoogleAuthCode)
	}
}
