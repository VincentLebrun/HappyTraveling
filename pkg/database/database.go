package database

import (
	"ApiHappyTravel/internal/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Connect(cfg config.DatabaseConfig) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		log.Printf("Cannot open sql %s", err)
	}

	if err = db.Ping(); err != nil {
		log.Printf("Cannot ping the adress %s", err)
	}

	return db, nil
}
