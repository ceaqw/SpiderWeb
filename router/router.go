package router

import (
	"SpiderWeb/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.New()
	router.Use(jwt.JWTAuth())
	// vi版本api
	v1Router := router.Group("/api/v1")

	{
		initUserRouter(v1Router)
		initDashBoardRouter(v1Router)
		initProjectRouter(v1Router)
		initTrendRouter(v1Router)
		initMemberRouter(v1Router)
	}
	return router
}
