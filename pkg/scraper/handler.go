package scraper

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/mdoddzz/f1-scraper-go/pkg/models"
	"github.com/mdoddzz/f1-scraper-go/pkg/storage/mongo"
)

type service struct {
	col          *colly.Collector
	race         models.RacesService
	race_results models.RaceResultService
	baseURL      string
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
	return &service{newCollector(), storage, storage, "https://www.formula1.com"}
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
