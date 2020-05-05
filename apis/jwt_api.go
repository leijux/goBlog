package apis

import (
	
	"net/http"
	
	"task-system/log"
	"task-system/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func JwtToUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		u, _ := c.Get(middleware.AuthMiddleware.IdentityKey)
		log.Logger.Info(u)
		c.JSON(http.StatusOK, gin.H{
			// "emeil": claims[middleware.AuthMiddleware.IdentityKey],
			"name":  claims["Name"],
			"user":u,
		})
	}
}

func JwtOk() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "jwtOK",
		})
	}
}
