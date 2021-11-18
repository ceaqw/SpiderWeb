package models

import (
	"SpiderWeb/models/req"
	"fmt"
	"time"
)

type CrawlerInformation struct {
	Id          uint64 `xorm:"pk autoincr" json:"id"`
	Project     string `xorm:"varchar(255) not null" json:"project"`
	Platform    string `xorm:"varchar(45) not null" json:"platform"`
	ScriptId    uint64 `xorm:"not null" json:"script_id"`
	Ip          string `xorm:"varchar(45) not null" json:"ip"`
	Comment     string `xorm:"text" json:"comment"`
	Server      uint8  `xorm:"not null" json:"server"`
	CriticalKpi string `xorm:"varchar(200)" json:"critical_kpi"`
	BindTable   string `xorm:"varchar(200)" json:"bind_table"`
	DelFlag     uint8  `json:"del_flag"`
}

func (CrawlerInformation) TableName() string {
	return "crawler_information"
}

type CrawlerInformationOrm struct {
	BaseDB
	CrawlerInformation
}

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
	BaseDB
	Project
}

type CrawlerStat struct {
	Project string    `xorm:"pk varchar(200) notnull 'project'" json:"project"`          //
	Date    time.Time `xorm:"pk index(date_hour) date notnull 'date'" json:"date"`       //
	Hour    int       `xorm:"pk index(date_hour) tinyint(4) notnull 'hour'" json:"hour"` //
	Success int       `xorm:"int(10) default 0 'success'" json:"success"`                //
	Fail    int       `xorm:"int(10) default 0 'fail'" json:"fail"`                      //
	Retry   int       `xorm:"int(10) default 0 'retry'" json:"retry"`                    //
	Undone  int       `xorm:"int(10) default 0 'undone'" json:"undone"`                  //
	Total   int       `xorm:"int(10) default 0 'total'" json:"total"`                    //
}

func (CrawlerStat) TableName() string {
	return "crawler_stat"
}

type CrawlerStatOrm struct {
	BaseDB
	Project
	CrawlerStat
	CrawlerInformation
}

func (m CrawlerInformationOrm) GetProjectListWithPlatform() ([]map[string][]byte, error) {
	result, err := m.Query("select platform, project from crawler_information where del_flag <> 1 order by platform, project")
	return result, err
}

func (m CrawlerStatOrm) GetAllDatas(filter req.Filter) ([]map[string][]byte, error) {
	sql := `select 
			sum(cs.success) as success,
			sum(cs.fail) as fail,
			sum(cs.undone) as undone,
			sum(cs.retry) as retry,
			sum(cs.total) as total,
			%s
			from crawler_stat cs
			join crawler_information ci using(project)
			where 1`
	// 拿去坐标字段
	// [天， 小时]
	if filter.PlatForm != "all" {
		sql = fmt.Sprintf("%s and ci.platform = '%s'", sql, filter.PlatForm)
	}
	if filter.Project != "all" {
		sql = fmt.Sprintf("%s and cs.project = '%s'", sql, filter.Project)
	}

	// 时间特殊处理
	req.FilterVerify(&filter)
	sql = fmt.Sprintf("%s and cs.date >= '%s' and cs.date <= '%s' ", sql, filter.StartDate, filter.EndDate)
	// [天， 小时]
	if filter.ShowType == 0 {
		sql = fmt.Sprintf("%s group by cs.date order by cs.date", sql)
		sql = fmt.Sprintf(sql, "concat(cs.date, '日') as date ")
	} else if filter.ShowType == 1 {
		sql = fmt.Sprintf("%s group by cs.date, cs.hour order by cs.date, cs.hour", sql)
		sql = fmt.Sprintf(sql, "concat(cs.date, '日', cs.hour, '时') as date ")
	}
	result, err := m.Query(sql)
	return result, err
}

// func (m CrawlerStatOrm) GetAlanyseDatas(filter req.Filter) ([]map[string][]byte, error) {

// }
