package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	myerr "task-system/err"
)

const (
	//Jsonfile 文件路径
	Jsonfile = "./config/config.json"
	//Mysql 字段
	Mysql = "mysql"
	//Sqlite 字段
	Sqlite = "sqlite"
)

type config struct {
	Gin struct {
		IsDebugMode bool   `json:"isDebugMode"`
		Open        string `json:"open"`
		Prot        string `json:"prot"`
	} `json:"gin"`
	Log struct {
		LogFilePath string `json:"logFilePath"`
		LogFileName string `json:"logFileName"`
	} `json:"log"`
	Database struct {
		Enable string `json:"enable"`
		Mysql  struct {
			DriverName     string `json:"driverName"`
			DataSourceName string `json:"dataSourceName"`
		} `json:"mysql"`
		Sqlite struct {
			DriverName     string `json:"driverName"`
			DataSourceName string `json:"dataSourceName"`
		} `json:"sqlite"`
		Redis struct {
			IsOpen   bool   `json:"isOpen"`
			Addr     string `json:"addr"`
			Password string `json:"password"`
			Db       int    `json:"db"`
		} `json:"redis"`
	} `json:"database"`
}

//Cfg 配置文件对象
var Cfg *config

func init() {
	f, err := ioutil.ReadFile(Jsonfile)
	if err != nil {
		log.Fatalln(myerr.ErrOpenFile, err)
	}
	err = json.Unmarshal(f, &Cfg)

	if err != nil {
		log.Fatalln(err)
	}
}

// func Save(){

// }
