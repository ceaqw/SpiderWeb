package v1

import (
	"SpiderWeb/models"
	"SpiderWeb/models/req"
	"SpiderWeb/services/resp"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ChartData struct {
	models.CrawlerStatOrm
}

func (a ChartData) AllChartData(c *gin.Context) {
	filter := req.Filter{}
	if err := c.BindJSON(&filter); err == nil {
		result, err := a.GetAllDatas(filter)
		if err == nil {
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
			response["title"] = filter.PlatForm + "_" + filter.Project
			response["pieDatas"] = []int{success, fail, undone}
			response["barDatas"] = barDatas
			c.JSON(200, resp.Success(response))
		} else {
			c.JSON(200, resp.ErrorResp(500, "查询出错"))
		}
	} else {
		c.JSON(200, resp.ErrorResp(400, "参数错误"))
	}
}
