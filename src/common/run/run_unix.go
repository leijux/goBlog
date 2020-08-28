package run

import (
	"net/http"

	"goBlog/log"

	"github.com/fvbock/endless"
)

func run(prot string, handler http.Handler) {
	if err := endless.ListenAndServe(prot, handler); err != nil {
		log.Logger.Fatalf("listen: %s\n", err)
	}
}
