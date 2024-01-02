package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/soda-zero/vegandb/internal/config"
)

func ConnectDB() (*sql.DB, error) {
	cfg := config.GetConfig(config.Dev)
	db, err := sql.Open("sqlite3", cfg.DatabaseConnectionString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to SQLite database: %v", err)
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("unable to ping dabatase: %v", err)

	}
	return db, nil
}
