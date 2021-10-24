package config

import (
	log "github.com/sirupsen/logrus"
	"strings"
)
import "github.com/spf13/viper"

// InitConfigs ReadConfigsFromFileAndEnv
func InitConfigs() {

	viper.AutomaticEnv() // Get the value of the environment variable
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigName("config")     // Define the profile name
	viper.SetConfigType("toml")       // Define format
	viper.AddConfigPath("./configs/") // Define file path
	err := viper.ReadInConfig()       // Read
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}

	// set logger level
	level, err := log.ParseLevel(viper.GetString("loglevel"))
	if err != nil {
		log.Fatalf("parse log level failed: %v", err)
	}
	log.SetLevel(level)

}
