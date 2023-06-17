package mongo

import (
	"context"

	"github.com/mdoddzz/f1-scraper-go/pkg/models"
)

// Get constructor standings by Year
func (s *Storage) GetConstructorStandingsSeason(year int) (*[]models.ConstructorStandingsSeason, error) {

	return &[]models.ConstructorStandingsSeason{}, nil

}

// Add a new constructor standing
func (s *Storage) AddConstructorStandingsSeason(result models.ConstructorStandingsSeason) error {

	_, err := s.constructor_standings_season.InsertOne(context.TODO(), result)
	if err != nil {
		return err
	}

	return nil

}
