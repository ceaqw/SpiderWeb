/*
 * @Date: 2021-12-10 16:44:53
 * @LastEditTime: 2021-12-10 17:04:09
 * @Author: ceaqw
 */
package models

type AllMongoProject struct {
	Platform string `xorm:"varchar(45) notnull default '' 'platform'" json:"platform"` //
	Project  string `xorm:"varchar(200) notnull default '' 'project'" json:"project"`  // 项目
}

//TableName Get table name
func (t AllMongoProject) TableName() string {
	return "all_mongo_project"
}

type AllMongoProjectOrm struct {
	MainDbHand
}

func (t AllMongoProjectOrm) GetAllMongoProject() ([]AllMongoProject, error) {
	rest := make([]AllMongoProject, 0)
	err := MainSqlDb.Where("platform <> 'spider_raw'").Find(&rest)
	return rest, err
}
