/*
 * @Date: 2021-12-09 14:09:51
 * @LastEditTime: 2021-12-09 14:21:36
 * @Author: ceaqw
 */
package models

import (
	"SpiderWeb/models/req"
	"fmt"
	"time"
)

type JobMonito struct {
	Platform    string    `xorm:"varchar(45) notnull 'platform'" json:"platform"`               //
	TableNames  string    `xorm:"varchar(100) notnull 'table_name'" json:"table_name"`          //
	Date        time.Time `xorm:"date notnull 'date'" json:"date"`                              //
	Hour        int       `xorm:"int(10) notnull default 0 'hour'" json:"hour"`                 //
	AllCount    int       `xorm:"int(10) notnull default 0 'all_count'" json:"all_count"`       //
	FinishCount int       `xorm:"int(10) notnull default 0 'finish_count'" json:"finish_count"` //
	UndoCount   int       `xorm:"int(10) default 0 'undo_count'" json:"undo_count"`             //
	FailCount   int       `xorm:"int(10) default 0 'fail_count'" json:"fail_count"`             //
}

//TableName Get table name
func (t *JobMonito) TableName() string {
	return "job_monito"
}

func (m CrawlerStatOrm) GetAnalyseDatas(filter req.Filter) ([]map[string][]byte, error) {
	sql := `select 
			%s
			sum(all_count) as total, 
			sum(finish_count) as success, 
			sum(undo_count) as undone, 
			sum(fail_count) as fail 
			FROM job_monitor jm
			where 1`
	sql = fmt.Sprintf("%s and jm.date >= '%s' and jm.date <= '%s' ", sql, filter.StartDate, filter.EndDate)
	sql = fmt.Sprintf("%s and platform='spider_raw'", sql)
	// if filter.PlatForm != "all" {
	// 	sql = fmt.Sprintf("%s and platform='%s'", sql, filter.PlatForm)
	// }
	if filter.Project != "all" {
		sql = fmt.Sprintf("%s and table_name='%s'", sql, filter.Project)
	}
	if filter.ShowType == 1 {
		sql = fmt.Sprintf(sql, "concat(date, ' ', hour) as date,")
		sql = fmt.Sprintf("%s group by jm.date, jm.hour ORDER BY jm.date, jm.hour", sql)
	} else if filter.ShowType == 0 {
		sql = fmt.Sprintf(sql, "concat(date, '日') as date,")
		sql = fmt.Sprintf("%s group by jm.date ORDER BY jm.date", sql)
	}
	result, err := m.Query(sql)
	return result, err
}
