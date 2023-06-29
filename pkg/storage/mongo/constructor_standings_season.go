package mongo

import (
	"context"

	"github.com/mdoddzz/f1-scraper-go/pkg/models"
)

// GetConstructorStandingsSeason : Get constructor standings by Year
func (s *Storage) GetConstructorStandingsSeason(year int) (*[]models.ConstructorStandingsSeason, error) {

	return &[]models.ConstructorStandingsSeason{}, nil

}

// AddConstructorStandingsSeason : Add a new constructor standing
func (s *Storage) AddConstructorStandingsSeason(result models.ConstructorStandingsSeason) error {

	_, err := s.constructorStandingsSeason.InsertOne(context.TODO(), result)
	if err != nil {
		return err
	}

	return nil

}
