package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// ConstructorStandings : model for the constructor standings per team/ race
type ConstructorStandings struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	RaceId primitive.ObjectID `json:"race_id" bson:"race_id"`
	Points float64            `json:"points" bson:"points"`
}
