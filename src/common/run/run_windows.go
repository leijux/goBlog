package run

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"goBlog/log"

	"go.uber.org/zap"
)

func run(prot string, handler http.Handler) {
	// srv := &fasthttp.Server{
	// 	Handler:      fasthttpadaptor.NewFastHTTPHandler(handler),
	// 	ReadTimeout:  15 * time.Second,
	// 	WriteTimeout: 5 * time.Second,
	// }

	srv := &http.Server{
		Handler:        handler,
		Addr:           prot,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   1 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		// 服务连接
		// if err := srv.ListenAndServe(prot); err != nil && err != fasthttp.ErrConnectionClosed {
		// 	log.Logger.Fatalf("listen: %s\n", err)
		// }

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen err", zap.Error(err))
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Infoln("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:",
			zap.Error(err),
		)
	}
	log.Infoln("Server exiting")
}
