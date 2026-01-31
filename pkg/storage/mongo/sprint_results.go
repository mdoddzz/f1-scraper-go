package mongo

import (
	"context"

	"github.com/mdoddzz/f1-scraper-go/pkg/models"
)

// GetSprintResult : Get sprint result
func (s *Storage) GetSprintResult(raceId interface{}) (*[]models.SprintResult, error) {

	return &[]models.SprintResult{}, nil

}

// GetSprintResultDriver : Get sprint results of a driver
func (s *Storage) GetSprintResultDriver(driver string) (*[]models.SprintResult, error) {

	return &[]models.SprintResult{}, nil

}

// AddSprintResult : Add a new sprint result
func (s *Storage) AddSprintResult(sprintResult models.SprintResult) error {

	_, err := s.sprintResults.InsertOne(context.TODO(), sprintResult)
	if err != nil {
		return err
	}

	return nil

}
