package main

import (
	"goBlog/config"
	"goBlog/database"
	"goBlog/database/cache"
	"goBlog/router"
	"goBlog/src/common"

	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
)

func main() {
	defer database.Db.Close() //关闭数据库
	defer cache.Close()       //关闭缓存

	r := setupRouter()
	common.Run(r)
}

func setupRouter() (r *gin.Engine) {
	r = gin.New()
	gin.DisableConsoleColor() //静止控制台颜色，防止有空格
	router.InitRouter(r)
	webURL := config.GetString("gin.open")
	if isDebugMode := config.GetBool("gin.isDebugMode"); isDebugMode { //判断模式，如果是debug模式则开启pprof
		ginpprof.Wrap(r)       //go tool pprof -http=:8080 cpu.prof
		go common.Open(webURL) // http://localhost:8080/
	} else {
		gin.SetMode(gin.ReleaseMode) //发布模式
	}
	return
}
