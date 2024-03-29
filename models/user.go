/*
 * @Date: 2021-11-16 09:08:27
 * @LastEditTime: 2021-12-14 14:33:16
 * @Author: ceaqw
 */
package models

import (
	"SpiderWeb/conf"
	"SpiderWeb/models/response"
	"fmt"
	"time"

	"github.com/druidcaesa/gotool"
)

type User struct {
	Mid            uint64    `xorm:"pk autoincr" json:"mid"`
	Name           string    `xorm:"varchar(255) notnull unique" json:"name"`
	Password       string    `xorm:"varchar(255)" json:"password"`
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
	DelFlag        uint8     `xorm:"tinyint default 0" json:"del_flag"`
	LastLoginTime  time.Time `json:"last_login_time"`
	LoginTimes     uint64    `xorm:"default 0" json:"login_times"`
	Inviter        uint64    `json:"inviter"`
	Role           uint8     `xorm:"default 2" json:"role"`
}

type UserOrm struct {
	MainDbHand
}

func (User) TableName() string {
	return "member"
}

func (m UserOrm) GetUserByUserName(userName string) *User {
	user := User{}
	row, err := MainSqlDb.Where("name = ?", userName).Get(&user)

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

func (m UserOrm) GetUserInfo(name string, cols string) ([]map[string][]byte, error) {
	return MainSqlDb.Query(fmt.Sprintf("select %s from member where name = ? limit 1", cols), name)
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

func (m UserOrm) AddUser(user User) error {
	if user.Role == 0 {
		user.Role = 2
	}
	_, err := MainSqlDb.InsertOne(user)
	return err
}
