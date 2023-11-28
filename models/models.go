package models

import "time"

type Personne struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     string `json:"age"`
}

type Event struct {
	IdEvent int       `json:"id_event"`
	Name    string    `json:"name"`
	Date    time.Time `json:"date"`
	Town    Town      `json:"town"`
}

type Town struct {
	IDVille int    `json:"id_town"`
	Name    string `json:"name"`
}

type Trajet struct {
	IDTrajet    int    `json:"id_trajet"`
	Conducteur  string `json:"conducteur"`
	Depart      Town   `json:"depart"`
	Destination Town   `json:"destination"`
	Horaire     string `json:"horaire"`
}
