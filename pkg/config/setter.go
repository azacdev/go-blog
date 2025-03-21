package config

import (
	"log"

	"github.com/spf13/viper"
)

func Set() {
	viper.SetConfigName("config") // name of config file (without extension
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("config") // path to look for the config file in // optionally look for config in the working directory
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading the config")
	}

	err := viper.Unmarshal(&configurations)
	if err != nil {
		log.Fatal("unable to decode into struct, %v", err)
	}
}
