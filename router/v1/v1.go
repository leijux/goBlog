package v1

import (
	"goBlog/apis"
	"goBlog/middleware"

	"github.com/gin-gonic/gin"
)

//V1 v1版本的路由
func V1(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		{ //logn
			v1.POST("/logn", apis.AddUserAPI)          //注册
			v1.GET("/emeilChack", apis.EmeilChackAPI)  //邮箱验证api
			v1.GET("/logn", middleware.LoginHandler()) //登入
		}

		{ //user
			v1.DELETE("/user") //删除

			v1.PUT("/user") //更新
			v1.PATCH("/user")

			v1.GET("/user", middleware.JwtMiddlewareFunc(), apis.JwtToUserAPI) //根据jwt里的信息解析用户信息
		}

		{ //验证jwt
			v1.GET("/jwt", middleware.JwtMiddlewareFunc(), apis.JwtOkAPI) //
		}

		{ //blog
			//添加文章
			v1.POST("/blog", middleware.JwtMiddlewareFunc(), apis.AddBlogAPI)
			//得到文章
			v1.GET("/blogs", apis.GetBlogsAPI)
			// 获取文章数量
			v1.GET("/blogSize", apis.BlogSizeAPI)
			//得到文章排名
			v1.GET("/getTop", apis.GetTopAPI)
		}
	}
}
