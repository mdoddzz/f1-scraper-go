package models

// Practice : model for the practice results
type Practice struct {
	ID            interface{} `json:"id" bson:"_id,omitempty"`
	RaceId        interface{} `json:"race_id" bson:"race_id"`
	SessionNumber int         `json:"session_number" bson:"session_number"`
	Position      int         `json:"position" bson:"position"`
	Driver        Driver      `json:"driver" bson:"driver"`
	Car           string      `json:"car" bson:"car"`
	Time          F1Time      `json:"time" bson:"time"`
	Gap           string      `json:"gap" bson:"gap"`
	Laps          int         `json:"laps" bson:"laps"`
}

// PracticeService : interface for the practice result model
type PracticeService interface {

	// Get practice result
	GetPracticeResult(raceId interface{}, session int) (*[]Practice, error)

	// Add a new practice result
	AddPractice(practice Practice) error
}
