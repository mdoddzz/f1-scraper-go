package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Practice : model for the practice results
type Practice struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	RaceId        primitive.ObjectID `json:"race_id" bson:"race_id"`
	SessionNumber int                `json:"session_number" bson:"session_number"`
	Position      int                `json:"position" bson:"position"`
	Driver        Driver             `json:"driver" bson:"driver"`
	Car           string             `json:"car" bson:"car"`
	Time          F1Time             `json:"time" bson:"time"`
	Gap           string             `json:"gap" bson:"gap"`
	Laps          int                `json:"laps" bson:"laps"`
}

// PracticeService : interface for the practice result model
type PracticeService interface {

	// Get practice result
	GetPracticeResult(raceId string, session int) (*[]Practice, error)

	// Add a new practice result
	AddPractice(practice Practice) error
}
