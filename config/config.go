package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

const (
	//Mysql 字段
	Mysql = "mysql"
	//Sqlite 字段
	Sqlite = "sqlite3"
)

var (

	//JSONPath 文件路径
	JSONPath = "./config/"
)

func init() {
	if gin.Mode() == gin.DebugMode {
		JSONPath = "E:\\OneDrive\\learning\\program\\Go\\myGoProjects\\goBlog\\config"
	}
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("json")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(JSONPath) // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	_ = viper.WriteConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file:%s Op:%s\n", e.Name, e.Op)
	})
}

//GetString 返回字符串
func GetString(key string) string {
	return viper.GetString(key)
}

//GetBool 返回bool
func GetBool(key string) bool {
	return viper.GetBool(key)
}

//GetInt 返回int
func GetInt(key string) int {
	return viper.GetInt(key)
}

//Set 写入值 写入内存不写入文件
func Set(key string, value interface{}) error {
	viper.Set(key, value)
	return viper.WriteConfig()
}
