package mongo

import (
	"context"

	"github.com/mdoddzz/f1-scraper-go/pkg/models"
)

// GetQualifyingResult : Get qualifying result by race ID
func (s *Storage) GetQualifyingResult(raceId interface{}) (*[]models.Qualifying, error) {
	return &[]models.Qualifying{}, nil
}

// AddQualifyingResult : Add a new qualifying result
func (s *Storage) AddQualifyingResult(qualifying models.Qualifying) error {

	_, err := s.qualifying.InsertOne(context.TODO(), qualifying)
	if err != nil {
		return err
	}

	return nil

}
