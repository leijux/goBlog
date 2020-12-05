package common

import (
	"net/http"
	"os/exec"
	"runtime"

	"goBlog/log"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
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
		err = errors.New("unsupported platform")
	}
	return
}

type MyHandler func(c *gin.Context) (bool, string, interface{})

func Handler() func(h MyHandler) gin.HandlerFunc {
	return func(h MyHandler) gin.HandlerFunc {
		return func(c *gin.Context) {
			code, msg, data := h(c)
			Rmsg(c, code, msg, data)
			return
		}
	}
}

//Rmsg 返回请求
func Rmsg(c *gin.Context, code bool, msg string, data ...interface{}) {
	if data == nil {
		rmsg(c, code, msg, nil)
	}
	rmsg(c, code, msg, data)
}

func rmsg(c *gin.Context, code bool, msg string, data interface{}) {
	json := gin.H{
		"code": code,
	}
	if msg != "" {
		json["msg"] = msg
	}
	if data != nil {
		json["data"] = data
	}

	c.JSON(http.StatusOK, json)
}

func Scrypt(paw string) (dk []byte, err error) {
	const salt = "leiju"
	dk, err = scrypt.Key([]byte(paw), []byte(salt), 16384, 8, 1, 32)

	if err != nil {
		log.Error("Scrypt err",
			zap.Error(err),
		)
		return
	}
	return
}
