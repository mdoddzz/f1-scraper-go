package models

// SprintShootout : model for the sprint shootout/qualifying results
type SprintShootout struct {
	ID       interface{} `json:"id" bson:"_id,omitempty"`
	RaceId   interface{} `json:"race_id" bson:"race_id"`
	Session  string      `json:"session,omitempty" bson:"session,omitempty"`
	Position int         `json:"position" bson:"position"`
	Number   int         `json:"number" bson:"number"`
	Driver   Driver      `json:"driver" bson:"driver"`
	Car      string      `json:"car" bson:"car"`
	Time     *F1Time     `json:"time,omitempty" bson:"time,omitempty"`
	Laps     int         `json:"laps,omitempty" bson:"laps,omitempty"`
}

// SprintShootoutService : interface for the sprint shootout result model
type SprintShootoutService interface {

	// GetSprintShootout : Get sprint shootout result by race ID
	GetSprintShootout(raceId interface{}) (*[]SprintShootout, error)

	// AddSprintShootout : Add a new sprint shootout result
	AddSprintShootout(sprintShootout SprintShootout) error
}
