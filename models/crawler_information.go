/*
 * @Date: 2021-12-09 14:00:10
 * @LastEditTime: 2021-12-09 14:57:21
 * @Author: ceaqw
 */
package models

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
