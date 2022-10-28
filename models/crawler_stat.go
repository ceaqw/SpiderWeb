package models

import (
	"SpiderWeb/models/req"
	"SpiderWeb/modules/utils"
	"errors"
	"fmt"
	"time"
)

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
	MainDbHand
	ArrayUtil utils.Array
}

func (m CrawlerInformationOrm) GetProjectListWithPlatform(platform ...string) ([]map[string][]byte, int, error) {
	var count int64
	sql := "select platform, project from crawler_information where del_flag <> 1 %s order by platform, project"
	if len(platform) > 0 {
		sql = fmt.Sprintf(sql, fmt.Sprintf(" and platform in (%s) ", ))
	} else {
		sql = fmt.Sprintf(sql, " and 1 ")
	}
	count, _ = MainSqlDb.Where("del_flag <> 1").Count(new(CrawlerInformation))
	result, err := m.Query(sql)
	return result, int(count), err
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

	sql = fmt.Sprintf("%s and cs.date >= '%s' and cs.date <= '%s' ", sql, filter.StartDate, filter.EndDate)
	// [天， 小时]
	if filter.ShowType == 0 {
		dateSelect := "(cs.`hour` = 23"
		toDay := time.Now().Format("2006-01-02")
		if filter.EndDate >= toDay {
			dateSelect = fmt.Sprintf("%s or (cs.`date` = '%s' and cs.hour = %d))", dateSelect, toDay, time.Now().Hour()-1)
		} else {
			dateSelect = fmt.Sprintf("%s)", dateSelect)
		}
		sql = fmt.Sprintf("%s and %s group by cs.date order by cs.date", sql, dateSelect)
		sql = fmt.Sprintf(sql, "concat(cs.date, '日') as date ")
	} else if filter.ShowType == 1 {
		sql = fmt.Sprintf("%s group by cs.date, cs.hour order by cs.date, cs.hour", sql)
		sql = fmt.Sprintf(sql, "concat(cs.date, ' ', cs.hour) as date ")
	}
	result, err := m.Query(sql)
	return result, err
}

func (m CrawlerStatOrm) GetCrawlerStatByProject(projectList [][]byte, cols string, filter req.Filter) ([]map[string][]byte, error) {
	sql := `select 
			%s
			from crawler_stat cs
			where project in ('%s')`
	sql = fmt.Sprintf(sql, cols, m.ArrayUtil.Join(projectList, "','"))
	sql = fmt.Sprintf("%s and date >= '%s' and date <= '%s'", sql, filter.StartDate, filter.EndDate)
	result := make([]map[string][]byte, 0)
	if filter.ShowType == 0 {
		dateSelect := "(cs.`hour` = 23"
		toDay := time.Now().Format("2006-01-02")
		if filter.EndDate >= toDay {
			dateSelect = fmt.Sprintf("%s or (cs.`date` = '%s' and cs.hour = %d))", dateSelect, toDay, time.Now().Hour()-1)
		} else {
			dateSelect = fmt.Sprintf("%s)", dateSelect)
		}
		sql = fmt.Sprintf("%s and %s group by cs.project, cs.date", sql, dateSelect)
	} else if filter.ShowType == 1 {
		sql = fmt.Sprintf("%s group by cs.project,cs.date,cs.hour", sql)
	} else {
		return result, errors.New("invalid showType")
	}
	sql = fmt.Sprintf("%s order by fail desc", sql)
	result, err := m.Query(sql)
	return result, err
}
