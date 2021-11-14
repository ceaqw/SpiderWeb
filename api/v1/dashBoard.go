package v1

import (
	"SpiderWeb/models"
	"SpiderWeb/services/resp"

	"github.com/gin-gonic/gin"
)

type DashBoard struct {
	models.CrawlerInformationOrm
}

func (a DashBoard) List(c *gin.Context) {

}

func (a DashBoard) ProjectList(c *gin.Context) {
	result, err := a.GetProjectListWithPlatform()
	if err != nil {
		c.JSON(200, resp.ErrorResp(500, "获取project错误"))
	}
	c.JSON(200, resp.Success(result))
}
