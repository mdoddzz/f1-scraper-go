package models

// RaceResult : model for the race results
type RaceResult struct {
	ID       interface{} `json:"id" bson:"_id,omitempty"`
	RaceId   interface{} `json:"race_id" bson:"race_id"`
	Position interface{} `json:"position" bson:"position"`
	Number   int         `json:"number" bson:"number"`
	Driver   Driver      `json:"driver" bson:"driver"`
	Car      string      `json:"car" bson:"car"`
	Laps     int         `json:"laps" bson:"laps"`
	Time     string      `json:"time" bson:"time"`
	Points   float64     `json:"points" bson:"points"`
}

// RaceResultService : interface for the race result model
type RaceResultService interface {

	// GetRaceResult : Get race result
	GetRaceResult(raceId interface{}) (*[]RaceResult, error)

	// GetRaceResultDriver : Get race results of a driver
	GetRaceResultDriver(driver string) (*[]RaceResult, error)

	// AddRaceResult : Add a new race result
	AddRaceResult(raceResult RaceResult) error
}
