package mongo

import (
	"context"

	"github.com/mdoddzz/f1-scraper-go/pkg/models"
)

// Get driver standings for a race
func (s *Storage) GetDriverStandings(raceId interface{}) (*[]models.DriverStandings, error) {
	return &[]models.DriverStandings{}, nil
}

// Add a new driver standings
func (s *Storage) AddDriverStandings(standing models.DriverStandings) error {

	_, err := s.driver_standings.InsertOne(context.TODO(), standing)
	if err != nil {
		return err
	}

	return nil

}
