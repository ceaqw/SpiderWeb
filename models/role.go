package models

type Role struct {
	Id   uint64 `xorm:"pk autoincr" json:"id"`
	Name string `xorm:"varchar(255)" json:"name"`
}

func (Role) TableName() string {
	return "role"
}
