package router

import (
	v1 "SpiderWeb/api/v1"

	"github.com/gin-gonic/gin"
)

func initProjectRouter(router *gin.RouterGroup) {
	projectApi := new(v1.Project)

	projectRouter := router.Group("/project")

	{
		projectRouter.GET("/list", projectApi.List)
	}
}
