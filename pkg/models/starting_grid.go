package models

import "time"

// StartingGrid : model for the race starting grid
type StartingGrid struct {
	RaceId   string
	Position int
	Number   int
	Driver   string
	Car      string
	Time     time.Duration
}

// StartingGridService : interface for the starting grid model
type StartingGridService interface {

	// Get starting grid of a race
	GetStartingGrid(raceId string) (*[]StartingGrid, error)

	// Add a new starting grid
	AddStartingGrid(grid StartingGrid) error
}
