package models

import (
	"SpiderWeb/conf"
	"SpiderWeb/models/response"
	"time"

	"github.com/druidcaesa/gotool"
)

type User struct {
	Mid            uint64    `xorm:"pk autoincr" json:"mid"`
	Name           string    `xorm:"varchar(255)" json:"name"`
	Password       string    `xorm:"varchar(255)" json:"-"`
	OriginPassword string    `xorm:"varchar(255)" json:"-"`
	Data           string    `xorm:"text" json:"data"`
	Sessions       string    `xorm:"varchar(1000)" json:"sessions"`
	Email          string    `xorm:"varchar(255)" json:"email"`
	Security       string    `xorm:"varchar(255)" json:"security"`
	Mobile         string    `xorm:"varchar(255)" json:"mobile"`
	Wechat         string    `xorm:"varchar(255)" json:"wechat"`
	Status         uint8     `xorm:"tinyint default 0" json:"status"`
	Created        time.Time `xorm:"created default current_timestamp" json:"created"`
	Modify         time.Time `json:"modify"`
	EffectiveTime  time.Time `json:"effectiveTime"`
	ExpireTime     time.Time `json:"expireTime"`
	DelFlag        uint8     `xorm:"tinyint default 0" json:"delFlag"`
	LastLoginTime  time.Time `json:"lastLoginTime"`
	LoginTimes     uint64    `xorm:"default 0" json:"loginTimes"`
	Inviter        uint64    `json:"inviter"`
	Role           uint8     `xorm:"default 2" json:"role"`
}

type UserOrm struct {
}

func (User) TableName() string {
	return "member"
}

func (m UserOrm) GetUserByUserName(user User) *User {
	row, err := MainSqlDb.Get(&user)

	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
	}
	if row {
		return &user
	}
	return nil
}

func (m UserOrm) UpdateInfo(id uint64, user User, cols ...string) (int64, error) {
	return MainSqlDb.ID(id).Cols(cols...).Update(&user)
}

func (m UserOrm) GetUserList(page int, pageSize int) ([]response.UserListResponse, int64) {
	var users []response.UserListResponse
	if conf.GetPageSize() < pageSize {
		pageSize = conf.GetPageSize()
	}
	total, _ := MainSqlDb.Where("1=1").Count(new(User))
	MainSqlDb.Cols("mid", "name", "status", "role").Where("1=1").Limit(pageSize, pageSize*(page-1)).Find(&users)
	return users, total
}
