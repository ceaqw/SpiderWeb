package utils

import (
	"SpiderWeb/models/response"

	"github.com/gin-gonic/gin"
)

type UserUtils struct {
}

func (h UserUtils) GetUserInfo(c *gin.Context) *response.UserResponse {
	user := &response.UserResponse{}
	return user
}
