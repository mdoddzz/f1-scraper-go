package models

// Driver : model for driver for the whole season
type Driver struct {
	FullName       string `json:"full_name" bson:"full_name"`
	FirstName      string `json:"first_name" bson:"first_name"`
	LastName       string `json:"last_name" bson:"last_name"`
	NameIdentifier string `json:"name_identifier" bson:"name_identifier"`
}
