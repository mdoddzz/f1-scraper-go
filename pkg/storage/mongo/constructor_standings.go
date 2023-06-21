package mongo

import (
	"context"

	"github.com/mdoddzz/f1-scraper-go/pkg/models"
)

// Get constructor standings for a race
func (s *Storage) GetConstructorStandings(raceId interface{}) (*[]models.ConstructorStandings, error) {
	return &[]models.ConstructorStandings{}, nil
}

// Add a new constructor standings
func (s *Storage) AddConstructorStandings(standing models.ConstructorStandings) error {
	_, err := s.constructor_standings.InsertOne(context.TODO(), standing)
	if err != nil {
		return err
	}

	return nil

}
