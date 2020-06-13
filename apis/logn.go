package apis

import (
	"fmt"
	"goBlog/models/logn"
	"goBlog/src/common"

	"github.com/gin-gonic/gin"
)

func EmeilChackAPI(c *gin.Context) {
	emeil := c.Query("emeil")

	var logn logn.Logn
	logn.Emeil = emeil
	b := logn.EmeilChack()
	if !b {
		msg := fmt.Sprintln("emeil chack err")
		common.Rmsg(c, false, msg)
		return
	}
	msg := fmt.Sprintln("emeil chack ok")
	common.Rmsg(c, true, msg)
}
