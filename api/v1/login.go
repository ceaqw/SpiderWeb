package v1

import "github.com/gin-gonic/gin"

type Login struct {
}

func (Login) Login(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}
