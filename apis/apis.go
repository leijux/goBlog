package apis

import (
	"net/http"

	"goBlog/models/blog"

	"github.com/gin-gonic/gin"
)

//Index 默认API
func Index(c *gin.Context) {
	//c.String(http.StatusOK, "this test!")
	var b blog.Blog
	i, _ := b.Count()
	c.JSON(http.StatusOK, gin.H{
		"msg":  "test success！",
		"data": i,
	})
}
