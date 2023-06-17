package models

// PitStop : model for the pitstop results
type PitStop struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	RaceId    string `json:"race_id" bson:"race_id"`
	Stops     int    `json:"stops" bson:"stops"`
	Number    int    `json:"number" bson:"number"`
	Driver    Driver `json:"driver" bson:"driver"`
	Lap       int    `json:"lap" bson:"lap"`
	TimeOfDay F1Time `json:"time_of_day" bson:"time_of_day"`
	Time      F1Time `json:"time" bson:"time"`
	Total     F1Time `json:"total" bson:"total"`
}
