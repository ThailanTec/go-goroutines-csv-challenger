package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBUsername string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
}

func LoadConfig() Config {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	config := Config{
		DBUsername: viper.GetString("DB_USERNAME"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetString("DB_PORT"),
	}

	return config
}
