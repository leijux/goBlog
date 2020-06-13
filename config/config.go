package config

import (
	myerr "goBlog/err"

	"github.com/golobby/config"
	"github.com/golobby/config/feeder"
	log "github.com/sirupsen/logrus"
)

const (
	//JSONPath 文件路径
	JSONPath = "./config/config.json"
	//Mysql 字段
	Mysql = "mysql"
	//Sqlite 字段
	Sqlite = "sqlite"
)

//cfg 配置文件对象
var cfg *config.Config

func init() {
	var err error
	cfg, err = config.New(config.Options{
		Feeder: feeder.Json{Path: JSONPath},
	})
	if err != nil {
		log.Fatalln(myerr.ErrOpenFile, err)
	}
}

//GetString 返回字符串
func GetString(s string) string {
	j, err := cfg.GetString(s)
	if err != nil {
		log.Errorln(err)
	}
	return j
}

//GetBool 返回bool
func GetBool(s string) bool {
	j, err := cfg.GetBool(s)
	if err != nil {
		log.Errorln(err)
	}
	return j
}

//GetInt 返回int
func GetInt(s string) int {
	j, err := cfg.GetInt(s)
	if err != nil {
		log.Errorln(err)
	}
	return j
}

//Set 写入值
func Set(key string, value interface{}) {
	cfg.Set(key, value)
}
