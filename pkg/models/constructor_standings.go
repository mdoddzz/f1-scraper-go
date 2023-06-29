package models

// ConstructorStandings : model for the constructor standings per team/ race
type ConstructorStandings struct {
	ID     interface{} `json:"id" bson:"_id,omitempty"`
	RaceId string      `json:"race_id" bson:"race_id"`
	Team   string      `json:"team" bson:"team"`
	Points float64     `json:"points" bson:"points"`
}

// ConstructorStandingsService : interface for the constructor standings
type ConstructorStandingsService interface {

	// GetConstructorStandings : Get constructor standings for a race
	GetConstructorStandings(raceId interface{}) (*[]ConstructorStandings, error)

	// AddConstructorStandings : Add a new constructor standings
	AddConstructorStandings(standing ConstructorStandings) error
}
