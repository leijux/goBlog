package apis

import (
	"fmt"
	"net/http"

	"task-system/middleware"
	"task-system/models/user"
	"task-system/src/common"

	//jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

//JwtToUserAPI 解析jwt里的数据
func JwtToUserAPI(c *gin.Context) {
	u, ok := c.Get(middleware.AuthMiddleware.IdentityKey)
	if ok {
		msg := fmt.Sprintln("User does not exist")
		common.Rmsg(c, http.StatusOK, msg, false)
		return
	}
	common.Rmsg(c, http.StatusOK, "success!", u)
}

//JwtOkAPI 测试jwt功能
func JwtOkAPI(c *gin.Context) {
	u:=new(user.User)
	u.Emeil = "leiju@outlook.com"
	u.GetUser()
	msg := "JwtOK"
	common.Rmsg(c, http.StatusOK, msg, u)
}
