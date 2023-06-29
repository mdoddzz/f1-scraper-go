package models

// DriverStandings : model for the driver standings per driver/ race
type DriverStandings struct {
	ID           interface{} `json:"id" bson:"_id,omitempty"`
	RaceId       interface{} `json:"race_id" bson:"race_id"`
	Driver       Driver      `json:"driver" bson:"driver"`
	Car          string      `json:"car" bson:"car"`
	RacePosition int         `json:"race_position" bson:"race_position"`
	Points       float64     `json:"points" bson:"points"`
}

// DriverStandingsService : interface for the driver standings
type DriverStandingsService interface {

	// GetDriverStandings : Get driver standings for a race
	GetDriverStandings(raceId interface{}) (*[]DriverStandings, error)

	// AddDriverStandings : Add a new driver standings
	AddDriverStandings(standing DriverStandings) error
}
