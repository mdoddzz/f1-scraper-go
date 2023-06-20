package scraper

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/mdoddzz/f1-scraper-go/pkg/models"
	"github.com/mdoddzz/f1-scraper-go/pkg/storage/mongo"
)

type service struct {
	col                          *colly.Collector
	race                         models.RacesService
	race_results                 models.RaceResultService
	driver_standings_season      models.DriverStandingSeasonService
	constructor_standings_season models.ConstructorStandingsSeasonService
	pit_stops                    models.PitStopService
	starting_grid                models.StartingGridService
	baseURL                      string
}

func newCollector() *colly.Collector {
	// Set new collector
	c := colly.NewCollector(
		// Visit only domains within the F1 website
		colly.AllowedDomains("formula1.com", "www.formula1.com", "https://formula1.com", "https://www.formula1.com"),
	)

	// Log on request start
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// Log on request end
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	// Log when finished all scraping on URL
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	return c
}

func NewWithMongo(storage *mongo.Storage) *service {
	return &service{newCollector(), storage, storage, storage, storage, storage, storage, "https://www.formula1.com"}
}

/*func NewWithMySQL(storage *sql.DB) *service {
	return &service{newCollector(), storage}
}*/

func (s *service) Start() {

	// Build all scraper functions
	s.ErrorHandler()
	s.HandleData()

	// Start getting URLs and get the data from those URLS
	s.CollectURLs("/en/results.html/2023/races.html")

}

func (s *service) ErrorHandler() {

	s.col.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

}

func getUrlEnd(url string) string {

	return url[strings.LastIndex(url, "/")+1:]

}

func getUrlPathByIndex(url string, index int) string {

	split_url := strings.Split(url, "/")

	return split_url[index]

}

func handleF1Time(str string, date_or_time string) models.F1Time {

	t := models.F1Time{
		String: str,
	}

	switch date_or_time {

	case "date":

		t.DateTime = handleF1DateFormat(str)

	case "time":

		t.DateTime = handleF1TimeFormat(str)

	}

	return t

}

func handleF1DateFormat(str string) time.Time {

	pd, err := time.Parse("02 Jan 2006", str)
	if err != nil {
		fmt.Println("Error Parsing date: ", str)
		pd = time.Time{} // Set empty date
	}

	return pd

}

func handleF1TimeFormat(str string) time.Time {

	// Check string length for format e.g. xx.xxx
	format := ""

	if len(str) == 6 {
		format = "05.999999999"
	} else {
		format = "15:04:05.999999999"
	}

	pt, err := time.Parse(format, str)
	if err != nil {
		fmt.Println("Error Parsing date: ", str)
		pt = time.Time{} // Set empty time
	}

	return pt

}

func handleF1Driver(el *colly.HTMLElement, querySelectorBase string) models.Driver {

	first_name := el.ChildText(querySelectorBase + " .hide-for-tablet")
	last_name := el.ChildText(querySelectorBase + " .hide-for-mobile")
	identifier := el.ChildText(querySelectorBase + " .hide-for-desktop")

	return models.Driver{
		FullName:       first_name + " " + last_name,
		FirstName:      first_name,
		LastName:       last_name,
		NameIdentifier: identifier,
	}

}

func handleF1Int(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error Parsing Int Data: ", str)
		i = 0 // Set empty ints
	}
	return i
}

func handleF1Float(str string) float64 {
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println("Error Parsing Float Data: ", str)
		f = 0 // Set empty ints
	}
	return f
}

func (s *service) getRaceId(url string) (interface{}, error) {

	race_id, err := s.race.GetRaceByUrlId(getIdFromURL(url))
	if err != nil {
		return "", err
	}

	return race_id.ID, nil

}
func getIdFromURL(url string) int {

	split := strings.Split(url, "/")

	return handleF1Int(split[5])

}
