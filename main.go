package main

import (
	"goBlog/config"
	"goBlog/database/cache"
	"goBlog/database/orm"
	_ "goBlog/docs"
	"goBlog/log"
	"goBlog/models/blog"
	"goBlog/models/comment"
	"goBlog/models/likes"
	"goBlog/models/user"
	"goBlog/router"
	"goBlog/src/common"
	"goBlog/src/common/run"
	"net/http"

	"github.com/DeanThompson/ginpprof"
	"github.com/arl/statsviz"
	"github.com/gin-gonic/gin"
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
	err := orm.Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		new(user.User),
		new(blog.Blog),
		new(likes.Likes),
		new(comment.Comment),
	)

	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	defer orm.Close()   //关闭gorm
	defer cache.Close() //关闭缓存


	isDebugMode = config.GetBool("gin.isDebugMode")

	if !isDebugMode { //判断模式，如果是debug模式则开启pprof
		gin.SetMode(gin.ReleaseMode) //发布模式
		//log.Logger.SetLevel(logrus.ErrorLevel) // 设置日志级别 在什么级别之上
	}

	//gin.DisableConsoleColor() //静止控制台颜色，防止有空格

	r := setupRouter()
	run.Run(r)
}

func setupRouter() (r *gin.Engine) {
	r = gin.New()
	router.InitRouter(r) //设置路由
	webURL := config.GetString("gin.open")
	if isDebugMode { //判断模式，如果是debug模式则开启pprof
		ginpprof.Wrap(r)       //go tool pprof -http=:8080 cpu.prof
		go common.Open(webURL) // http://localhost:8000/
		statsviz.RegisterDefault()
		go func() {
			// http://localhost:6060/debug/statsviz/
			log.Debug(http.ListenAndServe("localhost:6060", nil).Error())
		}()
	}
	return
}
