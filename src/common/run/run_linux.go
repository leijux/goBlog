package run

import (
	"goBlog/log"
	"net/http"
	"time"

	"github.com/fevin/gracehttp"
)

func run(prot string, handler http.Handler) {
	server := gracehttp.NewGraceHTTP()
	if _, err := server.AddServer(&gracehttp.ServerOption{
		HTTPServer: &http.Server{
			Addr:           prot,
			Handler:        handler,
			ReadTimeout:    1 * time.Second,
			WriteTimeout:   1 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}); err != nil {
		log.Fatalln(err)
	}
	if err := server.Run(); err != nil && err != http.ErrServerClosed {
		log.Fatalln(err)
	}
	// if err := endless.ListenAndServe(prot, handler); err != nil {
	// 	log.Logger.Fatalf("listen: %s\n", err)
	// }
}
