package apis

import (
	"fmt"
	"net/http"

	"task-system/log"
	"task-system/middleware"

	//jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

//JwtToUserAPI 解析jwt里的数据
func JwtToUserAPI(c *gin.Context) {
	u, ok := c.Get(middleware.AuthMiddleware.IdentityKey)
	if ok {
		msg := fmt.Sprintln("User does not exist")
		log.Logger.Info(msg)
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  msg,
			"data": false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  nil,
		"data": u,
	})
}

//JwtOkAPI 测试jwt功能
func JwtOkAPI(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "jwtOK",
		"data": true,
	})
}
