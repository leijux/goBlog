package middleware

import (
	"fmt"
	"goBlog/log"

	"github.com/casbin/casbin"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Authorize() gin.HandlerFunc {
	a, err := xormadapter.NewAdapter("mysql", "root:123456@tcp(47.101.147.15:3306)/leiju?parseTime=true", true)
	if err != nil {
		log.Fatalln(err)
	}

	e := casbin.NewEnforcer("./middleware/rbac_models.conf", a)
	_ = e.LoadPolicy()
	return func(c *gin.Context) {

		//获取请求的URI
		obj := c.Request.URL.RequestURI()
		//获取请求方法
		act := c.Request.Method
		//获取用户的角色
		sub := "admin"

		//判断策略中是否存在
		if ok := e.Enforce(sub, obj, act); ok {
			fmt.Println("恭喜您,权限验证通过")
			c.Next()
		} else {
			fmt.Println("很遗憾,权限验证没有通过")
			c.Abort()
		}
	}
}
