// base 板块子路由

package router

import (
	v1 "SpiderWeb/api/v1"

	"github.com/gin-gonic/gin"
)

func initBaseRouter(router *gin.RouterGroup) {
	baseApi := new(v1.Base)

	dashBoardRouter := router.Group("/base")

	{
		dashBoardRouter.GET("projectList", baseApi.ProjectList)
	}
}
