package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	router := gin.Default()

	// 配置跨域
	//router.Use(middlewares.Cors())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // 允许所有前端域
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true, // 允许携带 Cookie
	}))

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
