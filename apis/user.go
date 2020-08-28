package apis

import (
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
		const msg = "shoul bind err"
		log.Logger.Errorln(err)
		common.Rmsg(c, false, msg)
		return
	}

	msg, b := addUserAPI(u)
	common.Rmsg(c, b, msg)
}

func addUserAPI(u user.UserApi) (string, bool) {
	b, err := u.CreateUser()
	if b {
		const msg = "add user success"
		return msg, b
	}
	const msg = "add user err"
	log.Logger.Errorln(err)
	return msg, b
}
