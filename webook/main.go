package main

import (
	"GoBase/webook/internal/web"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func main() {
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		//AllowOrigins: []string{"http://localhost:3000/"},
		//AllowMethods: []string{"PUT", "PATCH", "POST", "GET"}, // 不写就是 都支持
		AllowHeaders: []string{"Content-Type", "Authorization"},
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
	}))

	u := web.NewUserHandler()
	//u.RegisterRoutesV1(server.Group("/users")) // v1 方法二
	u.RegisterRoutes(server)
	server.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
