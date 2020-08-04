package orm

import (
	"goBlog/config"
	"goBlog/log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func init() {
	DriverName := config.GetString("database.mysql.driverName")
	DataSourceName := config.GetString("database.mysql.dataSourceName")
	Db, err := gorm.Open(DriverName, DataSourceName)
	Db.SingularTable(true)
	if err != nil {
		log.Logger.Fatalln(err)
	}
}
func Close() {
	if Db != nil {
		Db.Close()
	}
}
