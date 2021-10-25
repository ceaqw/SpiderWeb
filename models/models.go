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

type BaseDB struct {
}

func init() {
	mainSqlDbCfg := conf.GetMainDbCfg()
	jdbc := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		mainSqlDbCfg.UserName, mainSqlDbCfg.Password, mainSqlDbCfg.Host, mainSqlDbCfg.Port, mainSqlDbCfg.Database,
	)
	var err error
	MainSqlDb, err = xorm.NewEngine(mainSqlDbCfg.Driver, jdbc)

	if err != nil {
		panic(fmt.Sprintf("db error: %#v\n", err.Error()))
	}
	err = MainSqlDb.Ping()
	if err != nil {
		panic(fmt.Sprintf("db connect error: %#v\n", err.Error()))
	}
	// MainSqlDb.ShowSQL(true)
}

func (baseDB *BaseDB) Query(sql string) ([]map[string][]byte, error) {
	results, err := MainSqlDb.Query(sql)
	return results, err
}
