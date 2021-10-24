package main

import (
	_ "SpiderWeb/conf"
	"SpiderWeb/models"
	"fmt"
)

func main() {
	baseDb := models.BaseDB{}
	results, err := baseDb.Query("select * from course")
	if err == nil {
		fmt.Println(string(results[0]["cname"]))
	}
}
