package models

// FastestLaps : model for the fastest laps in a race
type FastestLaps struct {
	ID        interface{} `json:"id" bson:"_id,omitempty"`
	RaceId    interface{} `json:"race_id" bson:"race_id"`
	Position  int         `json:"position" bson:"position"`
	Number    int         `json:"number" bson:"number"`
	Driver    Driver      `json:"driver" bson:"driver"`
	Lap       int         `json:"lap" bson:"lap"`
	TimeOfDay F1Time      `json:"time_of_day" bson:"time_of_day"`
	Time      F1Time      `json:"time" bson:"time"`
	AvgSpeed  float64     `json:"avg_speed" bson:"avg_speed"`
}

// FastestLapsService : interface for the fastest laps in a race
type FastestLapsService interface {

	// GetFastestLaps : Get fastest laps for a race
	GetFastestLaps(raceId interface{}) (*[]FastestLaps, error)

	// AddFastestLaps : Add a new fastest lap
	AddFastestLaps(fastestLaps FastestLaps) error
}
