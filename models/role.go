package models

type Role struct {
	Id   uint64 `xorm:"pk autoincr" json:"id"`
	Name string `xorm:"varchar(255)" json:"name"`
}

func (Role) TableName() string {
	return "role"
}

type RoleOrm struct {
	BaseDB
}

func (RoleOrm) GetRoles() []Role {
	var roles []Role
	MainSqlDb.Find(&roles)
	return roles
}

func (RoleOrm) GetRoleIdByName(roleName string) (uint8, error) {
	var role Role
	_, err := MainSqlDb.Where("role = ?", roleName).Get(&role)
	return uint8(role.Id), err
}
