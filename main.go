package main

import (
	"SpiderWeb/conf"
	"SpiderWeb/router"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	appCfg := conf.GetAppCfg()
	gin.SetMode("debug")
	r := router.Init()
	err := r.Run(fmt.Sprintf("%s:%s", appCfg.Host, appCfg.Port))

	if err != nil {
		log.Fatalf("Start server: %+v", err)
	} else {
		log.Fatalf("Start server: success")
	}
}
