package models

import "time"

// Qualifying : model for the qualifying results
type Qualifying struct {
	RaceId   string    `json:"race_id" bson:"race_id"`
	Position int       `json:"position" bson:"position"`
	Number   int       `json:"number" bson:"number"`
	Driver   Driver    `json:"driver" bson:"driver"`
	Q1       time.Time `json:"q1" bson:"q1"`
	Q2       time.Time `json:"q2" bson:"q2"`
	Q3       time.Time `json:"q3" bson:"q3"`
	Laps     int       `json:"laps" bson:"laps"`
}

// QualifyingService : interface for the qualifying result model
type QualifyingService interface {

	// Get qualifying result by race ID
	GetQualifyingResult(raceId string) (*[]Qualifying, error)

	// Add a new qualifying result
	AddQualifyingResult(result Qualifying) error
}
