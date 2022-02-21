package config

import (
	"github.com/spf13/viper"
)

// Config - Service configuration structure
type Config struct {
	AppName        string
	DefaultAppPort string
	DB             DataBase
}
type DataBase struct {
	Host     string
	Username string
	Password string
	Database string
	Port     string
}

// Settings - Actual service settings
var Settings Config

// ReadConfig provide application configuration from ENV variables
func ReadConfig() Config {

	viper.AutomaticEnv()

	viper.SetEnvPrefix("APP")
	viper.SetDefault("NAME", "Todo")
	viper.SetDefault("PORT", "8080")
	// PRODUCTION DB

	viper.SetDefault("POSTGRES_HOSTNAME", "localhost")
	viper.SetDefault("POSTGRES_USERNAME", "postgres")
	viper.SetDefault("POSTGRES_PASSWORD", "123")
	viper.SetDefault("POSTGRES_DATABASE", "postgres")
	viper.SetDefault("POSTGRES_PORT", "5432")

	return Config{
		AppName:        viper.GetString("NAME"),
		DefaultAppPort: viper.GetString("PORT"),
		DB: DataBase{
			Host:     viper.GetString("POSTGRES_HOSTNAME"),
			Username: viper.GetString("POSTGRES_USERNAME"),
			Password: viper.GetString("POSTGRES_PASSWORD"),
			Database: viper.GetString("POSTGRES_DATABASE"),
			Port:     viper.GetString("POSTGRES_PORT"),
		},
	}
}
