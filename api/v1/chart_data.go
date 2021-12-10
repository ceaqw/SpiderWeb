/*
 * @Date: 2021-11-29 09:13:10
 * @LastEditTime: 2021-12-10 14:39:16
 * @Author: ceaqw
 */
package v1

import (
	"SpiderWeb/models"
	"SpiderWeb/models/req"
	"SpiderWeb/models/response"
	"SpiderWeb/services"
	"SpiderWeb/services/resp"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type ChartData struct {
	response.CrawlerStat
	models.CrawlerStatOrm
	models.JobMonitorOrm
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
		// 时间特殊处理
		req.FilterVerify(&filter)
		result, err := a.GetAllDatas(filter)
		if filter.ShowType == 1 && filter.EndDate == time.Now().Format("2006-01-02") && (filter.PlatForm == "all" || filter.Project == "all") {
			result = result[:len(result)-1]
		}
		if err == nil {
			response := a.PackChartDatasByCrawler(result, filter)
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
		req.FilterVerify(&filter)
		result, err := a.GetDoDatas(filter)
		if err == nil {
			response := a.PackChartDatasByAnalyse(result)
			response["title"] = "spider_raw:" + filter.PlatForm + "_" + filter.Project
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

func (a ChartData) ProjectDatas(c *gin.Context) {
	filter := req.Filter{}
	if err := c.BindJSON(&filter); err == nil {
		result, err := a.GetDoDatas(filter)
		if err == nil {
			response := a.PackChartDatasByAnalyse(result)
			c.JSON(200, resp.Success(response))
		} else {
			resp.Error(c, "查询出错")
		}
	} else {
		c.JSON(200, resp.ErrorResp(400, "参数错误"))
	}
}
