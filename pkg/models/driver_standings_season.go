package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// DriverStandingsSeason : model for driver standing for the whole season
type DriverStandingsSeason struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Year        int                `json:"year" bson:"year"`
	Position    int                `json:"position" bson:"position"`
	Driver      Driver             `json:"driver" bson:"driver"`
	Nationality string             `json:"nationality" bson:"nationality"`
	Car         string             `json:"car" bson:"car"`
	Points      float64            `json:"points" bson:"points"`
}

// DriverStandingSeasonService : interface for the driver standing for the whole season model
type DriverStandingSeasonService interface {

	// Get driver standings by Year
	GetDriverStandings(year int) (*[]DriverStandingsSeason, error)

	// Add a new driver standing
	AddDriverStandingSeason(driverStandingsSeason DriverStandingsSeason) error
}
