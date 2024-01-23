package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"ApiHappyTravel/internal/config"
	"log"
)

func main() {
	conf := config.LoadConfiguration()

	db, err := sql.Open("mysql", conf.Database.DSN())
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()
}
