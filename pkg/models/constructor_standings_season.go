package models

// ConstructorStandingsSeason : model for constructor standing for the whole season
type ConstructorStandingsSeason struct {
	ID       interface{} `json:"id" bson:"_id,omitempty"`
	Year     int         `json:"year" bson:"year"`
	Position int         `json:"position" bson:"position"`
	Team     string      `json:"team" bson:"team"`
	Points   float64     `json:"points" bson:"points"`
}

// ConstructorStandingsSeasonService : interface for the constructor standing for the whole season model
type ConstructorStandingsSeasonService interface {

	// Get constructor standings by Year
	GetConstructorStandingsSeason(year int) (*[]ConstructorStandingsSeason, error)

	// Add a new constructor standing
	AddConstructorStandingsSeason(constructorStandingsSeason ConstructorStandingsSeason) error
}
