package myDb

import (
	"commons/nError"
	"commons/tools"
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	MysqlOnce     sync.Once
	MysqlInstance *gorm.DB
)

//GetMysqlClient 初始链接的时候需要传入参数 mysqlOption: 0user, 1pass, 2host, 3port, 4dbName, 5charset
func GetMysqlClient(mysqlOption ...string) *gorm.DB {
	MysqlOnce.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", mysqlOption[0], mysqlOption[1], mysqlOption[2], mysqlOption[3], mysqlOption[4], mysqlOption[5])
		tools.Log.Info(fmt.Sprintf("%s,%s", tools.RunFuncName(), dsn))
		db, err := gorm.Open("mysql", dsn)
		if err != nil {
			tools.Log.Error(err.Error())
			nError.Panic(nError.ErrorDbConnectFail, "ErrorDbConnectFail")
		}
		tools.Log.Info(tools.RunFuncName() + " Connect OK!")
		MysqlInstance = db
	})
	return MysqlInstance
}
