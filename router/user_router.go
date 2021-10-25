package router

import (
	v1 "SpiderWeb/api/v1"

	"github.com/gin-gonic/gin"
)

func initUserRouter(router *gin.RouterGroup) {
	userApi := new(v1.User)

	userRouter := router.Group("/user")

	{
		userRouter.GET("/login", userApi.Login)
		userRouter.GET("/logout", userApi.LoginOut)
	}
}
