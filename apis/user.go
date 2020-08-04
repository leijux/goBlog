package apis

import (
	"fmt"

	"goBlog/log"
	"goBlog/models/user"
	"goBlog/src/common"

	"github.com/gin-gonic/gin"
)

//AddUserAPI 添加用户
func AddUserAPI(c *gin.Context) {
	var u user.UserApi
	err := c.Bind(&u)
	if err != nil {
		msg := fmt.Sprintln("shoul bind err")
		log.Logger.Errorln(err)
		common.Rmsg(c, false, msg)
		return
	}

	b, err := u.AddUser()
	if !b {
		msg := fmt.Sprintln("add user err")
		log.Logger.Errorln(err)
		common.Rmsg(c, false, msg)
		return
	}

	msg := fmt.Sprintf("add user success")
	common.Rmsg(c, true, msg)
}
