package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Port string `mapstructure:"PORT"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
	User     string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"`
	Db       string `mapstructure:"DBNAME"`
}

type Config struct {
	App AppConfig      `mapstructure:"APP"`
	Db  DatabaseConfig `mapstructure:"DATABASE"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Warning: Could not read config file: %v\n", err)
	}

	config := new(Config)

	err := viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
