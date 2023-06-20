package mongo

import (
	"context"

	"github.com/mdoddzz/f1-scraper-go/pkg/models"
)

// Get pit stop for a race
func (s *Storage) GetPitStops(raceId interface{}) (*[]models.PitStop, error) {

	// Check ID a primitive.ObjectId

	return &[]models.PitStop{}, nil

}

// Add a new pit stop
func (s *Storage) AddPitStop(pitStop models.PitStop) error {

	_, err := s.pit_stops.InsertOne(context.TODO(), pitStop)
	if err != nil {
		return err
	}

	return nil

}
