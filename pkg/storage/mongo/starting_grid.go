package mongo

import (
	"context"

	"github.com/mdoddzz/f1-scraper-go/pkg/models"
)

// Get starting grid of a race
func (s *Storage) GetStartingGrid(raceId interface{}) (*[]models.StartingGrid, error) {

	return &[]models.StartingGrid{}, nil

}

// Add a new starting grid
func (s *Storage) AddStartingGrid(startingGrid models.StartingGrid) error {

	_, err := s.pit_stops.InsertOne(context.TODO(), startingGrid)
	if err != nil {
		return err
	}

	return nil

}
