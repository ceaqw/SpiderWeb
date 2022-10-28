package services

import (
	"SpiderWeb/middleware/jwt"
	"SpiderWeb/models"
	"time"

	"github.com/druidcaesa/gotool"
)

// UserService 用户操作业务逻辑
type UserService struct {
	userModel models.UserOrm
}

// GetUserByUserName 根据用户名查询用户
func (s UserService) GetUserByUserName(name string) *models.User {
	user := models.User{}
	user.Name = name
	return s.userModel.GetUserByUserName(user.Name)
}

func (s UserService) Login(name string, password string) (bool, string, *models.User) {
	user := s.GetUserByUserName(name)
	// fmt.Println(user)
	if user == nil || !gotool.BcryptUtils.CompareHash(user.Password, password) {
		return false, "用户名或密码错误", user
	}
	// 更新登陆信息
	user.LoginTimes += 1
	user.LastLoginTime = time.Now()
	s.userModel.UpdateInfo(user.Mid, *user, "login_times", "last_login_time")
	//生成token
	token, err := jwt.NewJWT().CreateUserToken(*user)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
		return false, "", user
	}
	//数据存储到redis中
	return true, token, user
}
