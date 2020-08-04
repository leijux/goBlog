package apis

import (
	"fmt"

	"goBlog/middleware"
	"goBlog/models/user"
	"goBlog/src/common"

	//jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

//JwtToUserAPI 解析jwt里的数据
func JwtToUserAPI(c *gin.Context) {
	u, ok := c.Get(middleware.GetIdentityKey())
	if ok {
		common.Rmsg(c, true, "success!", u)
		return
	}
	msg := fmt.Sprintln("User does not exist")
	common.Rmsg(c, false, msg)
}

//JwtOkAPI 测试jwt功能
func JwtOkAPI(c *gin.Context) {
	u := new(user.UserApi)
	u.Email = "leiju@outlook.com"
	// u.GetUser()
	msg := "JwtOK"
	common.Rmsg(c, true, msg, u)
}
