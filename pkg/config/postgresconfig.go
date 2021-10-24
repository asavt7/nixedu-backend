package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// PostgresConfig - config for connection to postgres db
type PostgresConfig struct {
	Host, Port, Username, Password, DBName string
	SSLMode                                string
}

func (p PostgresConfig) String() string {
	return fmt.Sprintf("postgres host=%s port=%s db=%s ssl=%s ", p.Host, p.Port, p.DBName, p.SSLMode)
}

// InitPostgresConfig read config from envs\config files and returns PostgresConfig
func InitPostgresConfig() PostgresConfig {
	op := PostgresConfig{
		Host:     viper.GetString("pg.host"),
		Port:     viper.GetString("pg.port"),
		Username: viper.GetString("pg.username"),
		Password: viper.GetString("pg.password"),
		DBName:   viper.GetString("pg.dbname"),
		SSLMode:  viper.GetString("pg.sslmode"),
	}
	log.Infof("Postgres config %v", op)
	return op
}
