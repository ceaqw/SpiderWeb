/*
 * @Date: 2021-11-16 09:08:27
 * @LastEditTime: 2021-12-07 12:46:25
 * @Author: ceaqw
 */
package response

import (
	"SpiderWeb/models/req"
	"strconv"
)

type CrawlerStat struct {
}

func (m CrawlerStat) PackChartDatasByCrawler(result []map[string][]byte, filter req.Filter) map[string]interface{} {
	response := make(map[string]interface{})
	barDatas := make([]map[string]interface{}, 0)
	var successCom, failCom, undoneCom, retryCom, totalCom int
	for _, item := range result {
		var success, fail, retry, undone, total int
		sub := make(map[string]interface{})
		success, _ = strconv.Atoi(string(item["success"]))
		fail, _ = strconv.Atoi(string(item["fail"]))
		retry, _ = strconv.Atoi(string(item["retry"]))
		undone, _ = strconv.Atoi(string(item["undone"]))
		total, _ = strconv.Atoi(string(item["total"]))
		if filter.ShowType == 1 {
			sub["success"] = success - successCom
			sub["fail"] = fail - failCom
			sub["undone"] = undone
			sub["retry"] = retry - retryCom
			sub["total"] = total - totalCom
			retryCom = retry
			successCom = success
			failCom = fail
			undoneCom = undone
			totalCom = total
		} else {
			sub["success"] = success
			sub["fail"] = fail
			sub["undone"] = undone
			sub["retry"] = retry
			sub["total"] = total
			retryCom = retryCom + retry
			successCom = successCom + success
			failCom = failCom + fail
			undoneCom = undoneCom + undone
			totalCom = totalCom + total
		}
		sub["date"] = string(item["date"])
		barDatas = append(barDatas, sub)
	}
	response["pieDatas"] = []int{successCom, failCom, undoneCom}
	response["barDatas"] = barDatas
	return response
}

func (m CrawlerStat) PackChartDatasByAnalyse(result []map[string][]byte) map[string]interface{} {
	response := make(map[string]interface{})
	barDatas := make([]map[string]interface{}, 0)
	var success, fail, undone int
	for _, item := range result {
		sub := make(map[string]interface{})
		sub["success"], _ = strconv.Atoi(string(item["success"]))
		sub["fail"], _ = strconv.Atoi(string(item["fail"]))
		sub["undone"], _ = strconv.Atoi(string(item["undone"]))
		sub["retry"], _ = strconv.Atoi(string(item["retry"]))
		sub["total"], _ = strconv.Atoi(string(item["total"]))
		sub["date"] = string(item["date"])
		barDatas = append(barDatas, sub)
		success += sub["success"].(int)
		fail += sub["fail"].(int)
		undone += sub["undone"].(int)
	}
	response["pieDatas"] = []int{success, fail, undone}
	response["barDatas"] = barDatas
	return response
}
