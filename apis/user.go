package apis

import (
	"goBlog/log"
	"goBlog/models/user"
	"goBlog/src/common"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

//AddUserAPI 添加用户
func AddUserAPI(c *gin.Context) {
	u := user.NewUser()

	err := c.Bind(&u)
	if err != nil {
		const msg = "shoul bind err"
		errors.WithMessage(err, msg)
		log.Errorf("%+v ", err)
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
	if err != nil {
		log.Errorf("%+v", err)
	}
	return msg, b
}
