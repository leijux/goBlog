package main

import (
	"log"
	"task-system/models/user"
	"testing"

	"task-system/config"
	"task-system/database/mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"xorm.io/xorm"
)

func Test_DbPing(t *testing.T) {
	err := mysql.Db.Ping()
	if err != nil {
		assert.Error(t, err, "发生错误")
	}

}
func Test_createTable(t *testing.T) {
	engine, err := xorm.NewEngine(config.Cfg.Sql.Mysql.DriverName, config.Cfg.Sql.Mysql.DataSourceName)
	if err != nil {
		log.Println(err)
	}
	err = engine.Sync2(new(user.User))
	if err != nil {
		log.Println(err)
	}

	i,err:=engine.Insert(&user.User{
		Name: "leiju",
		Emeil: "leiju@outlook.com",
		Password: "12345678",
		Authority: 0,
	})
	if err!=nil{
		log.Println(i,err)
	}
}
