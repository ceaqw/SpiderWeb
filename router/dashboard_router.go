// dashBoard 板块子路由

package router

import (
	v1 "SpiderWeb/api/v1"

	"github.com/gin-gonic/gin"
)

func initDashBoardRouter(router *gin.RouterGroup) {
	dashBoardApi := new(v1.DashBoard)

	dashBoardRouter := router.Group("/dashboard")

	{
		dashBoardRouter.GET("list", dashBoardApi.List)
		dashBoardRouter.GET("projectList", dashBoardApi.ProjectList)
	}
}
