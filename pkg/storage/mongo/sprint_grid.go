package mongo

import (
	"context"

	"github.com/mdoddzz/f1-scraper-go/pkg/models"
)

// GetSprintGrid : Get sprint grid of a race
func (s *Storage) GetSprintGrid(raceId interface{}) (*[]models.SprintGrid, error) {

	return &[]models.SprintGrid{}, nil

}

// AddSprintGrid : Add a new sprint grid
func (s *Storage) AddSprintGrid(sprintGrid models.SprintGrid) error {

	_, err := s.sprintGrid.InsertOne(context.TODO(), sprintGrid)
	if err != nil {
		return err
	}

	return nil

}
