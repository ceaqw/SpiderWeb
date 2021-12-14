/*
 * @Date: 2021-10-25 17:05:54
 * @LastEditTime: 2021-12-14 13:49:53
 * @Author: ceaqw
 */
package v1

import (
	"SpiderWeb/models"
	"SpiderWeb/models/req"
	"SpiderWeb/services/resp"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Project struct {
	crawlerInformationOrm models.CrawlerInformationOrm
	models.AllMongoProjectOrm
}

func (a Project) ProjectList(c *gin.Context) {
	respDatas := make(map[string][]string)
	result, _, err := a.crawlerInformationOrm.GetProjectListWithPlatform()
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
	var offset, limit int
	platform := c.Query("platform")
	project := c.Query("project")
	if pageSize, ok := c.GetQuery("pageSize"); ok {
		limit, _ = strconv.Atoi(pageSize)
		if limit > 50 {
			limit = 50
		}
	}
	if page, ok := c.GetQuery("page"); ok {
		pageInt, _ := strconv.Atoi(page)
		offset = (pageInt - 1) * limit
	}

	query := "1"
	args := make([]interface{}, 0)
	if platform != "all" {
		query = fmt.Sprintf("%s and platform=?", query)
		args = append(args, platform)
	}
	if project != "all" {
		query = fmt.Sprintf("%s and project=?", query)
		args = append(args, project)
	}
	total, datas, err := a.crawlerInformationOrm.GetProjectInfos(offset, limit, query, args...)
	if err == nil {
		data["datas"] = datas
		data["totalNumber"] = total
		c.JSON(200, resp.Success(data))
	} else {
		resp.Error(c, "查询错误")
	}

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

func (a Project) CreateProject(c *gin.Context) {
	createProjectForm := req.CreateProjectForm{}
	if c.BindJSON(&createProjectForm) == nil {
		crawlerInformation := models.CrawlerInformation{
			Project:     createProjectForm.Project,
			Platform:    createProjectForm.Platform,
			ScriptId:    createProjectForm.ScriptId,
			Ip:          createProjectForm.Ip,
			Comment:     createProjectForm.Comment,
			Server:      createProjectForm.Server,
			CriticalKpi: createProjectForm.CriticalKpi,
			BindTable:   createProjectForm.BindTables,
		}
		var err error
		if createProjectForm.Id == 0 {
			// 创建
			_, err = a.crawlerInformationOrm.Insert(crawlerInformation)
		} else {
			// 更新
			crawlerInformation.Id = createProjectForm.Id
			_, err = a.crawlerInformationOrm.Update(crawlerInformation)
		}

		if err == nil {
			c.JSON(200, resp.Success(nil, "ok"))
		} else {
			fmt.Println(err)
			c.JSON(200, resp.Success(nil, "fail"))
		}
	} else {
		resp.Error(c, "参数错误")
	}
}
