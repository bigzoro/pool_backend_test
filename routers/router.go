package routers

import (
	"github.com/gin-gonic/gin"
	"pool/middlewares"
)

func InitRouters() *gin.Engine {
	router := gin.Default()

	// 配置跨域
	router.Use(middlewares.Cors())

	// 设置上传文件保存的目录
	uploadDir := "./uploads"
	router.Static("/uploads", uploadDir)

	ApiGroup := router.Group("/api/v1")
	InitUserRouter(ApiGroup)
	InitPoolRouter(ApiGroup)
	//front_routers.InitFront(ApiGroup)
	//plum_routers.InitPlumRouters(ApiGroup)

	return router
}
