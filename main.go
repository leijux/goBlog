package main

import (
	"goBlog/log"
	"goBlog/config"
	"goBlog/database"
	"goBlog/database/cache"
	"goBlog/router"
	"goBlog/src/common"

	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var isDebugMode bool

func main() {
	defer database.Db.Close() //关闭数据库
	defer cache.Close()       //关闭缓存

	isDebugMode = config.GetBool("gin.isDebugMode")

	if !isDebugMode { //判断模式，如果是debug模式则开启pprof
		gin.SetMode(gin.ReleaseMode)       //发布模式
		log.Logger.SetLevel(logrus.ErrorLevel) // 设置日志级别 在什么级别之上
	}

	gin.DisableConsoleColor() //静止控制台颜色，防止有空格

	r := setupRouter()
	common.Run(r)
}

func setupRouter() (r *gin.Engine) {
	r = gin.New()
	router.InitRouter(r) //设置路由
	webURL := config.GetString("gin.open")
	if isDebugMode { //判断模式，如果是debug模式则开启pprof
		ginpprof.Wrap(r)       //go tool pprof -http=:8080 cpu.prof
		go common.Open(webURL) // http://localhost:8080/
	}
	return
}
