package services

import (
	"SpiderWeb/middleware/jwt"
	"SpiderWeb/models"

	"github.com/druidcaesa/gotool"
)

type LoginService struct {
	userService UserService
}

// Login 用户登录业务处理
func (s LoginService) Login(name string, password string) (bool, string) {
	// user := s.userService.GetUserByUserName(name)
	user := models.User{
		Name:     "ceaqw",
		Password: "$2a$10$aUKFgkqWVOxK2kidLtwry.ZpIWU//HO3BTqlE7cHvqSsSB1h/blXi",
		UserId:   1,
	}
	// if user == nil {
	// 	return false, "用户不存在"
	// }
	if !gotool.BcryptUtils.CompareHash(user.Password, password) {
		return false, "密码错误"
	}
	//生成token
	token, err := jwt.NewJWT().CreateUserToken(user)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
		return false, ""
	}
	//数据存储到redis中
	return true, token
}
