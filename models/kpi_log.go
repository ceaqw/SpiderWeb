/*
 * @Date: 2021-12-09 14:26:24
 * @LastEditTime: 2021-12-10 12:39:18
 * @Author: ceaqw
 */
package models

import (
	"SpiderWeb/models/req"
	"fmt"
	"strings"
)

// type KpiLog struct {

// }

type KpiLogOrm struct {
	SpiderDbHand
}

func (t KpiLogOrm) GetCriticalKpiDataByProjectAndP(filter req.Filter, projects map[string][]string) ([]map[string][]byte, error) {
	filter.StartDate = fmt.Sprintf("%s 00:00:00", filter.StartDate)
	filter.EndDate = fmt.Sprintf("%s 23:59:59", filter.EndDate)

	sql := "select project, p, count(*) as num from kpi_log where 1 %s group by project, p order by project, num desc"
	filterStr := ""
	filterStr = fmt.Sprintf("%s and date >='%s' and date<='%s'", filterStr, filter.StartDate, filter.EndDate)

	// 添加project&p过滤信息
	projectFilters := make([]string, 0)
	for project, ps := range projects {
		projectStr := fmt.Sprintf("(project=%s and %s)", project, "%s")
		if len(ps) > 0 {
			pInStr := fmt.Sprintf("p in (%s)", strings.Join(ps, ","))
			projectStr = fmt.Sprintf(projectStr, pInStr)
		} else {
			projectStr = fmt.Sprintf(projectStr, "1")
		}
		projectFilters = append(projectFilters, projectStr)
	}
	if len(projectFilters) > 0 {
		filterStr = fmt.Sprintf("%s and %s", filterStr, fmt.Sprintf("(%s)", strings.Join(projectFilters, ",")))
	}
	sql = fmt.Sprintf(sql, filterStr)
	return SpiderDb.Query(sql)
}
