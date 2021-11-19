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
			resp.Error(c, "查询出错")
		}
	} else {
		c.JSON(200, resp.ErrorResp(400, "参数错误"))
	}
}

func (a ChartData) AnalyseDatas(c *gin.Context) {
	filter := req.Filter{}
	if err := c.BindJSON(&filter); err == nil {
		result, err := a.GetAnalyseDatas(filter)
		if err == nil {
			response := make(map[string]interface{})
			response["title"] = filter.PlatForm + "_" + filter.Project
			response["pieDatas"] = []int{1, 2, 3}
			barDatas := make([]map[string]interface{}, 0)
			sub := make(map[string]interface{})
			sub["date"] = "2021-11-12日"
			sub["fail"] = 35575306
			sub["retry"] = 7612716
			sub["success"] = 190741941
			sub["total"] = 3162624
			sub["undone"] = 26902616
			barDatas = append(barDatas, sub)
			response["barDatas"] = barDatas
			c.JSON(200, resp.Success(response))
		} else {
			resp.Error(c, "查询出错")
		}
	} else {
		c.JSON(200, resp.ErrorResp(400, "参数错误"))
	}
}
