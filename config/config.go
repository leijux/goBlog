package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const (
	Jsonfile = "./config/config.json"
	Mysql    = "mysql"
	Sqlite   = "sqlite"
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
			Addr     string `json:"Addr"`
			Password string `json:"Password"`
		} `json:"redis"`
	} `json:"database"`
}

var Cfg *config

func init() {
	f, err := ioutil.ReadFile(Jsonfile)
	if err != nil {
		log.Fatalln(err, "ioutil.ReadFile")
	}
	err = json.Unmarshal(f, &Cfg)
	if err != nil {
		log.Fatalln(err)
	}
}
