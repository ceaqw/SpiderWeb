package router

import (
	"SpiderWeb/conf"
	"SpiderWeb/middleware/logger"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.New()
	// 身份验证
	// router.Use(jwt.JWTAuth())
	// 日志记录
	router.Use(logger.LoggerToFile())
	// 版本api
	apiVersion := conf.GetApiVersion()
	v1Router := router.Group(fmt.Sprintf("/api/%s", apiVersion))

	{
		initUserRouter(v1Router)
		initDashBoardRouter(v1Router)
		initProjectRouter(v1Router)
		initTrendRouter(v1Router)
		initMemberRouter(v1Router)
	}
	return router
}
