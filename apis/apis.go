package apis

import (
	"goBlog/models/blog"

	"github.com/gin-gonic/gin"
)

// @Summary 测试index
// @Description testApi
// @Tags 测试
// @Accept json
// @Success 200 {json} string "{"msg": "test success！"}"
// @Router / [get]
func Index(c *gin.Context) (bool, string, interface{}) {
	return index(c)
}
func index(c *gin.Context) (bool, string, interface{}) {
	i, _ := blog.Count()
	return true, "test success!", i
}
