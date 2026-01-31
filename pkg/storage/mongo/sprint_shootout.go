package mongo

import (
	"context"

	"github.com/mdoddzz/f1-scraper-go/pkg/models"
)

// GetSprintShootout : Get sprint shootout result by race ID
func (s *Storage) GetSprintShootout(raceId interface{}) (*[]models.SprintShootout, error) {

	return &[]models.SprintShootout{}, nil

}

// AddSprintShootout : Add a new sprint shootout result
func (s *Storage) AddSprintShootout(sprintShootout models.SprintShootout) error {

	_, err := s.sprintShootout.InsertOne(context.TODO(), sprintShootout)
	if err != nil {
		return err
	}

	return nil

}
