package router

import (
	"task-system/apis"
	"task-system/middleware"
	v1 "task-system/router/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	router.Use(gin.Recovery(), middleware.LoggerToFile())
	router.Static("/files", "./web")
	router.GET("/", apis.Index)
	v1.V1(router)
}
