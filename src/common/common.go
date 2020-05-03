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

	"task-system/log"

	"github.com/gin-gonic/gin"
)

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

func Run(router *gin.Engine, prot string) {
	srv := &http.Server{
		Addr:    prot,
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
	log.Logger.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Logger.Fatal("Server Shutdown:", err)
	}
	log.Logger.Println("Server exiting")
}
