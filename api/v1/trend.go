package v1

import "github.com/gin-gonic/gin"

type Trend struct {
	Kpi
	Fail
}

type Kpi struct {
}

type Fail struct {
}

func (Kpi) List(c *gin.Context) {

}

func (Fail) List(c *gin.Context) {

}
