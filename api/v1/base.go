package v1

import (
	"SpiderWeb/models"
	"SpiderWeb/services/resp"

	"github.com/gin-gonic/gin"
)

type Base struct {
	models.CrawlerInformationOrm
}

func (a Base) ProjectList(c *gin.Context) {
	respDatas := make(map[string][]string)
	result, err := a.GetProjectListWithPlatform()
	if err != nil {
		c.JSON(200, resp.ErrorResp(500, "获取project错误"))
	} else {
		for _, row := range result {
			respDatas[string(row["platform"])] = append(respDatas[string(row["platform"])], string(row["project"]))
		}
		c.JSON(200, resp.Success(respDatas))
	}
}
