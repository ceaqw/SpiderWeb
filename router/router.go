package router

import "github.com/gin-gonic/gin"

func Init() *gin.Engine {
	router := gin.New()

	// vi版本api
	v1Router := router.Group("/api/v1")

	{
		initLoginRouter(v1Router)
	}
	return router
}
