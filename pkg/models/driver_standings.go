package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// DriverStandings : model for the driver standings per driver/ race
type DriverStandings struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	RaceId       primitive.ObjectID `json:"race_id" bson:"race_id"`
	Driver       Driver             `json:"driver" bson:"driver"`
	Car          string             `json:"car" bson:"car"`
	RacePosition int                `json:"race_position" bson:"race_position"`
	Points       float64            `json:"points" bson:"points"`
}
