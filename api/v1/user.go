package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type User struct {
}

func (User) Login(c *gin.Context) {
	fmt.Println(c.PostForm("username"))
	c.JSON(200, gin.H{"message": "OK"})
}

func (User) LoginOut(c *gin.Context) {

}
