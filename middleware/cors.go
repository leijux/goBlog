package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//Cors 跨域中间件
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"access-control-allow-origin", "access-control-allow-headers", "Authorization"}
	return cors.New(config)
}
