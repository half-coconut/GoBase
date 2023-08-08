package main

import (
	"GoBase/webook/internal/repository"
	"GoBase/webook/internal/repository/dao"
	"GoBase/webook/internal/service"
	"GoBase/webook/internal/web"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"), &gorm.Config{})
	if err != nil {
		// 在初始化过程中，panic
		// panic 使得 goroutine 直接结束
		// 一旦初始化过程出错，应用就不要启动了
		panic(err)
	}
	ud := dao.NewUserDAO(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := web.NewUserHandler(svc)

	server := gin.Default()
	// 处理跨域问题
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
	//u.RegisterRoutesV1(server.Group("/users")) // v1 方法二
	u.RegisterRoutes(server)

	server.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
