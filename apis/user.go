package apis

import (
	"fmt"

	"goBlog/log"
	"goBlog/models/user"
	"goBlog/src/common"

	"github.com/gin-gonic/gin"
)

// @Summary AddUserAPI
// @Description 添加用户数据
// @Tags 测试
// @Accept json
// @Param email query string true "邮箱"
// @Success 200 {string} json "{"msg": "email check ok"}"
// @Failure 400 {string} json "{"msg": "email check err"}"
// @Router /emailCheck [get]
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
