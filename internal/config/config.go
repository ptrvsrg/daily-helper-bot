package config

import (
	"path/filepath"

	"github.com/spf13/viper"

	"daily-helper-bot/internal/log"
)

type BotConfig struct {
	Token string `mapstructure:"token"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"dbname"`
}

type APIServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type AllConfig struct {
	Bot       BotConfig       `mapstructure:"bot"`
	DB        DBConfig        `mapstructure:"db"`
	APIServer APIServerConfig `mapstructure:"api_server"`
}

var Config AllConfig

func LoadConfig(filePath string) {
	viper.SetConfigFile(filePath)
	viper.SetConfigType(filepath.Ext(filePath)[1:])

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Logger.Fatalf("Config error: %v", err)
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Logger.Fatalf("Config error: %v", err)
	}
}
