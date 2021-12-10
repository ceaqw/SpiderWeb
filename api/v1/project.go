/*
 * @Date: 2021-10-25 17:05:54
 * @LastEditTime: 2021-12-10 17:25:17
 * @Author: ceaqw
 */
package v1

import (
	"SpiderWeb/models"
	"SpiderWeb/services/resp"

	"github.com/gin-gonic/gin"
)

type Project struct {
	models.CrawlerInformationOrm
	models.AllMongoProjectOrm
}

func (a Project) ProjectList(c *gin.Context) {
	respDatas := make(map[string][]string)
	result, _, err := a.GetProjectListWithPlatform()
	if err != nil {
		c.JSON(200, resp.ErrorResp(500, "获取project错误"))
	} else {
		for _, row := range result {
			respDatas[string(row["platform"])] = append(respDatas[string(row["platform"])], string(row["project"]))
		}
		c.JSON(200, resp.Success(respDatas))
	}
}

func (a Project) GetProjectInfos(c *gin.Context) {
	data := make(map[string]interface{})
	datas := make([]map[string]interface{}, 0)
	test := make(map[string]interface{})
	test["platform"] = "shopee"
	test["project"] = "item"
	test["script_id"] = 7456
	test["ip"] = "192.168.0.1"
	test["comment"] = "shopee_item"
	test["bind_table"] = "item"
	datas = append(datas, test)
	data["datas"] = datas
	data["totalNumber"] = 1
	c.JSON(200, resp.Success(data))
}

func (a Project) GetAllPlatformAndProjectMap(c *gin.Context) {
	rest, err := a.GetAllMongoProject()
	if err == nil {
		platformAndProjectMap := make(map[string][]string)
		for _, item := range rest {
			if _, ok := platformAndProjectMap[item.Platform]; ok {
				platformAndProjectMap[item.Platform] = append(platformAndProjectMap[item.Platform], item.Project)
			} else {
				platformAndProjectMap[item.Platform] = []string{item.Project}
			}
		}
		c.JSON(200, resp.Success(platformAndProjectMap))
	} else {
		c.JSON(500, resp.ErrorResp("查询错误"))
	}
}
