//dashBoard 板块子路由
package router

import (
	v1 "SpiderWeb/api/v1"

	"github.com/gin-gonic/gin"
)

func initTrendRouter(router *gin.RouterGroup) {
	trendApi := new(v1.Trend)

	trendKpiRouter := router.Group("/trend/kpi")
	trendFailRouter := router.Group("/trend/fail")

	{
		trendKpiRouter.GET("/list", trendApi.Kpi.List)
	}

	{
		trendFailRouter.GET("/list", trendApi.Fail.List)
	}
}
