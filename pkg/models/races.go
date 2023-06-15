package models

import "time"

// Race : model for the race details
type Race struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	UrlID     int       `json:"url_id" bson:"url_id"`
	GrandPrix string    `json:"grand_prix" bson:"grand_prix"`
	Date      time.Time `json:"date" bson:"date"`
	Winner    Driver    `json:"winner" bson:"winner"`
	Car       string    `json:"car" bson:"car"`
	Laps      int       `json:"laps" bson:"laps"`
	Time      time.Time `json:"time" bson:"time"`
}

// RacesService : interface for the race details model
type RacesService interface {

	// Get all races
	GetRaces() (*[]Race, error)

	// Get a race by URL ID
	GetRaceByUrlId(id int) (*Race, error)

	// Get a race by ID
	GetRaceById(id string) (*Race, error)

	// Add a new race
	AddRace(race Race) error
}
