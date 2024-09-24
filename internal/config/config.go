package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port string
}

func NewConfig(configPath string) (*Config, error) {
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	return &Config{
		Port: ":" + viper.GetString("port"),
	}, nil
}
