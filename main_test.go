package main

import (
	"log"
	"testing"

	"task-system/config"
	"task-system/database"
	"task-system/models/user"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"xorm.io/xorm"
)

func Test_DbPing(t *testing.T) {
	err := database.Db.Ping()
	if err != nil {
		assert.Error(t, err, "发生错误")
	}

}
func Test_createTable(t *testing.T) {
	engine, err := xorm.NewEngine(config.Cfg.Database.Mysql.DriverName, config.Cfg.Database.Mysql.DataSourceName)
	if err != nil {
		log.Println(err)
	}
	err = engine.Sync2(new(user.User))
	if err != nil {
		log.Println(err)
	}

	// i, err := engine.Insert(&user.User{
	// 	Name:      "leiju",
	// 	Emeil:     "leiju@outlook.com",
	// 	Password:  "12345678",
	// 	Authority: 0,
	// })
	// if err != nil {
	// 	log.Println(i, err)
	// }
}

// func Test_Redis(t *testing.T) {
// 	s := cache.Redisdb.Ping().String()
// 	log.Println(s)
// }
