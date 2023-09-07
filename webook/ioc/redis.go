package ioc

import (
	"GoBase/webook/config"
	"github.com/redis/go-redis/v9"
)

func InitRedis() redis.Cmdable {
	redisClient := redis.NewClient(&redis.Options{
		Addr: config.Config.Redis.Addr,
	})
	return redisClient
}

func InitFreeCache() *freecache.Cache {
	// 缓存大小为1MB
	return freecache.NewCache(1024 * 1024)
}
