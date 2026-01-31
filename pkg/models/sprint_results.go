package models

// SprintResult : model for the sprint race results
type SprintResult struct {
	ID       interface{} `json:"id" bson:"_id,omitempty"`
	RaceId   interface{} `json:"race_id" bson:"race_id"`
	Position interface{} `json:"position" bson:"position"`
	Number   int         `json:"number" bson:"number"`
	Driver   Driver      `json:"driver" bson:"driver"`
	Car      string      `json:"car" bson:"car"`
	Laps     int         `json:"laps" bson:"laps"`
	Time     string      `json:"time" bson:"time"`
	Points   float64     `json:"points" bson:"points"`
}

// SprintResultService : interface for the sprint result model
type SprintResultService interface {

	// GetSprintResult : Get sprint result
	GetSprintResult(raceId interface{}) (*[]SprintResult, error)

	// GetSprintResultDriver : Get sprint results of a driver
	GetSprintResultDriver(driver string) (*[]SprintResult, error)

	// AddSprintResult : Add a new sprint result
	AddSprintResult(sprintResult SprintResult) error
}
