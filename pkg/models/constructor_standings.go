package models

// ConstructorStandings : model for the constructor standings per team/ race
type ConstructorStandings struct {
	ID     interface{} `json:"id" bson:"_id,omitempty"`
	RaceId string      `json:"race_id" bson:"race_id"`
	Points float64     `json:"points" bson:"points"`
}
