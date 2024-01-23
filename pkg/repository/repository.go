package repository

import (
	"ApiHappyTravel/internal/config"
	"ApiHappyTravel/pkg/models"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func getDbConnection() *sql.DB {
	cfg := config.LoadConfiguration()

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	return db
}
func InsertUserIntoDB(user models.User) error {
	db := getDbConnection()
	if db == nil {
		log.Println("Failed to establish database connection")
		return fmt.Errorf("database connection error")
	}
	defer db.Close()

	query := `INSERT INTO users (name, surname, age, phone, mail, password) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(query, user.Name, user.Surname, user.Age, user.Phone, user.Mail, user.Password)
	if err != nil {
		log.Printf("Error inserting user into database: %s", err)
		return err
	}

	return nil
}
