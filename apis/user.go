package apis

import (
	"fmt"
	"time"

	"goBlog/log"
	"goBlog/models/user"
	"goBlog/src/common"

	"github.com/gin-gonic/gin"
)

//AddUserAPI 添加用户
func AddUserAPI(c *gin.Context) {
	var u user.User
	err := c.Bind(&u)
	if err != nil {
		msg := fmt.Sprintln("shoul bind err")
		log.Logger.Errorln(err)
		common.Rmsg(c, false, msg)
		return
	}
	t := time.Now()
	u.Created = t
	u.Updated = t
	id, err := u.AddUser()
	if err != nil {
		msg := fmt.Sprintln("add user err")
		log.Logger.Errorln(err)
		common.Rmsg(c, false, msg)
		return
	}

	msg := fmt.Sprintf("add user id")
	common.Rmsg(c, true, msg, id)
}
