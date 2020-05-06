package database

import (
	"task-system/config"
	myerr "task-system/err"
	"task-system/log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	var err error

	switch config.Cfg.Database.Enable {
	case config.Mysql:
		Db, err = sqlx.Connect(config.Cfg.Database.Mysql.DriverName, config.Cfg.Database.Mysql.DataSourceName)
	case config.Sqlite:
		Db, err = sqlx.Connect(config.Cfg.Database.Sqlite.DriverName, config.Cfg.Database.Sqlite.DataSourceName)
	default:
		log.Logger.Fatalln(myerr.ErrEnableValue)
	}
	if err != nil {
		log.Logger.Fatalln(err)
	}
	Db.SetMaxIdleConns(10) //设置连接池中的保持连接的最大连接数
}
