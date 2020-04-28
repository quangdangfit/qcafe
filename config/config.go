package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Schema struct {
	MongoDB struct {
		Host         string `mapstructure:"host"`
		Port         int    `mapstructure:"port"`
		AuthDatabase string `mapstructure:"auth_database"`
		Database     string `mapstructure:"database"`
		User         string `mapstructure:"user"`
		Password     string `mapstructure:"password"`
	}
}

var Config Schema

func init() {
	config := viper.New()
	config.SetConfigName("config")
	config.AddConfigPath("./config") // Look for config in the working directory.

	err := config.ReadInConfig() // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	err = config.Unmarshal(&Config)
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
