package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"go.uber.org/zap"
	"net/http"
)

/**
得到一个指针，用 & 取地址
申明一个指针，用 * 指针
*/

func main() {
	initViperV1()
	initLogger()
	keys := viper.AllKeys()
	println(keys)
	setting := viper.AllSettings()
	println(setting)
	server := InitWebServer()
	server.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello, go go go!")
	})
	server.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}

func initLogger() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	zap.L().Error("这是之前")
	//
	zap.ReplaceGlobals(logger)
	zap.L().Info("这是之后")
	type Demo struct {
		Name string `json:"name"`
	}
	zap.L().Info("这是实验参数",
		zap.Error(errors.New("这是一个 error")),
		zap.Int64("id", 123),
		zap.Any("一个结构体", Demo{Name: "hello go"}))
}

func initViperReader() {
	viper.SetConfigType("yaml")
	cfg := `
db.mysql:
  dsn: "root:root@tcp(localhost:13316)/webook"

redis:
  addr: "localhost:6379"
`
	err := viper.ReadConfig(bytes.NewReader([]byte(cfg)))
	if err != nil {
		panic(err)
	}
}

func initViperV1() {
	cfile := pflag.String("config", "config/config.yaml", "指定配置文件路径")
	pflag.Parse()
	viper.SetConfigFile(*cfile)
	// 实时监听配置变更
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println(in.Name, in.Op)
		fmt.Println(viper.GetString("db.dsn"))
	})
	//viper.SetDefault("db.mysql.dsn", "root:root@tcp(localhost:3306)/webook")
	//viper.SetConfigFile("config/dev.yaml")
	//viper.KeyDelimiter("-")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initViperRemote() {
	viper.SetConfigType("yaml")
	// 通过 webook 和其他使用 etcd 的区别出来
	err := viper.AddRemoteProvider("etcd3", "http://127.0.0.1:12379", "/webook")
	if err != nil {
		panic(err)
	}
	err = viper.ReadRemoteConfig()
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
