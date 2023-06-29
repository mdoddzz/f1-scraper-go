package models

// StartingGrid : model for the race starting grid
type StartingGrid struct {
	ID       interface{} `json:"id" bson:"_id,omitempty"`
	RaceId   interface{} `json:"race_id" bson:"race_id"`
	Position int         `json:"position" bson:"position"`
	Number   int         `json:"number" bson:"number"`
	Driver   Driver      `json:"driver" bson:"driver"`
	Car      string      `json:"car" bson:"car"`
	Time     F1Time      `json:"time" bson:"time"`
}

// StartingGridService : interface for the starting grid model
type StartingGridService interface {

	// GetStartingGrid : Get starting grid of a race
	GetStartingGrid(raceId interface{}) (*[]StartingGrid, error)

	// AddStartingGrid : Add a new starting grid
	AddStartingGrid(startingGrid StartingGrid) error
}
