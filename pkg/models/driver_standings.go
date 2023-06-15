package models

// DriverStandings : model for the driver standings per driver/ race
type DriverStandings struct {
	RaceId       string
	Driver       Driver
	Car          string
	RacePosition int
	Points       float64
}
