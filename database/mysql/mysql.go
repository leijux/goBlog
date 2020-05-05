package mysql

import (
	"task-system/config"
	"task-system/log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	var err error
	Db, err = sqlx.Connect(config.Cfg.Sql.Mysql.DriverName, config.Cfg.Sql.Mysql.DataSourceName)
	if err != nil {
		log.Logger.Fatalln(err)
	}
	Db.SetMaxIdleConns(10)//设置连接池中的保持连接的最大连接数
}
