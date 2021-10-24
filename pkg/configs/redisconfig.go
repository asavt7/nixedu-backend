package configs

import (
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// InitRedisConf read configs from envs\config files and return *redis.Options
func InitRedisConf() *redis.Options {
	op := &redis.Options{
		Addr: viper.GetString("redis.host") + ":" + viper.GetString("redis.port"),
	}
	log.Infof("Redis configs %v", op)
	return op
}
