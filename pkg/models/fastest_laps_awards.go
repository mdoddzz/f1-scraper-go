package models

// FastestLapAward : model for the fastest lap for each race
type FastestLapAward struct {
	ID     interface{} `json:"id" bson:"_id,omitempty"`
	RaceId interface{} `json:"race_id" bson:"race_id"`
	Driver Driver      `json:"driver" bson:"driver"`
	Car    string      `json:"car" bson:"car"`
	Time   F1Time      `json:"time" bson:"time"`
}

// FastestLapAwardService : interface for the fastest lap result model
type FastestLapAwardService interface {

	// Get practice result
	GetFastestLapAward(raceId interface{}) (*FastestLapAward, error)

	// Add a new practice result
	AddFastestLapAward(fastestLapAward FastestLapAward) error
}
