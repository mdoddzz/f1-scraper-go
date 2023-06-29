package mongo

import (
	"context"

	"github.com/mdoddzz/f1-scraper-go/pkg/models"
)

// GetDriverStandingsByYear : Get driver standings by Year
func (s *Storage) GetDriverStandingsByYear(year int) (*[]models.DriverStandingsSeason, error) {

	return &[]models.DriverStandingsSeason{}, nil

}

// AddDriverStandingSeason : Add a new driver standing
func (s *Storage) AddDriverStandingSeason(result models.DriverStandingsSeason) error {

	_, err := s.driverStandingsSeason.InsertOne(context.TODO(), result)
	if err != nil {
		return err
	}

	return nil

}
