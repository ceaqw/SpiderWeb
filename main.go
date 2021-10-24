package main

import (
	"SpiderWeb/conf"
	"SpiderWeb/router"
	"fmt"
)

func main() {
	appCfg := conf.GetAppCfg()
	r := router.Init()
	r.Run(fmt.Sprintf("%s:%s", appCfg.Host, appCfg.Port))
}
