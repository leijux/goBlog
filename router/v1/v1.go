package v1

import (
	"goBlog/apis"
	"goBlog/middleware"

	"github.com/gin-gonic/gin"
)

func V1(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		{ //logn
			v1.POST("/logn", apis.AddUserAPI) //注册
			v1.GET("/emeilChack", apis.EmeilChackAPI)
			v1.GET("/logn", middleware.AuthMiddleware.LoginHandler) //登入
		}

		{ //suer
			v1.DELETE("/user") //删除

			v1.PUT("/user") //更新
			v1.PATCH("/user")

			v1.GET("/user", middleware.AuthMiddleware.MiddlewareFunc(), apis.JwtToUserAPI)
		}
		{ //验证jwt
			v1.GET("/jwt", middleware.AuthMiddleware.MiddlewareFunc(), apis.JwtOkAPI) //.Use(middleware.AuthMiddleware.MiddlewareFunc())

		}

		{ //blog
			v1.POST("/blog", middleware.AuthMiddleware.MiddlewareFunc(), apis.AddBlogAPI)
			//添加文章
			v1.GET("/blogs", apis.GetBlogsAPI)
			// 获取文章数量
			v1.GET("/blogSize", apis.BlogSizeAPI)
			v1.GET("/getTop", apis.GetTopAPI)
		}
	}
}
