package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const (
	Jsonfile = "./config/config.json"
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
	Sql struct {
		Mysql struct {
			DriverName     string `json:"driverName"`
			DataSourceName string `json:"dataSourceName"`
		} `json:"mysql"`
		Sqlist struct {
			DriverName     string `json:"driverName"`
			DataSourceName string `json:"dataSourceName"`
		} `json:"sqlist"`
	} `json:"sql"`
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
