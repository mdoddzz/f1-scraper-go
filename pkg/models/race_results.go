package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// RaceResult : model for the race results
type RaceResult struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	RaceId   primitive.ObjectID `json:"race_id" bson:"race_id"`
	Position int                `json:"position" bson:"position"`
	Number   int                `json:"number" bson:"number"`
	Driver   Driver             `json:"driver" bson:"driver"`
	Car      string             `json:"car" bson:"car"`
	Laps     int                `json:"laps" bson:"laps"`
	Time     string             `json:"time" bson:"time"`
	Points   float64            `json:"points" bson:"points"`
}

// RaceResultService : interface for the race result model
type RaceResultService interface {

	// Get race result
	GetRaceResult(raceId string) (*[]RaceResult, error)

	// Get race results of a driver
	GetRaceResultDriver(driver string) (*[]RaceResult, error)

	// Add a new race result
	AddRaceResult(raceResult RaceResult) error
}
