package models

import "time"

type Personne struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Age      string `json:"age"`
	Phone    string `json:"phone"`
	Mail     string `json:"mail"`
	Password string `json:"password" bson:"password" validate:"required=true"`
}

type Event struct {
	IdEvent int       `json:"id_event"`
	Name    string    `json:"name"`
	Date    time.Time `json:"date"`
	Town    []Town    `json:"town"`
}

type Town struct {
	IdTown int    `json:"id_town"`
	Name   string `json:"name"`
}

type Travel struct {
	IdTravel    int       `json:"id_travel"`
	Driver      string    `json:"driver"`
	Start       Town      `json:"start"`
	Destination Town      `json:"destination"`
	Hourly      time.Time `json:"hourly"`
}
