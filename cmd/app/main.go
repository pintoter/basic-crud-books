package main

import (
	"books-app/internal/app"
	"books-app/pkg/logger"
)

const (
	ConfigPath = "configs"
	ConfigFile = "main"
	EnvFile    = ".env"
)

func main() {
	if err := app.Run(ConfigPath, ConfigFile, EnvFile); err != nil {
		logger.Error(err.Error())
	}
}
