/*
 * @Date: 2021-10-25 17:00:59
 * @LastEditTime: 2021-12-10 17:12:57
 * @Author: ceaqw
 */
package router

import (
	v1 "SpiderWeb/api/v1"

	"github.com/gin-gonic/gin"
)

func initProjectRouter(router *gin.RouterGroup) {
	projectApi := new(v1.Project)

	projectRouter := router.Group("/project")

	{
		projectRouter.GET("projectList", projectApi.ProjectList)
		projectRouter.GET("projectInfos", projectApi.GetProjectInfos)
		projectRouter.GET("allPlatformAndProjectMap", projectApi.GetAllPlatformAndProjectMap)
	}
}
