package ioc

import (
	"GoBase/webook/internal/web"
	"GoBase/webook/internal/web/middleware"
	"GoBase/webook/pkg/ginx/middlewares/ratelimit"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"
)

func InitWebServer(mdls []gin.HandlerFunc, userHdl *web.UserHandler) *gin.Engine {
	server := gin.Default()
	server.Use(mdls...)
	userHdl.RegisterRoutes(server)
	return server
}

func InitMiddlewares(redisClient redis.Cmdable) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		corsHdl(),
		middleware.NewLoginJWTMiddlewareBuilder().
			IgnorePaths("/users/login").
			IgnorePaths("/hello").
			IgnorePaths("/users/login_sms/code/send").
			IgnorePaths("/users/login_sms").
			IgnorePaths("/users/signup").Builder(),
		ratelimit.NewBuilder(redisClient, time.Second, 100).Build(),
	}
}
func corsHdl() gin.HandlerFunc {
	return cors.New(cors.Config{
		//AllowOrigins: []string{"http://localhost:3000/"},
		//AllowMethods: []string{"PUT", "PATCH", "POST", "GET"}, // 不写就是 都支持
		AllowHeaders: []string{"Content-Type", "Authorization"},
		// 前端可以拿到这个值
		ExposeHeaders: []string{"x-jwt-token"},
		//ExposeHeaders:    []string{"Content-Length"},
		// 是否允许带 cookie 之类
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				// 开发环境
				return true
			}
			return strings.Contains(origin, "coconut.com")
		},
		MaxAge: 12 * time.Hour,
	})
}
