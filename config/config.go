package config

import (
	"fmt"

	"github.com/cyrus28214/comments-area/database"
	"github.com/cyrus28214/comments-area/logging"
	"github.com/cyrus28214/comments-area/server"
	"github.com/spf13/viper"
)

type Config struct {
	Server   server.ServerConfig     `mapstructure:"server"`
	Logging  logging.LoggingConfig   `mapstructure:"logging"`
	Database database.DatabaseConfig `mapstructure:"database"`
}

func GetConfig(configFile string) Config {
	v := viper.New()
	v.SetConfigFile(configFile)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config file: %s", err))
	}
	var c Config
	if err := v.Unmarshal(&c); err != nil {
		panic(fmt.Errorf("failed to unmarshal config: %s", err))
	}
	return c
}
