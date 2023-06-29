package mongo

import (
	"context"

	"github.com/mdoddzz/f1-scraper-go/pkg/models"
)

// GetFastestLapAward : Get practice result
func (s *Storage) GetFastestLapAward(raceId interface{}) (*models.FastestLapAward, error) {
	return &models.FastestLapAward{}, nil
}

// AddFastestLapAward : Add a new practice result
func (s *Storage) AddFastestLapAward(fastestLapAward models.FastestLapAward) error {

	_, err := s.driverStandings.InsertOne(context.TODO(), fastestLapAward)
	if err != nil {
		return err
	}

	return nil

}
