package main

import (
	"task-system/config"
	"task-system/router"
	"task-system/src/common"

	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
)

func main() {
	//defer db.SqlDB.Close()
	r := gin.New()
	gin.DisableConsoleColor() //静止控制台颜色，防止有空格
	router.InitRouter(r)
	if config.Cfg.Gin.IsDebugMode { //判断模式，如果是debug模式则开启pprof
		ginpprof.Wrap(r)                              //go tool pprof -http=:8080 cpu.prof
		go common.Open(config.Cfg.Gin.Open) // http://localhost:8080/
	}else{
		gin.SetMode(gin.ReleaseMode) //发布模式
	}
	common.Run(r, config.Cfg.Gin.Prot)
}
