package models

import "time"

type User struct {
	UserId         uint64    `xorm:"pk autoincr" json:"userId"`
	Name           string    `xorm:"varchar(255)" json:"name"`
	Password       string    `xorm:"varchar(255)" json:"-"`
	OriginPassword string    `xorm:"varchar(255)" json:"-"`
	Data           string    `xorm:"text" json:"data"`
	Sessions       string    `xorm:"varchar(1000)" json:"sessions"`
	Email          string    `xorm:"varchar(255)" json:"email"`
	Security       string    `xorm:"varchar(255)" json:"security"`
	Mobile         string    `xorm:"varchar(255)" json:"mobile"`
	Wechat         string    `xorm:"varchar(255)" json:"wechat"`
	Statue         uint8     `xorm:"tinyint" json:"status"`
	Created        time.Time `xorm:"created" json:"created"`
	Modify         time.Time `json:"modify"`
	EffectiveTime  time.Time `json:"effectiveTime"`
	ExpireTime     time.Time `json:"expireTime"`
	DelFlag        uint8     `xorm:"tinyint" json:"delFlag"`
	LastLoginTime  time.Time `json:"lastLoginTime"`
	LoginTimes     uint64    `json:"loginTimes"`
	Inviter        uint64    `json:"inviter"`
	Role           uint8     `json:"role"`
}

func (User) TableName() string {
	return "user"
}
