package router

import (
	"goBlog/apis"
	"goBlog/middleware"
	v1 "goBlog/router/v1"
	"goBlog/src/common"
	"goBlog/validator/vuser"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitRouter(router *gin.Engine) {
	router.Use(gin.Recovery(), middleware.Log(), middleware.Cors(), middleware.Authorize())
	router.Static("/files", "./web")
	router.GET("/", common.Handler()(apis.Index))

	//添加自定义 tag
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("vuserName", vuser.VuserName)
	}

	v1.V1(router)
}
