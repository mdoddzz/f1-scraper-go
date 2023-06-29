package mongo

import (
	"context"

	"github.com/mdoddzz/f1-scraper-go/pkg/models"
)

// GetConstructorStandings : Get constructor standings for a race
func (s *Storage) GetConstructorStandings(raceId interface{}) (*[]models.ConstructorStandings, error) {
	return &[]models.ConstructorStandings{}, nil
}

// AddConstructorStandings : Add a new constructor standings
func (s *Storage) AddConstructorStandings(standing models.ConstructorStandings) error {
	_, err := s.constructorStandings.InsertOne(context.TODO(), standing)
	if err != nil {
		return err
	}

	return nil

}
