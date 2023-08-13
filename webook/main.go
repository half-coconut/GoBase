package main

import (
	"GoBase/webook/internal/repository"
	"GoBase/webook/internal/repository/dao"
	"GoBase/webook/internal/service"
	"GoBase/webook/internal/web"
	"GoBase/webook/internal/web/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

/**
得到一个指针，用 & 取地址
申明一个指针，用 * 指针
*/

func main() {
	db := initDB()
	u := initUser(db)
	server := initWebServer()
	//u.RegisterRoutesV1(server.Group("/users")) // v1 方法二
	u.RegisterRoutes(server)
	server.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}

func initWebServer() *gin.Engine {
	server := gin.Default()
	server.Use(func(c *gin.Context) {
		println("这是第一个 middleware")
	})
	server.Use(func(c *gin.Context) {
		println("这是第二个 middleware")
	})

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
	//store := cookie.NewStore([]byte("secret"))
	// 基于 redis 实现的 store
	store, err := redis.NewStore(16, "tcp", "localhost:6379", "",
		// authentication key 身份认证, encryption key 数据加密
		[]byte("iyI1vQON0NmwDnaOMZAgdcJQZ7N6TYbD"),
		[]byte("OOPmqabOfgrBdeXk1545Dc1pS6JbCkUg"))
	if err != nil {
		panic(err)
	}
	//myStore := &sqlx_store.Store{}
	server.Use(sessions.Sessions("mysession", store))

	server.Use(middleware.NewLoginMiddlewareBuilder().
		IgnorePaths("/users/login").
		IgnorePaths("/users/signup").Builder())
	return server
}

func initUser(db *gorm.DB) *web.UserHandler {
	ud := dao.NewUserDAO(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := web.NewUserHandler(svc)
	return u
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"), &gorm.Config{})
	if err != nil {
		// 在初始化过程中，panic
		// panic 使得 goroutine 直接结束
		// 一旦初始化过程出错，应用就不要启动了
		panic(err)
	}
	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return db
}
