package router

import (
	"SpiderWeb/middleware/logger"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.New()
	// 身份验证
	// router.Use(jwt.JWTAuth())
	// 日志记录
	router.Use(logger.LoggerToFile())
	// vi版本api
	v1Router := router.Group("/api/v1")

	{
		initUserRouter(v1Router)
		initDashBoardRouter(v1Router)
		initProjectRouter(v1Router)
		initTrendRouter(v1Router)
		initMemberRouter(v1Router)
	}
	return router
}
