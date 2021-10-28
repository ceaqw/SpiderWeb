package router

import (
	v1 "SpiderWeb/api/v1"

	"github.com/gin-gonic/gin"
)

func initUserRouter(router *gin.RouterGroup) {
	userApi := new(v1.User)

	userRouter := router.Group("/user")

	{
		userRouter.POST("/login", userApi.Login)
		userRouter.GET("/getInfo", userApi.GetInfo)
		userRouter.POST("/logout", userApi.LoginOut)
	}
}
