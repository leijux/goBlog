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
	Log struct {
		LOGFILEPATH string `json:"LOG_FILE_PATH"`
		LOGFILENAME string `json:"LOG_FILE_NAME"`
	} `json:"log"`
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
