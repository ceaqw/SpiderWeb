package response

import "strconv"

type CrawlerStat struct {
}

func (m CrawlerStat) PackChartDatas(result []map[string][]byte) map[string]interface{} {
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
