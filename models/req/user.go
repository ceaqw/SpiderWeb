/*
 * @Date: 2021-11-08 09:05:42
 * @LastEditTime: 2021-12-14 14:13:03
 * @Author: ceaqw
 */
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

type UserUpdate struct {
	Mid      uint64 `json:"mid"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Wechat   string `json:"wechat"`
}
