package mongo

import (
	"context"

	"github.com/mdoddzz/f1-scraper-go/pkg/models"
)

// GetRaceResult : Get race result
func (s *Storage) GetRaceResult(raceId interface{}) (*[]models.RaceResult, error) {

	// Check ID a primitive.ObjectId

	return &[]models.RaceResult{}, nil
}

// GetRaceResultDriver : Get race results of a driver
func (s *Storage) GetRaceResultDriver(driver string) (*[]models.RaceResult, error) {

	return &[]models.RaceResult{}, nil
}

// AddRaceResult : Add a new race result
func (s *Storage) AddRaceResult(result models.RaceResult) error {

	_, err := s.raceResults.InsertOne(context.TODO(), result)
	if err != nil {
		return err
	}

	return nil

}
