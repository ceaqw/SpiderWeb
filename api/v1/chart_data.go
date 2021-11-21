package v1

import (
	"SpiderWeb/models"
	"SpiderWeb/models/req"
	"SpiderWeb/models/response"
	"SpiderWeb/services"
	"SpiderWeb/services/resp"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ChartData struct {
	models.CrawlerStatOrm
	response.CrawlerStat
	services.CrawlerStatService
}

func (a ChartData) AllChartData(c *gin.Context) {
	filter := req.Filter{}
	platform, ok := c.Params.Get("platform")
	if ok {
		filter.PlatForm = platform
		filter.Project = "all"
	}
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

func (a ChartData) TopErrorDatas(c *gin.Context) {
	filter := req.Filter{}
	if err := c.BindJSON(&filter); err == nil {
		result, count, err := a.GetTopError(filter)
		response := make(map[string]interface{})
		if err == nil {
			response["tableDatas"] = result
			response["count"] = count
			response["title"] = fmt.Sprintf("%s_%s", filter.PlatForm, filter.Project)
			c.JSON(200, resp.Success(response))
		} else {
			resp.Error(c, "查询出错")
		}
	} else {
		c.JSON(200, resp.ErrorResp(400, "参数错误"))
	}
}
