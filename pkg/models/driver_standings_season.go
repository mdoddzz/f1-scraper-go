package models

// DriverStandingsSeason : model for driver standing for the whole season
type DriverStandingsSeason struct {
	Year        int
	Position    int
	Driver      Driver
	Nationality string
	Car         string
	Points      float64
}
