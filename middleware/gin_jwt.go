package middleware

import (
	"time"

	"goBlog/log"
	"goBlog/models/logn"
	"goBlog/models/user"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var identityKey = "Emeil"

var AuthMiddleware *jwt.GinJWTMiddleware

func init() {
	var err error
	AuthMiddleware, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     24 * time.Hour,
		MaxRefresh:  24 * time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims { //负载
			if v, ok := data.(*user.User); ok {
				return jwt.MapClaims{
					identityKey: v.Emeil,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} { //解析负载
			claims := jwt.ExtractClaims(c)
			return &user.User{
				Emeil: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) { //登入
			var loginVals logn.Logn
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			if b, User, _ := loginVals.PwdCheck(); b {
				return &User, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool { //验证
			if _, ok := data.(*user.User); ok {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Logger.Fatalln("JWT Error:" + err.Error())
	}
}

// TokenLookup：token检索模式，用于提取token，默认值为header:Authorization。
// SigningAlgorithm：签名算法，默认值为HS256
// Timeout：token过期时间，默认值为time.Hour
// TimeFunc：测试或服务器在其他时区可设置该属性，默认值为time.Now
// TokenHeadName：token在请求头时的名称，默认值为Bearer
// IdentityKey：身份验证的key值，默认值为identity
// Realm：可以理解成该中间件的名称，用于展示，默认值为gin jwt
// CookieName：Cookie名称，默认值为jwt
// privKey：私钥
// pubKey：公钥
// Authenticator函数：根据登录信息对用户进行身份验证的回调函数
// PayloadFunc函数：登录期间的回调的函数
// IdentityHandler函数：解析并设置用户身份信息
// Authorizator函数：接收用户信息并编写授权规则，本项目的API权限控制就是通过该函数编写授权规则的
// Unauthorized函数：处理不进行授权的逻辑
// LoginResponse函数：完成登录后返回的信息，用户可自定义返回数据，默认返回
