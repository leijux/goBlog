package apis

import (
	"goBlog/log"
	"goBlog/models"
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

	var l models.LoginApi
	l.Email = email
	b, err := l.EmailCheck()
	if !b {
		log.Errorln(err)
		const msg = "email check err"
		common.Rmsg(c, false, msg)
		return
	}
	const msg = "email check ok"
	common.Rmsg(c, true, msg)
}
