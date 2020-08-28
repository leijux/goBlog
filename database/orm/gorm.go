package orm

import (
	"goBlog/config"
	"goBlog/log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func init() {
	var err error
	DriverName := config.GetString("database.mysql.driverName")
	DataSourceName := config.GetString("database.mysql.dataSourceName")
	Db, err = gorm.Open(DriverName, DataSourceName)
	Db.SingularTable(true)//不加s
	//Db.SetLogger(log.Logger)
	Db.LogMode(false)
	if err != nil {
		log.Logger.Fatalln(err)
	}
}

func Close() {
	if Db != nil {
		Db.Close()
	}
}
