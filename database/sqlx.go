package database

import (
	"goBlog/config"
	myerr "goBlog/err"
	"goBlog/log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jmoiron/sqlx"
)

//Db 数据库链接对象
var Db *sqlx.DB

func init() {
	var err error
	switch config.GetString("database.enable") {
	case config.Mysql:
		DriverName := config.GetString("database.mysql.driverName")
		DataSourceName := config.GetString("database.mysql.dataSourceName")
		Db, err = sqlx.Connect(DriverName, DataSourceName)
	case config.Sqlite:
		DriverName := config.GetString("database.sqlite.driverName")
		DataSourceName := config.GetString("database.sqlite.dataSourceName")
		Db, err = sqlx.Connect(DriverName, DataSourceName)
	default:
		log.Logger.Fatalln(myerr.ErrEnableValue)
	}
	if err != nil {
		log.Logger.Fatalln(err)
	}
	Db.SetMaxIdleConns(10) //设置连接池中的保持连接的最大连接数
}
