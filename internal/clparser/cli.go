package clparser

import (
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Parse() {
	pflag.StringP("config", "c", "config/config.json", "Config file path")
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		log.Fatalf("Command line parsser error: %v", err)
	}
}

func GetConfigFilePath() string {
	return viper.GetString("config")
}
