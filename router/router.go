package router

import (
	"goBlog/apis"
	"goBlog/middleware"
	v1 "goBlog/router/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	router.Use(gin.Recovery(), middleware.LoggerToFile(),middleware.Cors())
	router.Static("/files", "./web")
	router.GET("/", apis.Index)
	v1.V1(router)
}
