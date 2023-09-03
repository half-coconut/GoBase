package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
得到一个指针，用 & 取地址
申明一个指针，用 * 指针
*/

func main() {
	server := InitWebServer()
	server.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello,gogogo!")
	})
	server.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
