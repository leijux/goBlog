package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowHeaders = []string{"access-control-allow-origin, access-control-allow-headers,Authorization"}
	return cors.New(config)
}
