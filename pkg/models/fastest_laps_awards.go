package models

import "time"

// FastestLapsAwards : model for the fastest lap for each race
type FastestLapsAwards struct {
	RaceId string    `json:"race_id" bson:"race_id"`
	Driver Driver    `json:"driver" bson:"driver"`
	Car    string    `json:"car" bson:"car"`
	Time   time.Time `json:"time" bson:"time"`
}
