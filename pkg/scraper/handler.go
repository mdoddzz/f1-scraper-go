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

type Service struct {
	col       *colly.Collector
	baseURL   string
	f1Service models.F1Services
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

func NewWithMongo(storage *mongo.Storage) *Service {
	return &Service{newCollector(), "https://www.formula1.com", storage}
}

/*func NewWithMySQL(storage *sql.DB) *service {
	return &service{newCollector(), storage}
}*/

func (s *Service) Start() {

	// Build all scraper functions
	s.ErrorHandler()
	s.HandleData()

	// Start getting URLs and get the data from those URLS
	s.CollectURLs("/en/results.html/2023/races.html")

}

func (s *Service) ErrorHandler() {

	s.col.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

}

func getUrlEnd(url string) string {

	return url[strings.LastIndex(url, "/")+1:]

}

func getUrlPathByIndex(url string, index int) string {

	splitUrl := strings.Split(url, "/")

	return splitUrl[index]

}

func handleF1Time(str string, dateOrTime string) *models.F1Time {

	t := models.F1Time{
		String: str,
	}

	switch dateOrTime {

	case "date":

		t.DateTime = handleF1DateFormat(str)

	case "time":

		t.DateTime = handleF1TimeFormat(str)

	}

	return &t

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

	firstName := el.ChildText(querySelectorBase + " .hide-for-tablet")
	lastName := el.ChildText(querySelectorBase + " .hide-for-mobile")
	identifier := el.ChildText(querySelectorBase + " .hide-for-desktop")

	return models.Driver{
		FullName:       firstName + " " + lastName,
		FirstName:      firstName,
		LastName:       lastName,
		NameIdentifier: identifier,
	}

}

func handleF1Int(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error Parsing Int Data: ", str)
		i = 0 // Set empty int
	}
	return i
}

func handleF1Float(str string) float64 {
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println("Error Parsing Float Data: ", str)
		f = 0 // Set empty float
	}
	return f
}

func handleF1IntOrString(str string) interface{} {
	var i interface{}
	i, err := strconv.Atoi(str)
	if err != nil {
		i = str
	}
	return i
}

func (s *Service) getRaceId(url string) (interface{}, error) {

	raceId, err := s.f1Service.GetRaceByUrlId(getIdFromURL(url))
	if err != nil {
		return "", err
	}

	return raceId.ID, nil

}

func getIdFromURL(url string) int {

	split := strings.Split(url, "/")

	return handleF1Int(split[5])

}

// GetStringInBetween Returns empty string if no start string found
func GetStringInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str[s:], end)
	if e == -1 {
		return
	}
	e += s + e - 1
	return str[s:e]
}
