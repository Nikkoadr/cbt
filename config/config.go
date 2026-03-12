package config

import (
	"github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
	AppName  string `mapstructure:"APP_NAME"`
	AppPort  string `mapstructure:"APP_PORT"`
	DBHost   string `mapstructure:"DB_HOST"`
	DBPort   string `mapstructure:"DB_PORT"`
	DBUser   string `mapstructure:"DB_USER"`
	DBPass   string `mapstructure:"DB_PASS"`
	DBName   string `mapstructure:"DB_NAME"`
	JWTSecret string `mapstructure:"JWT_SECRET"`
	JWTExpire int    `mapstructure:"JWT_EXPIRE"`
}

// LoadConfig loads configuration from .env file
func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
