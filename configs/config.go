package configs

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading .env file: %v", err)
	}

	// Default values
	viper.SetDefault("APP_PORT", "8080")
}

func GetConfig(key string) string {
	return viper.GetString(key)
}
