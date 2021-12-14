/*
 * @Date: 2021-11-08 09:05:42
 * @LastEditTime: 2021-12-14 14:09:35
 * @Author: ceaqw
 */
package router

import (
	v1 "SpiderWeb/api/v1"

	"github.com/gin-gonic/gin"
)

func initUserRouter(router *gin.RouterGroup) {
	userApi := new(v1.User)

	userRouter := router.Group("user")

	{
		userRouter.POST("login", userApi.Login)
		userRouter.POST("logout", userApi.LoginOut)
		userRouter.POST("userList", userApi.UserList)
		userRouter.GET("user", userApi.GetUserInfo)
		userRouter.POST("option", userApi.Option)
		userRouter.POST("getRoles", userApi.GetRoles)
		userRouter.POST("register", userApi.Register)
		userRouter.POST("update", userApi.Update)
		userRouter.GET("validator/:type", userApi.Validator)
	}
}
