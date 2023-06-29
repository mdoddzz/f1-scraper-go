package mongo

import (
	"context"

	"github.com/mdoddzz/f1-scraper-go/pkg/models"
)

// GetDriverStandings : Get driver standings for a race
func (s *Storage) GetDriverStandings(raceId interface{}) (*[]models.DriverStandings, error) {
	return &[]models.DriverStandings{}, nil
}

// AddDriverStandings : Add a new driver standings
func (s *Storage) AddDriverStandings(standing models.DriverStandings) error {

	_, err := s.driverStandings.InsertOne(context.TODO(), standing)
	if err != nil {
		return err
	}

	return nil

}
