package req

type LoginBody struct {
	UserName string `json:"username"` //用户名
	Password string `json:"password"` //密码
}
