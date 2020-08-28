package apis

import (
	"fmt"
	"goBlog/models/login"
	"goBlog/src/common"

	"github.com/gin-gonic/gin"
)

// @Summary EmailCheckApi
// @Description 验证邮箱
// @Tags 测试
// @Accept json
// @Param email query string true "邮箱"
// @Success 200 {string} json "{"msg": "email check ok"}"
// @Failure 400 {string} json "{"msg": "email check err"}"
// @Router /emailCheck [get]
func EmailCheckApi(c *gin.Context) {
	email := c.Query("email")

	var login login.Login
	login.Email = email
	b := login.EmailCheck()
	if !b {
		msg := fmt.Sprintln("email check err")
		common.Rmsg(c, false, msg)
		return
	}
	msg := fmt.Sprintln("email check ok")
	common.Rmsg(c, true, msg)
}
