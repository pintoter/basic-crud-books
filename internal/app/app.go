package app

import (
	"github.com/pintoter/basic-crud-books/internal/config"
	"github.com/pintoter/basic-crud-books/internal/repository"
	"github.com/pintoter/basic-crud-books/internal/server"
	"github.com/pintoter/basic-crud-books/internal/service"
	"github.com/pintoter/basic-crud-books/internal/transport"
	"github.com/pintoter/basic-crud-books/pkg/database/postgres"
	"github.com/pintoter/basic-crud-books/pkg/hash"
	"github.com/pintoter/basic-crud-books/pkg/logger"
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	shutdownTimeout = 5 * time.Second
)

func Run(configDir, configName, envFile string) error {
	cfg, err := InitConfig(configDir, configName, envFile)
	if err != nil {
		logger.Fatal("config-error", err)
		return err
	}

	db, err := postgres.NewDB(&cfg.DB)
	if err != nil {
		logger.Fatal("database-connection", err)
		return err
	}
	defer db.Close()
	logger.Info("database", "postgres connected")

	service := service.New(service.Deps{
		Repos:  repository.New(db),
		Hasher: hash.NewHasher(cfg.Auth.Salt),
	})

	handler := transport.NewHandler(service) // у макса по другому, отдельная функция

	server := server.New(handler.InitRoutes(), cfg.HTTPServer.Host, cfg.HTTPServer.Port)
	go func() {
		if err := server.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal("server", err)
		}
	}()
	logger.Info("server", "server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	logger.Info("server", "shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := server.Stop(ctx); err != nil {
		return err
	}
	logger.Info("server", "server stopped")

	return nil
}

func InitConfig(configDir, configName, envFile string) (*config.Config, error) {
	err := config.InitENV(envFile)
	if err != nil {
		return nil, err
	}

	cfg, err := config.New(configDir, configName)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
