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
		userRouter.POST("/logout", userApi.LoginOut)
		userRouter.POST("/userList", userApi.UserList)
		userRouter.POST("/option", userApi.Option)
		userRouter.POST("/getRoles", userApi.GetRoles)
		userRouter.POST("/register", userApi.Register)
		userRouter.GET("/validator/:type", userApi.Validator)
	}
}
