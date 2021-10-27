package services

import (
	"SpiderWeb/models"
)

// UserService 用户操作业务逻辑
type UserService struct {
	userModel models.UserOrm
}

// GetUserByUserName 根据用户名查询用户
func (s UserService) GetUserByUserName(name string) *models.User {
	user := models.User{}
	user.Name = name
	return s.userModel.GetUserByUserName(user)
}
