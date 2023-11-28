package database

import (
	"ApiHappyTravel/internal/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connect(cfg config.DatabaseConfig) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
