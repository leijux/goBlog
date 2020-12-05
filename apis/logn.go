package apis

import (
	"goBlog/log"
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

	var login login.LoginApi
	login.Email = email
	b, err := login.EmailCheck()
	if !b {
		log.Errorf("%+v", err)
		const msg = "email check err"
		common.Rmsg(c, false, msg)
		return
	}
	const msg = "email check ok"
	common.Rmsg(c, true, msg)
}
