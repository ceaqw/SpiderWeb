/*
 * @Date: 2021-12-09 14:09:51
 * @LastEditTime: 2021-12-14 15:52:57
 * @Author: ceaqw
 */
package models

import (
	"SpiderWeb/models/req"
	"fmt"
	"strings"
	"time"
)

type JobMonitor struct {
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
func (t JobMonitor) TableName() string {
	return "job_monitor"
}

type JobMonitorOrm struct {
	MainDbHand
	JobMonitor
}

func (m JobMonitorOrm) WarpperSql(filter req.Filter) string {
	filterStr := ""
	filterStr = fmt.Sprintf("%s and date >= '%s' and date <= '%s' ", filterStr, filter.StartDate, filter.EndDate)
	if filter.ShowType == 1 {
		filterStr = fmt.Sprintf(
			"%s group by %s.date, %s.hour ORDER BY %s.date, %s.hour",
			filterStr,
			m.TableName(),
			m.TableName(),
			m.TableName(),
			m.TableName(),
		)
	} else if filter.ShowType == 0 {
		filterStr = fmt.Sprintf(
			"%s group by %s.date ORDER BY %s.date",
			filterStr,
			m.TableName(),
			m.TableName(),
		)
	}
	return filterStr
}

func (m JobMonitorOrm) GetDoDatas(filter req.Filter) ([]map[string][]byte, error) {
	sql := `select 
			%s
			sum(all_count) as total, 
			sum(finish_count) as success, 
			sum(undo_count) as undone, 
			sum(fail_count) as fail 
			FROM %s
			where 1`
	sql = fmt.Sprintf("%s and platform='spider_raw'", sql)
	if filter.ShowType == 1 {
		sql = fmt.Sprintf(sql, "concat(date, ' ', hour) as date,", m.TableName())
	} else if filter.ShowType == 0 {
		sql = fmt.Sprintf(sql, "concat(date, '日') as date,", m.TableName())
	}
	if filter.Project != "all" {
		sql = fmt.Sprintf("%s and table_name='%s'", sql, filter.Project)
	} else if filter.PlatForm != "all" {
		sql = fmt.Sprintf("%s and table_name like '%s%%'", sql, filter.PlatForm)
	}
	sql = fmt.Sprintf("%s %s", sql, m.WarpperSql(filter))
	return m.Query(sql)
}

func (m JobMonitorOrm) GetAnalyseDatas(filter req.Filter, cols ...string) ([]map[string][]byte, error) {
	colsStr := "*"
	if len(cols) > 0 {
		colsStr = strings.Join(cols, ",")
	}
	sql := fmt.Sprintf("select %s from %s where 1", colsStr, m.TableName())
	// 打包过滤条件
	if filter.PlatForm != "all" {
		sql = fmt.Sprintf("%s and platform != 'spider_raw'", sql)
	} else {
		sql = fmt.Sprintf("%s and platform = '%s'", sql, filter.PlatForm)
		if filter.Project != "all" {
			sql = fmt.Sprintf("%s and project = '%s'", sql, filter.Project[len(filter.PlatForm)+1:])
		}
	}
	sql = fmt.Sprintf("%s and %s", sql, m.WarpperSql(filter))
	return m.Query(sql)
}
