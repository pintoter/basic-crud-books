package config

import (
	"books-app/pkg/database/postgres"
	"books-app/pkg/logger"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

type Config struct {
	HTTPServer Server
	Auth       AuthConfig
	DB         postgres.DBConfig
}

type (
	Server struct {
		Host string
		Port int
	}

	AuthConfig struct {
		Salt   string
		Secret string
	}
)

func InitENV(filename string) error {
	return godotenv.Load(filename)
}

func New(configPath, configName string) (*Config, error) {
	cfg := &Config{} // need refactor

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	if err := envconfig.Process("db", &cfg.DB); err != nil {
		logger.Fatal("db config", err.Error())
	}

	if err := envconfig.Process("hash", &cfg.Auth); err != nil {
		logger.Fatal("hash envs", err.Error())
	}

	return cfg, nil
}
