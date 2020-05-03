package v1

import "github.com/gin-gonic/gin"

func V1(router *gin.Engine){
	v1:=router.Group("v1")
	{
		v1.GET("/x",)
	}
}