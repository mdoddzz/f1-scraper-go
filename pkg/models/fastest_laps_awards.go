package models

// FastestLapsAwards : model for the fastest lap for each race
type FastestLapsAwards struct {
	ID     interface{} `json:"id" bson:"_id,omitempty"`
	RaceId interface{} `json:"race_id" bson:"race_id"`
	Driver Driver      `json:"driver" bson:"driver"`
	Car    string      `json:"car" bson:"car"`
	Time   F1Time      `json:"time" bson:"time"`
}
