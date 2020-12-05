package orm

import (
	"goBlog/config"
	"goBlog/log"
	"time"

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

	Db, err = gorm.Open(dialector, &gorm.Config{
		Logger: logger.New(
			log.NewStdLog(),
			logger.Config{
				SlowThreshold: time.Second,   // 慢 SQL 阈值
				LogLevel:      logger.Silent, // Log level
				Colorful:      false,         // 禁用彩色打印},
			},
		),
		//禁用 事务
		SkipDefaultTransaction: true,
		//创建 prepared statement 并缓存，可以提高后续的调用速度
		PrepareStmt: true,
		//禁用外键
		// DisableForeignKeyConstraintWhenMigrating: true,
	})

	// Db.SingularTable(true) //不加s
	// //Db.SetLogger(log.Logger)
	// Db.LogMode(false)

	if err != nil {
		log.Fatalln(err)
	}
	// timeoutContext, _ := context.WithTimeout(context.Background(), time.Second)
	// Db = db.WithContext(timeoutContext)

	sqlDB, err := Db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

}

//TODO 带判断
func Close() {
	if Db != nil {
		sqlDB, err := Db.DB()
		if err != nil {
			log.Fatalln(err)
		}
		sqlDB.Close()
	}
}

func Create(value interface{}) *gorm.DB {
	return Db.Create(value)
}
