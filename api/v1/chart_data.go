package v1

import (
	"SpiderWeb/models"
	"SpiderWeb/models/req"
	"SpiderWeb/models/response"
	"SpiderWeb/services/resp"

	"github.com/gin-gonic/gin"
)

type ChartData struct {
	models.CrawlerStatOrm
	response.CrawlerStat
}

func (a ChartData) AllChartData(c *gin.Context) {
	filter := req.Filter{}
	if err := c.BindJSON(&filter); err == nil {
		result, err := a.GetAllDatas(filter)
		if err == nil {
			response := a.PackChartDatas(result)
			response["title"] = filter.PlatForm + "_" + filter.Project
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
			response := a.PackChartDatas(result)
			response["title"] = filter.PlatForm + "_" + filter.Project
			c.JSON(200, resp.Success(response))
		} else {
			resp.Error(c, "查询出错")
		}
	} else {
		c.JSON(200, resp.ErrorResp(400, "参数错误"))
	}
}
