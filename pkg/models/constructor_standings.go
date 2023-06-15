package models

// ConstructorStandings : model for the constructor standings per team/ race
type ConstructorStandings struct {
	RaceId string  `json:"race_id" bson:"race_id"`
	Points float64 `json:"points" bson:"points"`
}
