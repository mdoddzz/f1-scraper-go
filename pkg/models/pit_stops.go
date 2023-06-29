package models

// PitStop : model for the pit stop results
type PitStop struct {
	ID        interface{} `json:"id" bson:"_id,omitempty"`
	RaceId    interface{} `json:"race_id" bson:"race_id"`
	Stops     int         `json:"stops" bson:"stops"`
	Number    int         `json:"number" bson:"number"`
	Driver    Driver      `json:"driver" bson:"driver"`
	Car       string      `json:"car" bson:"car"`
	Lap       int         `json:"lap" bson:"lap"`
	TimeOfDay F1Time      `json:"time_of_day" bson:"time_of_day"`
	Time      F1Time      `json:"time" bson:"time"`
	Total     F1Time      `json:"total" bson:"total"`
}

// PitStopService : interface for the pit stop model
type PitStopService interface {

	// GetPitStops : Get pit stop for a race
	GetPitStops(raceId interface{}) (*[]PitStop, error)

	// AddPitStop : Add a new pit stop
	AddPitStop(pitStop PitStop) error
}
