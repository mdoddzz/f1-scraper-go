package models

// Practice : model for the practice results
type Practice struct {
	ID       interface{} `json:"id" bson:"_id,omitempty"`
	RaceId   interface{} `json:"race_id" bson:"race_id"`
	Session  string      `json:"session" bson:"session"`
	Position int         `json:"position" bson:"position"`
	Number   int         `json:"number" bson:"number"`
	Driver   Driver      `json:"driver" bson:"driver"`
	Car      string      `json:"car" bson:"car"`
	Time     F1Time      `json:"time" bson:"time"`
	Gap      string      `json:"gap" bson:"gap"`
	Laps     int         `json:"laps,omitempty" bson:"laps,omitempty"`
}

// PracticeService : interface for the practice result model
type PracticeService interface {

	// Get practice result
	GetPracticeResult(raceId interface{}, session string) (*[]Practice, error)

	// Add a new practice result
	AddPractice(practice Practice) error
}
