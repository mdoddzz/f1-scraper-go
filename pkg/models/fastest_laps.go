package models

import "time"

// FastestLaps : model for the fastest laps results
type FastestLaps struct {
	RaceId    string    `json:"race_id" bson:"race_id"`
	Position  int       `json:"position" bson:"position"`
	Number    int       `json:"number" bson:"number"`
	Driver    Driver    `json:"driver" bson:"driver"`
	Lap       int       `json:"lap" bson:"lap"`
	TimeOfDay time.Time `json:"time_of_day" bson:"time_of_day"`
	Time      time.Time `json:"time" bson:"time"`
	AvgSpeed  float64   `json:"avg_speed" bson:"avg_speed"`
}
