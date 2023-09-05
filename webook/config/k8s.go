//go:build k8s

// go build -tags=k8s -o webook .
package config

var Config = config{
	DB: DBConfig{
		DNS: "root:root@tcp(webook-live-mysql:3308)/webook",
	},
	Redis: RedisConfig{
		Addr: "webook-live-redis:6380",
	},
}
