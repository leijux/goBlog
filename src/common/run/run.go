package run

import (
	"goBlog/config"
	"goBlog/log"

	"github.com/gin-gonic/gin"
)

//Run 运行服务
func Run(router *gin.Engine) {
	prot := config.GetString("gin.prot")
	if prot == "" {
		log.Logger.Fatalln("prot is empty")
	}
	run(prot, router)
}
