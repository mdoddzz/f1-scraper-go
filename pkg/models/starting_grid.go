package models

// StartingGrid : model for the race starting grid
type StartingGrid struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	RaceId   string `json:"race_id" bson:"race_id"`
	Position int    `json:"position" bson:"position"`
	Number   int    `json:"number" bson:"number"`
	Driver   string `json:"driver" bson:"driver"`
	Car      string `json:"car" bson:"car"`
	Time     F1Time `json:"time" bson:"time"`
}

// StartingGridService : interface for the starting grid model
type StartingGridService interface {

	// Get starting grid of a race
	GetStartingGrid(raceId string) (*[]StartingGrid, error)

	// Add a new starting grid
	AddStartingGrid(startingGrid StartingGrid) error
}
