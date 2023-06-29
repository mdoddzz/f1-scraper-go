package mongo

import (
	"context"

	"github.com/mdoddzz/f1-scraper-go/pkg/models"
)

// GetStartingGrid : Get starting grid of a race
func (s *Storage) GetStartingGrid(raceId interface{}) (*[]models.StartingGrid, error) {

	return &[]models.StartingGrid{}, nil

}

// AddStartingGrid : Add a new starting grid
func (s *Storage) AddStartingGrid(startingGrid models.StartingGrid) error {

	_, err := s.pitStops.InsertOne(context.TODO(), startingGrid)
	if err != nil {
		return err
	}

	return nil

}
