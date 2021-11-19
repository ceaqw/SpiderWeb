// base 板块子路由

package router

import (
	v1 "SpiderWeb/api/v1"

	"github.com/gin-gonic/gin"
)

func initChartRouter(router *gin.RouterGroup) {
	chartDataApi := new(v1.ChartData)

	dashBoardRouter := router.Group("/chart")

	{
		dashBoardRouter.POST("all", chartDataApi.AllChartData)
		dashBoardRouter.POST("analyse", chartDataApi.AnalyseDatas)
	}
}
