package services

import (
	"SpiderWeb/models"
	"SpiderWeb/models/req"
	"strconv"
)

type CrawlerStatService struct {
	models.CrawlerStatOrm
}

func (s CrawlerStatService) GetTopError(filter req.Filter) (map[string]map[string]interface{}, int, error) {
	req.FilterVerify(&filter)
	result := make(map[string]map[string]interface{})
	// 先获取平台项目信息
	projectInfo, err := s.GetPRojectInfo(
		filter,
		`ci.platform as platform, 
		ci.project as project, 
		p.full_name as full_name`,
	)
	if err != nil {
		return nil, 0, err
	}
	// 根据获取的项目列表获取任务进度
	projectList := s.ArrayUtil.GetColumnByKey(projectInfo, "project")
	crawerInfo, err := s.GetCrawlerInfoByProject(projectList, "project, sum(fail) as fail, sum(undone) as undone, sum(success) as success", filter)
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
