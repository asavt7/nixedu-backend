package configs

import "github.com/go-redis/redis/v8"

func InitRedisConf() *redis.Options {
	return &redis.Options{
		Addr: "localhost:6379",
	}
}
