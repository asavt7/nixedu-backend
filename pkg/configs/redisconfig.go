package configs

import "github.com/go-redis/redis/v8"

// InitRedisConf read configs from envs\config files and return *redis.Options
func InitRedisConf() *redis.Options {
	return &redis.Options{
		Addr: "localhost:6379",
	}
}
