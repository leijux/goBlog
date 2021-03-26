package main

import (
	"net/http"

	"goBlog/config"
	"goBlog/database/orm"
	_ "goBlog/docs"
	"goBlog/log"
	"goBlog/models"
	"goBlog/router"
	"goBlog/src/common"
	"goBlog/src/common/run"

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

func init() {
	// 创建表时添加表后缀
	var err = orm.Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		new(models.User),
		new(models.Blog),
		new(models.Likes),
		new(models.Comment),
	)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	mod := config.GetString("gin.mode")
	gin.SetMode(mod) //发布模式
	//log.Logger.SetLevel(logrus.ErrorLevel) // 设置日志级别 在什么级别之上
	gin.DisableConsoleColor() //静止控制台颜色，防止有空格
	r := setupRouter()
	run.Run(r)
}

func setupRouter() (r *gin.Engine) {
	r = gin.New()
	router.InitRouter(r)             //设置路由
	if gin.Mode() == gin.DebugMode { //判断模式，如果是debug模式则开启pprof
		webURL := config.GetString("gin.open")
		ginpprof.Wrap(r)       //go tool pprof -http=:8080 cpu.prof
		go common.Open(webURL) // http://localhost:8000/
		_ = statsviz.RegisterDefault()
		go func() {
			// http://localhost:6060/debug/statsviz/
			log.Debugln(http.ListenAndServe("localhost:6060", nil))
		}()

	}
	return
}
