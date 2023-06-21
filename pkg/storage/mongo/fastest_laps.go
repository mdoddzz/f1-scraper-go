package mongo

import (
	"context"

	"github.com/mdoddzz/f1-scraper-go/pkg/models"
)

// Get fastest laps for a race
func (s *Storage) GetFastestLaps(raceId interface{}) (*[]models.FastestLaps, error) {
	return &[]models.FastestLaps{}, nil
}

// Add a new fastest lap
func (s *Storage) AddFastestLaps(fastestLaps models.FastestLaps) error {

	_, err := s.driver_standings.InsertOne(context.TODO(), fastestLaps)
	if err != nil {
		return err
	}

	return nil

}
