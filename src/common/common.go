package common

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"time"

	"task-system/config"
	"task-system/log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

//Run 运行服务
func Run(router *gin.Engine) {
	srv := &http.Server{
		Addr:    config.Cfg.Gin.Prot,
		Handler: router,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Logger.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Logger.Infoln("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Logger.Fatal("Server Shutdown:", err)
	}
	log.Logger.Infoln("Server exiting")
}

//Rmsg 返回请求
func Rmsg(c *gin.Context, code int, msg string, data interface{}) {
	json := gin.H{
		"code": http.StatusOK,
		"msg":  msg,
		"data": data,
	}
	log.Logger.WithFields(logrus.Fields(json)).Infoln()
	c.JSON(code, json)
}
