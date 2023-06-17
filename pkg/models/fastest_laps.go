package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// FastestLaps : model for the fastest laps results
type FastestLaps struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	RaceId    primitive.ObjectID `json:"race_id" bson:"race_id"`
	Position  int                `json:"position" bson:"position"`
	Number    int                `json:"number" bson:"number"`
	Driver    Driver             `json:"driver" bson:"driver"`
	Lap       int                `json:"lap" bson:"lap"`
	TimeOfDay F1Time             `json:"time_of_day" bson:"time_of_day"`
	Time      F1Time             `json:"time" bson:"time"`
	AvgSpeed  float64            `json:"avg_speed" bson:"avg_speed"`
}
