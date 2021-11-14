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
	baseRouter := router.Group(fmt.Sprintf("/api/%s", apiVersion))

	{
		initUserRouter(baseRouter)
		initDashBoardRouter(baseRouter)
		initProjectRouter(baseRouter)
		initTrendRouter(baseRouter)
		initMemberRouter(baseRouter)
		initBaseRouter(baseRouter)
		initChartRouter(baseRouter)
	}
	return router
}
