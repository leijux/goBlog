package apis

import (
	"fmt"
	"net/http"

	"task-system/log"
	"task-system/middleware"

	//jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func JwtToUser() gin.HandlerFunc {
	return func(c *gin.Context) {
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
}

func JwtOk() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":http.StatusOK,
			"msg": "jwtOK",
			"data":true,
		})
	}
}
