package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"

	"fmt"
)

const (
	driverName = "postgres"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

// Host     string = "localhost"
// Port     int = 5432
// User     string = "postgres"
// Password string = "123qweASD"
// Name   string = "books"
// SSLMode  string = "disable"

func NewDB(cfg *DBConfig) (*sql.DB, error) {
	mycfg := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Name, cfg.Password, cfg.SSLMode)

	db, err := sql.Open(driverName, mycfg)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
