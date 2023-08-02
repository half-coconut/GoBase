package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
https://gin-gonic.com/zh-cn/docs/quickstart/
go get -u github.com/gin-gonic/gin
go env -w GOPROXY=https://goproxy.cn
*/
func main() {
	r := gin.Default()
	// 静态路由
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello,go!")
	})
	r.POST("/post", func(c *gin.Context) {
		c.String(http.StatusOK, "hello,post method")
	})

	// 参数路由
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "hello，这是参数路由，name为 %s", name)
	})
	// 通配符路由
	r.GET("/view/*.html", func(c *gin.Context) {
		path := c.Param(".html")
		c.String(http.StatusOK, "匹配的值是 %s", path)
	})
	r.GET("/order", func(c *gin.Context) {
		id := c.Query("id")
		c.String(http.StatusOK, "order id 为 "+id)
	})
	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
