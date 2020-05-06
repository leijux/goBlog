package database

import (
	"task-system/config"
	myerr "task-system/err"
	"task-system/log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)
//Db 数据库链接对象
var Db *sqlx.DB

func init() {
	var err error
	switch config.Cfg.Database.Enable {
	case config.Mysql:
		DriverName := config.Cfg.Database.Mysql.DriverName
		DataSourceName := config.Cfg.Database.Mysql.DataSourceName
		Db, err = sqlx.Connect(DriverName, DataSourceName)
	case config.Sqlite:
		DriverName := config.Cfg.Database.Sqlite.DriverName
		DataSourceName := config.Cfg.Database.Sqlite.DataSourceName
		Db, err = sqlx.Connect(DriverName, DataSourceName)
	default:
		log.Logger.Fatalln(myerr.ErrEnableValue)
	}
	if err != nil {
		log.Logger.Fatalln(err)
	}
	Db.SetMaxIdleConns(10) //设置连接池中的保持连接的最大连接数
}
