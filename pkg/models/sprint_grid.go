package models

// SprintGrid : model for the sprint race starting grid
type SprintGrid struct {
	ID       interface{} `json:"id" bson:"_id,omitempty"`
	RaceId   interface{} `json:"race_id" bson:"race_id"`
	Position int         `json:"position" bson:"position"`
	Number   int         `json:"number" bson:"number"`
	Driver   Driver      `json:"driver" bson:"driver"`
	Car      string      `json:"car" bson:"car"`
	Time     F1Time      `json:"time" bson:"time"`
}

// SprintGridService : interface for the sprint grid model
type SprintGridService interface {

	// GetSprintGrid : Get sprint grid of a race
	GetSprintGrid(raceId interface{}) (*[]SprintGrid, error)

	// AddSprintGrid : Add a new sprint grid
	AddSprintGrid(sprintGrid SprintGrid) error
}
