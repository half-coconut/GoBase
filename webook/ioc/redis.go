package ioc

import "github.com/redis/go-redis/v9"

func InitRedis() redis.Cmdable {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "webook-live-redis:6380",
	})
	return redisClient
}
