package apis

import (
	"goBlog/middleware"
	"goBlog/models"
	"goBlog/src/common"

	//jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

//JwtToUserAPI 解析jwt里的数据
func JwtToUserAPI(c *gin.Context) {
	u, ok := c.Get(middleware.GetIdentityKey())
	if ok {
		msg := "success!"
		common.Rmsg(c, true, msg, u)
		return
	}
	msg := "User does not exist"
	common.Rmsg(c, false, msg)
}

//JwtOkAPI 测试jwt功能
func JwtOkAPI(c *gin.Context) {
	u := new(models.UserApi)
	u.Email = "leiju@outlook.com"
	//u.GetUser()
	msg := "JwtOK"
	common.Rmsg(c, true, msg, u)
}
