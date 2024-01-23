package handlers

import (
	"ApiHappyTravel/pkg/models"
	"ApiHappyTravel/pkg/repository"
	"encoding/json"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Cannot decode user data: %s", err)
		http.Error(w, "Invalid user data", http.StatusBadRequest)
		return
	}

	err = repository.InsertUserIntoDB(user)
	if err != nil {
		log.Printf("Cannot insert user into database: %s", err)
		http.Error(w, "Error inserting user into database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
