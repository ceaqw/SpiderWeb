/*
 * @Date: 2021-11-16 09:08:27
 * @LastEditTime: 2021-12-10 15:32:47
 * @Author: ceaqw
 */
package models

import (
	"SpiderWeb/conf"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	MainSqlDb, Db, SpiderDb, DbInternet, DbCn, DbWeb *xorm.Engine
)

type MainDbHand struct {
}

type SpiderDbHand struct {
}

func init() {
	InitMainDb()
	InitSpiderDb()
}

func InitMainDb() {
	mainSqlDbCfg := conf.GetDbCfgByName("maindb")
	jdbc := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		mainSqlDbCfg.UserName, mainSqlDbCfg.Password, mainSqlDbCfg.Host, mainSqlDbCfg.Port, mainSqlDbCfg.Database,
	)
	var err error
	MainSqlDb, err = xorm.NewEngine(mainSqlDbCfg.Driver, jdbc)

	if err != nil {
		panic(fmt.Sprintf("maindb error: %#v\n", err.Error()))
	}
	err = MainSqlDb.Ping()
	if err != nil {
		panic(fmt.Sprintf("maindb connect error: %#v\n", err.Error()))
	}
	// MainSqlDb.ShowSQL(true)
	// err = Main	SqlDb.Sync(
	// 	new(Role),
	// 	new(User),
	// 	new(CrawlerInformation),
	// 	new(CrawlerStat),
	// 	new(Project),
	// )
	// if err != nil {
	// 	panic(fmt.Sprintf("maindb sync error: %#v\n", err.Error()))
	// }
}

func InitSpiderDb() {
	spiderDbCfg := conf.GetDbCfgByName("spiderdb")
	jdbc := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		spiderDbCfg.UserName, spiderDbCfg.Password, spiderDbCfg.Host, spiderDbCfg.Port, spiderDbCfg.Database,
	)
	var err error
	SpiderDb, err = xorm.NewEngine(spiderDbCfg.Driver, jdbc)

	if err != nil {
		panic(fmt.Sprintf("spiderdb error: %#v\n", err.Error()))
	}
	err = SpiderDb.Ping()
	if err != nil {
		panic(fmt.Sprintf("spiderdb connect error: %#v\n", err.Error()))
	}
}

func (MainDbHand) Query(sql string) ([]map[string][]byte, error) {
	results, err := MainSqlDb.Query(sql)
	return results, err
}

func (MainDbHand) Excute(sql string, arg ...interface{}) error {
	_, err := MainSqlDb.Exec(sql)
	return err
}
