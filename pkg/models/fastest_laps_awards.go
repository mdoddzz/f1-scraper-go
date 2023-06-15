package models

import "time"

// FastestLapsAwards : model for the fastest lap for each race
type FastestLapsAwards struct {
	RaceId string
	Driver Driver
	Car    string
	Time   time.Time
}
