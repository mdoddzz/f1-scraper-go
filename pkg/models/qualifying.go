package models

// Qualifying : model for the qualifying results
type Qualifying struct {
	ID       interface{} `json:"id" bson:"_id,omitempty"`
	RaceId   interface{} `json:"race_id" bson:"race_id"`
	Session  string      `json:"session,omitempty" bson:"session,omitempty"`
	Position int         `json:"position" bson:"position"`
	Number   int         `json:"number" bson:"number"`
	Driver   Driver      `json:"driver" bson:"driver"`
	Car      string      `json:"car" bson:"car"`
	Q1       *F1Time     `json:"q1,omitempty" bson:"q1,omitempty"`
	Q2       *F1Time     `json:"q2,omitempty" bson:"q2,omitempty"`
	Q3       *F1Time     `json:"q3,omitempty" bson:"q3,omitempty"`
	Time     *F1Time     `json:"time,omitempty" bson:"time,omitempty"`
	Laps     int         `json:"laps,omitempty" bson:"laps,omitempty"`
}

// QualifyingService : interface for the qualifying result model
type QualifyingService interface {

	// GetQualifyingResult : Get qualifying result by race ID
	GetQualifyingResult(raceId interface{}) (*[]Qualifying, error)

	// AddQualifyingResult : Add a new qualifying result
	AddQualifyingResult(qualifying Qualifying) error
}
