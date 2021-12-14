/*
 * @Date: 2021-12-09 14:00:10
 * @LastEditTime: 2021-12-14 13:34:25
 * @Author: ceaqw
 */
package models

import (
	"SpiderWeb/models/response"
)

type CrawlerInformation struct {
	Id          uint64 `xorm:"pk autoincr" json:"id"`
	Project     string `xorm:"varchar(255) not null" json:"project"`
	Platform    string `xorm:"varchar(45) not null" json:"platform"`
	ScriptId    uint64 `xorm:"not null" json:"script_id"`
	Ip          string `xorm:"varchar(45) not null" json:"ip"`
	Comment     string `xorm:"text" json:"comment"`
	Server      uint8  `xorm:"not null" json:"server"`
	CriticalKpi string `xorm:"varchar(200)" json:"critical_kpi"`
	BindTable   string `xorm:"varchar(200)" json:"bind_table"`
	DelFlag     uint8  `json:"del_flag"`
}

func (CrawlerInformation) TableName() string {
	return "crawler_information"
}

type CrawlerInformationOrm struct {
	MainDbHand
	CrawlerInformation
}

func (m CrawlerInformationOrm) Insert(crawlerInformations ...CrawlerInformation) (int64, error) {
	infoList := make([]interface{}, 0)
	for _, crawlerInformation := range crawlerInformations {
		crawlerInformation.DelFlag = 0
		infoList = append(infoList, crawlerInformation)
	}
	return MainSqlDb.Insert(infoList...)
}

func (m CrawlerInformationOrm) Update(crawlerInformation CrawlerInformation) (int64, error) {
	crawlerInformation.DelFlag = 0
	return MainSqlDb.Id(crawlerInformation.Id).Update(crawlerInformation)
}

func (m CrawlerInformationOrm) GetProjectInfos(offset int, limit int, query string, args ...interface{}) (int64, []response.CrawlerInformation, error) {
	result := make([]response.CrawlerInformation, 0)
	total, err := MainSqlDb.Where(query, args...).Limit(limit, offset).FindAndCount(&result)
	return total, result, err
}
