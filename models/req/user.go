package req

type LoginBody struct {
	UserName string `json:"username"` //用户名
	Password string `json:"password"` //密码
}

type UserListBody struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

type UserOption struct {
	Option string `json:"option"`
	Mid    int    `json:"mid"`
	Name   string `json:"name"`
	Role   int    `json:"role"`
}
