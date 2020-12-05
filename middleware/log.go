package middleware

import (
	"time"

	"goBlog/log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//LoggerToFile 日志记录到文件
func Log() gin.HandlerFunc {

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()

		latencyTime := time.Since(startTime) // 执行时间
		reqMethod := c.Request.Method        // 请求方式
		reqURI := c.Request.RequestURI       // 请求路由
		statusCode := c.Writer.Status()      // 状态码
		clientIP := c.ClientIP()             // 请求IP
		log.Info("Router",
			zap.Int("status_code", statusCode),
			zap.Duration("latencyTime", latencyTime),
			zap.String("client_ip", clientIP),
			zap.String("req_method", reqMethod),
			zap.String("req_uri", reqURI),
		)
		// 日志格式
		// log.Logger.WithFields(logrus.Fields{
		// 	"status_code":  statusCode,
		// 	"latency_time": latencyTime,
		// 	"client_ip":    clientIP,
		// 	"req_method":   reqMethod,
		// 	"req_uri":      reqURI,
		// 	// userPhone:c.MustGet("claims").(*myjwt.CustomClaims).Phone
		// })
	}
}

// //LoggerToMongo 日志记录到 MongoDB
// func LoggerToMongo() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 	}
// }

// //LoggerToES 日志记录到 ES
// func LoggerToES() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 	}
// }

// //LoggerToMQ 日志记录到 MQ
// func LoggerToMQ() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 	}
// }
