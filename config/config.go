package config

import (
	"log"
	"sync"

	"github.com/golobby/config"
	"github.com/golobby/config/feeder"
	"github.com/pkg/errors"
)

const (
	//Mysql 字段
	Mysql = "mysql"
	//Sqlite 字段
	Sqlite = "sqlite3"
)

var (
	ErrOpenFile = errors.New("open err")

	//JSONPath 文件路径
	JSONPath = "./config/config.json"
)

//cfg 配置文件对象
var cfg *config.Config
var once sync.Once

// func Init(path string) {
// 	once.Do(func() {
// 		configInit(path)
// 	})
// }

// func configInit(path string) {
// 	var err error
// 	cfg, err = config.New(config.Options{
// 		Feeder: feeder.Json{Path: path},
// 	})
// 	if err != nil {
// 		log.Fatalf("%+v", errors.Wrap(err, ErrOpenFile.Error()))
// 	}
// }
func init() {
	var err error
	cfg, err = config.New(config.Options{
		Feeder: feeder.Json{Path: JSONPath},
	})
	if err != nil {
		log.Fatalf("%+v", errors.Wrap(err, ErrOpenFile.Error()))
	}
}

//GetString 返回字符串
func GetString(key string) string {
	j, err := cfg.GetString(key)
	if err != nil {
		log.Fatalln(err)
		return ""
	}
	return j
}

//GetBool 返回bool
func GetBool(key string) bool {
	j, err := cfg.GetBool(key)
	if err != nil {
		log.Fatalln(err)
		return false
	}
	return j
}

//GetInt 返回int
func GetInt(key string) int {
	j, err := cfg.GetInt(key)
	if j == 0 {
		return j
	}
	if err != nil {
		log.Fatalln(err)
		return 0
	}
	return j
}

//Get 获取数据
func Get(key string) (value interface{}) {
	value, err := cfg.Get(key)
	if err != nil {
		log.Fatalln(err)
		return
	}
	return
}

//Set 写入值 写入内存不写入文件
func Set(key string, value interface{}) {
	cfg.Set(key, value)
}
