package v1

import (
	"SpiderWeb/models/req"
	"SpiderWeb/modules/utils"
	"SpiderWeb/services"
	"SpiderWeb/services/resp"
	"fmt"

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

func (h User) GetInfo(c *gin.Context) {
	m := make(map[string]interface{})
	user := h.loginService.LoginUser(c)
	fmt.Println(user)
	//查询用户角色集合
	// roleKeys := a.permissionService.GetRolePermissionByUserId(user)
	// 权限集合
	// perms := a.permissionService.GetMenuPermission(user)
	// m["roles"] = roleKeys
	// m["permissions"] = perms
	// m["user"] = user
	c.JSON(200, resp.Success(m))
}

func (h User) LoginOut(c *gin.Context) {
	name := utils.UserUtils.GetUserInfo(utils.UserUtils{}, c).Name
	fmt.Println(name)
	// cache.RemoveKey(name)
	resp.OK(c)
}
