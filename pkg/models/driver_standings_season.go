package models

// DriverStandingsSeason : model for driver standing for the whole season
type DriverStandingsSeason struct {
	ID          interface{} `json:"id" bson:"_id,omitempty"`
	Year        int         `json:"year" bson:"year"`
	Position    int         `json:"position" bson:"position"`
	Driver      Driver      `json:"driver" bson:"driver"`
	Nationality string      `json:"nationality" bson:"nationality"`
	Car         string      `json:"car" bson:"car"`
	Points      float64     `json:"points" bson:"points"`
}

// DriverStandingSeasonService : interface for the driver standing for the whole season model
type DriverStandingSeasonService interface {

	// GetDriverStandingsByYear : Get driver standings by Year
	GetDriverStandingsByYear(year int) (*[]DriverStandingsSeason, error)

	// AddDriverStandingSeason : Add a new driver standing
	AddDriverStandingSeason(driverStandingsSeason DriverStandingsSeason) error
}
