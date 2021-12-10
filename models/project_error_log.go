/*
 * @Date: 2021-12-10 14:38:28
 * @LastEditTime: 2021-12-10 15:16:00
 * @Author: ceaqw
 */
/*
 * @Date: 2021-12-10 14:38:28
 * @LastEditTime: 2021-12-10 15:11:23
 * @Author: ceaqw
 */
package models

import (
	"SpiderWeb/models/req"
	"fmt"
	"strings"
	"time"
)

type ProjectErrorLog struct {
	Platform     string    `xorm:"varchar(45) notnull 'platform'" json:"platform"`      //
	TableNames   string    `xorm:"varchar(100) notnull 'table_name'" json:"table_name"` //
	Date         time.Time `xorm:"date notnull 'date'" json:"date"`                     //
	ErrorContent string    `xorm:"longtext  'error_content'" json:"error_content"`      //
	H0           int       `xorm:"int(10) default 0 '0'" json:"0"`                      //
	H1           int       `xorm:"int(10) default 0 '1'" json:"1"`                      //
	H2           int       `xorm:"int(10) default 0 '2'" json:"2"`                      //
	H3           int       `xorm:"int(10) default 0 '3'" json:"3"`                      //
	H4           int       `xorm:"int(10) default 0 '4'" json:"4"`                      //
	H5           int       `xorm:"int(10) default 0 '5'" json:"5"`                      //
	H6           int       `xorm:"int(10) default 0 '6'" json:"6"`                      //
	H7           int       `xorm:"int(10) default 0 '7'" json:"7"`                      //
	H8           int       `xorm:"int(10) default 0 '8'" json:"8"`                      //
	H9           int       `xorm:"int(10) default 0 '9'" json:"9"`                      //
	H10          int       `xorm:"int(10) default 0 '10'" json:"10"`                    //
	H11          int       `xorm:"int(10) default 0 '11'" json:"11"`                    //
	H12          int       `xorm:"int(10) default 0 '12'" json:"12"`                    //
	H13          int       `xorm:"int(10) default 0 '13'" json:"13"`                    //
	H14          int       `xorm:"int(10) default 0 '14'" json:"14"`                    //
	H15          int       `xorm:"int(10) default 0 '15'" json:"15"`                    //
	H16          int       `xorm:"int(10) default 0 '16'" json:"16"`                    //
	H17          int       `xorm:"int(10) default 0 '17'" json:"17"`                    //
	H18          int       `xorm:"int(10) default 0 '18'" json:"18"`                    //
	H19          int       `xorm:"int(10) default 0 '19'" json:"19"`                    //
	H20          int       `xorm:"int(10) default 0 '20'" json:"20"`                    //
	H21          int       `xorm:"int(10) default 0 '21'" json:"21"`                    //
	H22          int       `xorm:"int(10) default 0 '22'" json:"22"`                    //
	H23          int       `xorm:"int(10) default 0 '23'" json:"23"`                    //
}

//TableName Get table name
func (t ProjectErrorLog) TableName() string {
	return "project_error_lo"
}

type ProjectErrorLogOrm struct {
	ProjectErrorLog
	MainDbHand
}

func (m ProjectErrorLogOrm) WarpperSql(filter req.Filter) string {
	filterStr := ""
	filterStr = fmt.Sprintf("%s and date >= '%s' and date <= '%s' order by %s.date", filterStr, filter.StartDate, filter.EndDate, m.TableName())
	return filterStr
}

func (m ProjectErrorLogOrm) GetErrDatasByProjects(filter req.Filter, type_ string, cols ...string) ([]map[string][]byte, error) {
	colsStr := "*"
	if len(cols) > 0 {
		colsStr = strings.Join(cols, ",")
	}
	sql := fmt.Sprintf("select %s from %s where 1", colsStr, m.TableName())
	// 打包过滤条件
	switch type_ {
	case "do":
		sql = fmt.Sprintf("%s and platform = 'spider_raw'", sql)
		sql = fmt.Sprintf("%s and table_name = '%s_do'", sql, filter.Project)
	case "analyse":
		sql = fmt.Sprintf("%s and platform = 'spider_raw'", sql)
		sql = fmt.Sprintf("%s and table_name = '%s_analyse'", sql, filter.Project)
	case "script":
		sql = fmt.Sprintf("%s and platform = '%s'", sql, filter.PlatForm)
		sql = fmt.Sprintf("%s and table_name = '%s'", sql, filter.Project[len(filter.PlatForm)+1:])
	}
	sql = fmt.Sprintf("%s and %s", sql, m.WarpperSql(filter))
	return m.Query(sql)
}

func (m ProjectErrorLogOrm) GetDoErrDatas(filter req.Filter, cols ...string) ([]map[string][]byte, error) {
	return m.GetErrDatasByProjects(filter, "do", cols...)
}
