package models

// ConstructorStandingsSeason : model for constructor standing for the whole season
type ConstructorStandingsSeason struct {
	Position int     `json:"position" bson:"position"`
	Team     string  `json:"team" bson:"team"`
	Points   float64 `json:"points" bson:"points"`
}
