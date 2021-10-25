package router

import (
	v1 "SpiderWeb/api/v1"

	"github.com/gin-gonic/gin"
)

func initMemberRouter(router *gin.RouterGroup) {
	memberApi := new(v1.Member)

	memberRouter := router.Group("/member")

	{
		memberRouter.GET("/list", memberApi.List)
	}
}
