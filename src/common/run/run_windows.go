package run

import (
	"net/http"
	"os"
	"os/signal"
	"time"

	"goBlog/log"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func run(prot string, handler http.Handler) {
	srv := &fasthttp.Server{
		Handler:     fasthttpadaptor.NewFastHTTPHandler(handler),
		ReadTimeout: 15 * time.Second,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(prot); err != nil && err != fasthttp.ErrConnectionClosed {
			log.Logger.Fatalf("listen: %s\n", err)
		}

		// if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		// 	log.Logger.Fatalf("listen: %s\n", err)
		// }
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Logger.Infoln("Shutdown Server ...")

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	if err := srv.Shutdown(); err != nil {
		log.Logger.Fatal("Server Shutdown:", err)
	}
	log.Logger.Infoln("Server exiting")
}
