package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

/**
得到一个指针，用 & 取地址
申明一个指针，用 * 指针
*/

func main() {
	initViperV1()
	server := InitWebServer()
	server.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello, go go go!")
	})
	server.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}

func initViperV1() {
	viper.SetConfigFile("config/dev.yaml")
	//viper.KeyDelimiter("-")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initViper() {
	// 配置文件的名字，但是不包含文件扩展名
	// 不包含 .go  .yaml 之类的后缀
	viper.SetConfigName("dev")
	// 告诉 viper 我的配置用的是 yaml 格式
	// 有很多种格式：json、xml、yaml、toml、ini
	viper.SetConfigType("yaml")
	// 当前工作目录下的 config 子目录
	viper.AddConfigPath("./config")
	//viper.AddConfigPath("/tmp/config")
	//viper.AddConfigPath("/etc/webook")
	// 读取配置到 viper 里面，可以理解为加载到内存
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// 可以有多个 viper 的实例
	//otherViper := viper.New()
	//otherViper.SetConfigName("myJson")
	//otherViper.AddConfigPath("./config")
	//otherViper.SetConfigType("json")
}
