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
func GetString(key string) string {
	j, err := cfg.GetString(key)
	if err != nil {
		log.Errorln(err)
		return ""
	}
	return j
}

//GetBool 返回bool
func GetBool(key string) bool {
	j, err := cfg.GetBool(key)
	if err != nil {
		log.Errorln(err)
		return false
	}
	return j
}

//GetInt 返回int
func GetInt(key string) int {
	j, err := cfg.GetInt(key)
	if err != nil {
		//log.Errorln(err)
		return 0
	}
	return j
}

//Get 获取数据
func Get(key string) (value interface{}) {
	value, err := cfg.Get(key)
	if err != nil {
		log.Errorln(err)
		return
	}
	return
}

//Set 写入值 写入内存不写入文件
func Set(key string, value interface{}) {
	cfg.Set(key, value)
}
