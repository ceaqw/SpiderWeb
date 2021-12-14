/*
 * @Date: 2021-11-22 09:49:48
 * @LastEditTime: 2021-12-14 14:09:01
 * @Author: ceaqw
 */
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
		initProjectRouter(baseRouter)
		initTrendRouter(baseRouter)
		initChartRouter(baseRouter)
	}
	return router
}
