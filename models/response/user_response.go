package response

import (
	"time"
)

type B struct {
}

func (B) TableName() string {
	return "member"
}

// UserResponse 用户实体返回结构体
type UserAllResponse struct {
	Mid            uint64    `json:"mid"`
	Name           string    `json:"name"`
	Password       string    `json:"-"`
	OriginPassword string    `json:"-"`
	Data           string    `json:"data"`
	Sessions       string    `json:"sessions"`
	Email          string    `json:"email"`
	Security       string    `json:"security"`
	Mobile         string    `json:"mobile"`
	Wechat         string    `json:"wechat"`
	Status         uint8     `json:"status"`
	Created        time.Time `json:"created"`
	Modify         time.Time `json:"modify"`
	EffectiveTime  time.Time `json:"effectiveTime"`
	ExpireTime     time.Time `json:"expireTime"`
	DelFlag        uint8     `json:"delFlag"`
	LastLoginTime  time.Time `json:"lastLoginTime"`
	LoginTimes     uint64    `json:"loginTimes"`
	Inviter        uint64    `json:"inviter"`
	Role           uint8     `json:"role"`
}

type UserListResponse struct {
	B
	Mid    uint64 `json:"mid"`
	Name   string `json:"name"`
	Status uint8  `json:"status"`
	Role   uint8  `json:"role"`
}
