package router

import (
	v1 "SpiderWeb/api/v1"

	"github.com/gin-gonic/gin"
)

func initLoginRouter(router *gin.RouterGroup) {
	loginApi := new(v1.Login)

	loginRouter := router.Group("/")

	{
		loginRouter.GET("/login", loginApi.Login)
	}
}
