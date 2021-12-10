/*
 * @Date: 2021-11-22 09:49:48
 * @LastEditTime: 2021-12-10 16:41:22
 * @Author: ceaqw
 */
package services

import (
	"SpiderWeb/models"
	"SpiderWeb/models/req"
	"SpiderWeb/modules/utils"
	"strconv"
)

type CrawlerStatService struct {
	CrawlerStatOrm models.CrawlerStatOrm
	ProjectOrm     models.ProjectOrm
	ArrayUtil      utils.Array
}

func (s CrawlerStatService) GetTopError(filter req.Filter) (map[string]map[string]interface{}, int, error) {
	req.FilterVerify(&filter)
	result := make(map[string]map[string]interface{})
	// 先获取平台项目信息
	projectInfo, err := s.ProjectOrm.GetProjectInfo(
		filter,
		"ci.platform as platform",
		"ci.project as project",
		"p.full_name as full_name",
	)
	if err != nil {
		return nil, 0, err
	}
	// 根据获取的项目列表获取任务进度
	projectList := s.ArrayUtil.GetColumnByKey(projectInfo, "project")
	filter.ShowType = 0
	crawerInfo, err := s.CrawlerStatOrm.GetCrawlerStatByProject(projectList, "project, sum(fail) as fail, sum(undone) as undone, sum(success) as success", filter)
	if err != nil {
		return nil, 0, err
	}
	projectMap := s.ArrayUtil.GenerateMapByKey(projectInfo, "project")
	var count int
	addPk := make([]string, 0)
	for _, item := range crawerInfo {
		pk := string(item["project"])
		if s.ArrayUtil.In(pk, addPk) {
			tmpFail, _ := strconv.Atoi(string(item["fail"]))
			tmpUndone, _ := strconv.Atoi(string(item["undone"]))
			tmpSuccess, _ := strconv.Atoi(string(item["success"]))
			result[pk]["fail"] = tmpFail + result[pk]["fail"].(int)
			result[pk]["undone"] = tmpUndone + result[pk]["undone"].(int)
			result[pk]["success"] = tmpSuccess + result[pk]["success"].(int)
		} else {
			count += 1
			addPk = append(addPk, pk)
			tmp := make(map[string]interface{})
			tmp["project"] = pk
			tmp["fullName"] = string(projectMap[pk]["full_name"])
			tmp["platform"] = string(projectMap[pk]["platform"])
			tmp["fail"], _ = strconv.Atoi(string(item["fail"]))
			tmp["undone"], _ = strconv.Atoi(string(item["undone"]))
			tmp["success"], _ = strconv.Atoi(string(item["success"]))
			result[pk] = tmp
		}
	}
	return result, count, nil
}
