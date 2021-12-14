package v1

import (
	"SpiderWeb/models"
	"SpiderWeb/models/req"
	"SpiderWeb/modules/redis_helper"
	"SpiderWeb/services"
	"SpiderWeb/services/resp"
	"strconv"

	"github.com/druidcaesa/gotool"
	"github.com/gin-gonic/gin"
)

type User struct {
	userService services.UserService
	redisHelper redis_helper.RedisHelper
	userModel   models.UserOrm
	roleModel   models.RoleOrm
}

func (h User) Login(c *gin.Context) {
	loginBody := req.LoginBody{}
	if c.BindJSON(&loginBody) == nil {
		result := make(map[string]interface{})
		login, token, user := h.userService.Login(loginBody.UserName, loginBody.Password)
		if login {
			//将token存入到redis中
			h.redisHelper.SaveRedisToken(loginBody.UserName, token)
			result["token"] = token
			result["userAuth"] = user.Role
			result["userId"] = user.Mid
			c.JSON(200, resp.Success(result))
		} else {
			c.JSON(200, resp.ErrorResp(token))
		}
	} else {
		c.JSON(200, resp.ErrorResp(500, "参数绑定错误"))
	}
}

func (h User) LoginOut(c *gin.Context) {
	name, _ := c.Cookie("user")
	if h.redisHelper.RemoveRedisToken(name) != nil {
		resp.Error(c)
	}
	resp.OK(c)
}

func (h User) UserList(c *gin.Context) {
	userListBody := req.UserListBody{}
	if c.BindJSON(&userListBody) == nil {
		userList, total := h.userModel.GetUserList(userListBody.Page, userListBody.PageSize)
		result := make(map[string]interface{})
		result["userList"] = userList
		result["total"] = total
		c.JSON(200, resp.Success(result))
	} else {
		c.JSON(200, resp.ErrorResp(500, "参数错误"))
	}
}

func (h User) GetUserInfo(c *gin.Context) {
	user, ok := c.GetQuery("user")
	if ok {
		userInfo, err := h.userModel.GetUserInfo(user, "mid,email,mobile,wechat")
		if err == nil {
			response := make(map[string]interface{})
			response["name"] = user
			if len(userInfo) > 0 {
				response["mid"], _ = strconv.Atoi(string(userInfo[0]["mid"]))
				response["email"] = string(userInfo[0]["email"])
				response["mobile"] = string(userInfo[0]["mobile"])
				response["wechat"] = string(userInfo[0]["wechat"])
			}
			c.JSON(200, resp.Success(response))
		} else {
			resp.Error(c, "查询错误")
		}
	} else {
		resp.Error(c, "错误参数")
	}
}

func (h User) Option(c *gin.Context) {
	optionBody := req.UserOption{}
	if c.BindJSON(&optionBody) == nil {
		user := models.User{}

		if optionBody.Option == "del" {
			user.Status = 1
			user.DelFlag = 1
		} else if optionBody.Option == "recover" {
			user.Status = 0
			user.DelFlag = 0
		} else if optionBody.Option == "role" {
			user.Role = uint8(optionBody.Role)
		}
		var err error
		if optionBody.Option == "role" {
			_, err = h.userModel.UpdateInfo(uint64(optionBody.Mid), user, "role")
		} else {
			_, err = h.userModel.UpdateInfo(uint64(optionBody.Mid), user, "status", "del_flag")
		}
		if err != nil {
			c.JSON(200, resp.ErrorResp(500, "操作执行错误"))
		} else {
			resp.OK(c)
		}
	} else {
		c.JSON(200, resp.ErrorResp(500, "参数错误"))
	}
}

func (h User) GetRoles(c *gin.Context) {
	roles := h.roleModel.GetRoles()
	result := make(map[string]interface{})
	result["roles"] = roles
	c.JSON(200, resp.Success(result))
}

func (h User) Register(c *gin.Context) {
	user := models.User{}
	if c.BindJSON(&user) == nil {
		if h.userModel.GetUserByUserName(user.Name) != nil {
			c.JSON(200, resp.ErrorResp(500, "用户已存在"))
		} else {
			user.Password = string(gotool.BcryptUtils.Generate(user.Password))
			if h.userModel.AddUser(user) != nil {
				c.JSON(200, resp.ErrorResp(500, "新增用户失败"))
			} else {
				resp.OK(c)
			}
		}
	} else {
		resp.ParamError(c)
	}
}

func (h User) Update(c *gin.Context) {
	user := req.UserUpdate{}
	if c.BindJSON(&user) == nil {
		cols := []string{"name", "email", "mobile", "wechat"}
		userInfo := models.User{
			Name:   user.Name,
			Email:  user.Email,
			Mobile: user.Mobile,
			Wechat: user.Wechat,
		}
		if user.Password != "" {
			userInfo.Password = string(gotool.BcryptUtils.Generate(user.Password))
			cols = append(cols, "password")
		}
		_, err := h.userModel.UpdateInfo(user.Mid, userInfo, cols...)
		if err == nil {
			resp.OK(c, "ok")
		} else {
			resp.Error(c, "fail")
		}
	} else {
		resp.ParamError(c)
	}
}

func (h User) Validator(c *gin.Context) {
	optionTpye := c.Param("type")
	user := models.User{}
	user.Name = c.Query("name")
	if user.Name != "" {
		if optionTpye == "name" {
			if h.userModel.GetUserByUserName(user.Name) != nil {
				c.JSON(200, resp.ErrorResp(200, "exits"))
			} else {
				c.JSON(200, resp.ErrorResp(200, "ok"))
			}
		}
	} else {
		c.JSON(200, resp.ErrorResp(500, "参数错误"))
	}
}
