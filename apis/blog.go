package apis

import (
	"fmt"
	"net/http"
	"task-system/log"
	"task-system/models/blog"
	"task-system/src/common"
	"time"

	"github.com/gin-gonic/gin"
)

func AddBlogAPI(c *gin.Context) {
	var b blog.Blog
	err := c.Bind(&b)
	if err != nil {
		msg := fmt.Sprintln("shoul bind err")
		log.Logger.Errorln(err)
		common.Rmsg(c, http.StatusOK, msg, false)
		return
	}
	t := time.Now()
	b.Created = t
	b.Updated = t
	id, err := b.AddBlog()
	if err != nil {
		msg := fmt.Sprintln("add blog err")
		log.Logger.Errorln(err)
		common.Rmsg(c, http.StatusOK, msg, false)
		return
	}

	msg := fmt.Sprintf("add blog id")
	common.Rmsg(c, http.StatusOK, msg, id)
}
