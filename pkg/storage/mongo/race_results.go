package mongo

import (
	"context"

	"github.com/mdoddzz/f1-scraper-go/pkg/models"
)

// Get race result
func (s *Storage) GetRaceResult(raceId string) (*[]models.RaceResult, error) {

	return &[]models.RaceResult{}, nil
}

// Get race results of a driver
func (s *Storage) GetRaceResultDriver(driver string) (*[]models.RaceResult, error) {

	return &[]models.RaceResult{}, nil
}

// Add a new race result
func (s *Storage) AddRaceResult(result models.RaceResult) error {

	_, err := s.race_results.InsertOne(context.TODO(), result)
	if err != nil {
		return err
	}

	return nil

}
