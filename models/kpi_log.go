/*
 * @Date: 2021-12-09 14:26:24
 * @LastEditTime: 2021-12-09 17:16:31
 * @Author: ceaqw
 */
package models

import (
	"SpiderWeb/models/req"
	"fmt"
)

// type KpiLog struct {

// }

type KpiLogOrm struct {
	SpiderDbHand
}

func (t KpiLogOrm) CriticalKpiDataByProjectAndP(filter req.Filter, cols string, projects map[string][]string) ([]map[string][]byte, error) {
	filter.StartDate = fmt.Sprintf("%s 00:00:00", filter.StartDate)
	filter.EndDate = fmt.Sprintf("%s 23:59:59", filter.EndDate)

	sql := "select project, p, count(*) as num from kpi_log where 1 %s group by project, p order by project, num desc"
	filterStr := ""
	filterStr = fmt.Sprintf("%s and date >='%s' and date<='%s'", filterStr, filter.StartDate, filter.EndDate)

	// 添加project&p过滤信息
	filterStr = fmt.Sprintf("%s and (")
	// for project, ps := range projects {
	// 	projectStr := fmt.Sprintf("(project=%s", project)
	// 	for _, p := range ps {

	// 	}
	// }
	sql = fmt.Sprintf(sql, filterStr)
	return SpiderDb.Query(sql)
}
