package apis

import (
	"goBlog/log"
	"goBlog/models/user"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

//AddUserAPI 添加用户
func AddUserAPI(c *gin.Context) (code bool, msg string, data interface{}) {
	u := user.NewUser()

	err := c.Bind(&u)
	if err != nil {
		const m = "shoul bind err"
		err = errors.WithMessage(err, msg)
		msg = err.Error()
		log.Errorf("%+v ", err)
		code = false
		return
	}

	msg, code = addUserAPI(u)
	data = nil
	return

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
