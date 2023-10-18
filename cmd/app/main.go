package main

import (
	"github.com/pintoter/basic-crud-books/internal/app"
	"github.com/pintoter/basic-crud-books/pkg/logger"
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
