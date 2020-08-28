package apis

import (
	"net/http"

	"goBlog/models/blog"
	"github.com/gin-gonic/gin"
)

// @Summary 测试index
// @Description testApi
// @Tags 测试
// @Accept json
// @Success 200 {json} string "{"msg": "test success！"}"
// @Router / [get]
func Index(c *gin.Context) {
	//c.String(http.StatusOK, "this test!")
	var b blog.Blog
	i, _ := b.Count()
	c.JSON(http.StatusOK, gin.H{
		"msg":  "test success!",
		"data": i,
	})
}
