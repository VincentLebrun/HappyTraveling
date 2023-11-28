package main

import (
	"ApiHappyTravel/internal/config"
	"ApiHappyTravel/pkg/database"
	"log"
)

func main() {
	cfg := config.LoadConfiguration()

	db, err := database.Connect(cfg.Database)
	if nil != err {
		log.Fatal("Could not connect ", err)
	}
	defer db.Close()
}
