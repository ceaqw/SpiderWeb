package v1

import "github.com/gin-gonic/gin"

type User struct {
}

func (User) Login(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}

func (User) LoginOut(c *gin.Context) {

}
