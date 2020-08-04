package common

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"

	"goBlog/log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/scrypt"
)

//Open 用默认程序打开文件或者网站
func Open(file string) (err error) {
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", file).Start()
	case "windows":
		err = exec.Command("cmd", "/C", "start", file).Start()
	case "darwin":
		err = exec.Command("open", file).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	return
}

//Rmsg 返回请求
func Rmsg(c *gin.Context, code bool, msg string, data ...interface{}) {
	var json gin.H = gin.H{
		"code": code,
		"msg":  msg,
		"data": "",
	}
	if data != nil {
		json["data"] = data[0]
	}
	log.Logger.WithFields(logrus.Fields(json)).Infoln()
	c.JSON(http.StatusOK, json)
}

func Scrypt(paw string) (dk []byte, err error) {
	salt := "leiju"
	dk, err = scrypt.Key([]byte(paw), []byte(salt), 16384, 8, 1, 32)
	if err != nil {
		log.Logger.Errorln(err)
		return
	}
	return
}
