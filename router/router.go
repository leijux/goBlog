package router

import (
	"goBlog/apis"
	"goBlog/log"
	"goBlog/middleware"
	r "goBlog/router/v1"
	"goBlog/src/common"
	validator2 "goBlog/validator"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitRouter(router *gin.Engine) {
	router.Use(gin.Recovery(), middleware.Log(), middleware.Cors())
	router.Static("/files", "./web")
	router.GET("/", common.Handler()(apis.Index))

	//添加自定义 tag
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("vuserName", validator2.VUserName)
		if err != nil {
			log.Fatalln(err)
		}
	}

	r.V1(router)
}
