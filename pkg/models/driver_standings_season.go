package models

// DriverStandingsSeason : model for driver standing for the whole season
type DriverStandingsSeason struct {
	ID          string  `json:"id" bson:"_id,omitempty"`
	Year        int     `json:"year" bson:"year"`
	Position    int     `json:"position" bson:"position"`
	Driver      Driver  `json:"driver" bson:"driver"`
	Nationality string  `json:"nationality" bson:"nationality"`
	Car         string  `json:"car" bson:"car"`
	Points      float64 `json:"points" bson:"points"`
}
