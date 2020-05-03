package router

import (
	v1 "task-system/router/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	router.Use( gin.Recovery())
	router.Static("/files", "./web")

	v1.V1(router)

}
