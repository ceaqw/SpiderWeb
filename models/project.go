/*
 * @Date: 2021-12-09 14:01:13
 * @LastEditTime: 2021-12-10 16:41:00
 * @Author: ceaqw
 */
package models

import (
	"SpiderWeb/models/req"
	"fmt"
	"strings"
)

type Project struct {
	Id        int    `xorm:"int(10) notnull pk autoincr 'id'" json:"id"`                        //
	Name      string `xorm:"varchar(45) unique notnull 'name'" json:"name"`                     //
	FullName  string `xorm:"varchar(255) notnull 'full_name'" json:"full_name"`                 //
	Priority  int    `xorm:"int(10) notnull default 0 'priority'" json:"priority"`              //
	TaskTable string `xorm:"varchar(45) notnull default 'task' 'task_table'" json:"task_table"` //
	Status    int    `xorm:"tinyint(3) notnull default 0 'status'" json:"status"`               // 0: normal; 1: deleted
	Lua       string `xorm:"varchar(255) default NULL 'lua'" json:"lua"`                        // LUA SCRIPT ID
	Script    string `xorm:"varchar(255) default NULL 'script'" json:"script"`                  // insert/do script
}

func (t *Project) TableName() string {
	return "project"
}

type ProjectOrm struct {
	MainDbHand
	Project
}

func (m ProjectOrm) GetProjectInfo(filter req.Filter, cols ...string) ([]map[string][]byte, error) {
	sql := `select 
			%s
			from crawler_information ci 
			left join project p on ci.project = p.name 
			where ci.del_flag = 0`
	sql = fmt.Sprintf(sql, strings.Join(cols, ","))
	if filter.PlatForm != "all" {
		sql = fmt.Sprintf("%s and ci.platform = '%s'", sql, filter.PlatForm)
	}
	if filter.Project != "all" {
		sql = fmt.Sprintf("%s and ci.project = '%s'", sql, filter.Project)
	}
	projectInfo, err := m.Query(sql)
	return projectInfo, err
}
