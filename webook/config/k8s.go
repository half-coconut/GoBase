//go:build k8s

package config

var Config = config{
	DB: DBConfig{
		DNS: "root:root@tcp(webook-live-mysql:3308)/webook",
	},
	Redis: RedisConfig{
		Addr: "webook-live-redis:6380",
	},
}
