package apis

import (
	"net/http"
	"task-system/models/blog"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	//c.String(http.StatusOK, "this test!")
	var b blog.Blog
	a, _ := b.Count()
	c.JSON(http.StatusOK, gin.H{
		"a":a,
	})
}
