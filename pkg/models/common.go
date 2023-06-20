package models

import "time"

// Driver : model for driver for the whole season
type Driver struct {
	FullName       string `json:"full_name" bson:"full_name"`
	FirstName      string `json:"first_name" bson:"first_name"`
	LastName       string `json:"last_name" bson:"last_name"`
	NameIdentifier string `json:"name_identifier" bson:"name_identifier"`
}

// F1Time : model for saving time in both datetime format and a string
type F1Time struct {
	String   string    `json:"string,omitempty" bson:"string,omitempty"`
	DateTime time.Time `json:"datetime,omitempty" bson:"datetime,omitempty"`
}
