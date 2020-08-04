package router

import (
	"goBlog/apis"
	"goBlog/middleware"
	v1 "goBlog/router/v1"
	"goBlog/validator/vuser"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	
)

func InitRouter(router *gin.Engine) {
	router.Use(gin.Recovery(), middleware.LoggerToFile(), middleware.Cors())
	router.Static("/files", "./web")
	router.GET("/", apis.Index)
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("vuserName", vuser.VuserName)
	}

	v1.V1(router)
}
