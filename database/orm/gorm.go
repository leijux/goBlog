package orm

import (
	"time"

	"goBlog/config"
	"goBlog/log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func init() {
	var err error

	//DriverName := config.GetString("database.mysql.driverName")
	var dialector gorm.Dialector
	switch config.GetString("database.enable") {
	case config.Mysql:
		DataSourceName := config.GetString("database.mysql.dataSourceName")
		dialector = mysql.Open(DataSourceName)
	case config.Sqlite:
		DataSourceName := config.GetString("database.sqlite.dataSourceName")
		dialector = sqlite.Open(DataSourceName)
	default:
		log.Fatalln("db err")
	}
	ormLogger := logger.New(
		log.NewStdLog(),
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // 禁用彩色打印},
		})
	ormConfig := &gorm.Config{
		//禁用 事务
		SkipDefaultTransaction: true,
		//创建 prepared statement 并缓存，可以提高后续的调用速度
		PrepareStmt: true,
		//禁用外键
		DisableForeignKeyConstraintWhenMigrating: true,

		Logger: ormLogger,
	}
	Db, err = gorm.Open(dialector, ormConfig)

	// Db.SingularTable(true) //不加s
	// Db.SetLogger(log.Logger)
	// Db.LogMode(false)
	if err != nil {
		log.Fatalln(err)
	}
}
