package mongo

import (
	"context"
	"errors"
	"time"

	"github.com/mdoddzz/f1-scraper-go/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

// GetRaces : Get all races
func (s *Storage) GetRaces() (*[]models.Race, error) {

	return &[]models.Race{}, nil
}

// GetRaceByUrlId : Get a race by date
func (s *Storage) GetRaceByUrlId(id int) (*models.Race, error) {

	filter := bson.D{{Key: "url_id", Value: id}}

	cursor, err := s.race.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []models.Race
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	if len(results) == 0 {
		return nil, errors.New("race not found")
	}

	return &results[0], nil
}

// GetRaceByGPDate : Get a race by GP and year
func (s *Storage) GetRaceByGPDate(gp string, date time.Time) (*models.Race, error) {

	filter := bson.D{{Key: "grand_prix", Value: gp}, {Key: "date.datetime", Value: date}}

	cursor, err := s.race.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []models.Race
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	if len(results) == 0 {
		return nil, errors.New("race not found")
	}

	return &results[0], nil
}

// GetRaceById : Get a race by ID
func (s *Storage) GetRaceById(id interface{}) (*models.Race, error) {

	// Check ID a primitive.ObjectId

	return &models.Race{}, nil
}

// AddRace : Add a new race
func (s *Storage) AddRace(race models.Race) error {

	_, err := s.race.InsertOne(context.TODO(), race)
	if err != nil {
		return err
	}

	return nil

}
