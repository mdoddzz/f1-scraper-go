package mongo

import (
	"context"

	"github.com/mdoddzz/f1-scraper-go/pkg/models"
)

// GetPracticeResult : Get practice result
func (s *Storage) GetPracticeResult(raceId interface{}, session string) (*[]models.Practice, error) {

	return &[]models.Practice{}, nil

}

// AddPractice : Add a new practice result
func (s *Storage) AddPractice(practice models.Practice) error {

	_, err := s.practices.InsertOne(context.TODO(), practice)
	if err != nil {
		return err
	}

	return nil

}
