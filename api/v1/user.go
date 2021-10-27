package v1

import (
	"SpiderWeb/models/req"
	"SpiderWeb/services"
	"SpiderWeb/services/resp"

	"github.com/gin-gonic/gin"
)

type User struct {
	loginService services.LoginService
}

func (h User) Login(c *gin.Context) {
	loginBody := req.LoginBody{}
	if c.BindJSON(&loginBody) == nil {
		m := make(map[string]string)
		login, s := h.loginService.Login(loginBody.UserName, loginBody.Password)
		if login {
			//将token存入到redis中
			// user_util.SaveRedisToken(loginBody.UserName, s)
			m["token"] = s
			c.JSON(200, resp.Success(m))
		} else {
			c.JSON(200, resp.ErrorResp(s))
		}
	} else {
		c.JSON(200, resp.ErrorResp(500, "参数绑定错误"))
	}
}

func (User) LoginOut(c *gin.Context) {

}
