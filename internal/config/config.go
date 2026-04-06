package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	AppName   string
	AppEnv    string
	Port      string
	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	DBName    string
	JWTSecret string
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	viper.AutomaticEnv()

	cfg := &Config{
		AppName:   viper.GetString("APP_NAME"),
		AppEnv:    viper.GetString("APP_ENV"),
		Port:      viper.GetString("PORT"),
		DBHost:    viper.GetString("DB_HOST"),
		DBPort:    viper.GetString("DB_PORT"),
		DBUser:    viper.GetString("DB_USER"),
		DBPass:    viper.GetString("DB_PASSWORD"),
		DBName:    viper.GetString("DB_NAME"),
		JWTSecret: viper.GetString("JWT_SECRET"),
	}

	if cfg.Port == "" {
		log.Fatal("PORT not set")
	}

	return cfg
}

