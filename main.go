package main

import (
	"goBlog/config"
	"goBlog/database"
	"goBlog/database/cache"
	"goBlog/database/orm"
	_ "goBlog/docs"
	"goBlog/log"
	"goBlog/models/blog"
	"goBlog/router"
	"goBlog/src/common"
	"os/user"

	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celled server.
// @termsOfService https://www.topgoer.com

// @contact.name goBlog
// @contact.url
// @contact.email leijuxx@outlook.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /v1

var isDebugMode bool

func init() {

	// 创建表时添加表后缀
	orm.Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(new(user.User), new(blog.Blog))
}

func main() {

	defer database.Db.Close() //关闭数据库
	defer cache.Close()       //关闭缓存

	isDebugMode = config.GetBool("gin.isDebugMode")

	if !isDebugMode { //判断模式，如果是debug模式则开启pprof
		gin.SetMode(gin.ReleaseMode)           //发布模式
		log.Logger.SetLevel(logrus.ErrorLevel) // 设置日志级别 在什么级别之上
	}

	//gin.DisableConsoleColor() //静止控制台颜色，防止有空格

	r := setupRouter()
	common.Run(r)
}

func setupRouter() (r *gin.Engine) {
	r = gin.New()
	router.InitRouter(r) //设置路由
	webURL := config.GetString("gin.open")
	if isDebugMode { //判断模式，如果是debug模式则开启pprof
		ginpprof.Wrap(r)       //go tool pprof -http=:8080 cpu.prof
		go common.Open(webURL) // http://localhost:8000/
	}
	return
}
