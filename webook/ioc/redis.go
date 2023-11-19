package ioc

import (
	"fmt"
	"github.com/coocood/freecache"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func InitRedis() redis.Cmdable {
	addr := viper.GetString("redis.addr")
	//viper.GetDuration()
	//viper.GetFloat64() 注意精度
	fmt.Println(addr)
	redisClient := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return redisClient
}

func InitFreeCache() *freecache.Cache {
	// 缓存大小为1MB
	return freecache.NewCache(1024 * 1024)
}

//func NewRateLimiter()redis.Limiter  {
//
//}
