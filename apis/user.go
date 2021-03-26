package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"goBlog/log"
	"goBlog/models"
)

//AddUserAPI 添加用户
func AddUserAPI(c *gin.Context) (code bool, msg string, data interface{}) {
	u := new(models.UserApi)
	err := c.Bind(u)
	if err != nil {
		err = errors.WithMessage(err, "shoul bind err")
		msg = err.Error()
		log.Errorf("%+v ", err)
		code = false
		return
	}

	msg, code = addUserAPI(u)
	data = nil
	return

}

func addUserAPI(u *models.UserApi) (string, bool) {
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
