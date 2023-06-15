package scraper

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/mdoddzz/f1-scraper-go/pkg/models"
)

func (s *service) HandleData() {

	// Set collector HTML
	s.col.OnHTML(".resultsarchive-table > tbody", func(e *colly.HTMLElement) {

		// Get URL path
		path := e.Request.URL.Path

		// Get table type based on URL
		switch getUrlEnd(path) {

		case "races.html":

			e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
				tableData := models.Race{
					UrlID:     getIdFromURL(el.ChildAttr("td:nth-child(2) a", "href")),
					GrandPrix: el.ChildText("td:nth-child(2)"),
					Date: models.F1Time{
						String:   el.ChildText("td:nth-child(3)"),
						DateTime: handleF1Date(el.ChildText("td:nth-child(3)")),
					},
					Winner: handleF1Driver(el, "td:nth-child(4)"),
					Car:    el.ChildText("td:nth-child(5)"),
					Laps:   handleF1Int(el.ChildText("td:nth-child(6)")),
					Time: models.F1Time{
						String:   el.ChildText("td:nth-child(7)"),
						DateTime: handleF1Time(el.ChildText("td:nth-child(7)")),
					},
				}
				s.race.AddRace(tableData)
			})

		case "race-result.html":

			// Get race ID from URL
			race_id, err := s.getRaceId(path)
			if err != nil {
				fmt.Println("Unable to get raceID")
				break
			}

			e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
				tableData := models.RaceResult{
					RaceId:   race_id,
					Position: handleF1Int(el.ChildText("td:nth-child(2)")),
					Number:   handleF1Int(el.ChildText("td:nth-child(3)")),
					Driver:   handleF1Driver(el, "td:nth-child(4)"),
					Car:      el.ChildText("td:nth-child(5)"),
					Laps:     handleF1Int(el.ChildText("td:nth-child(6)")),
					Time:     el.ChildText("td:nth-child(7)"),
					Points:   handleF1Float(el.ChildText("td:nth-child(8)")),
				}
				s.race_results.AddRaceResult(tableData)
			})

		case "drivers.html":

		case "team.html":

		case "fastest-laps.html":

		default:

			if strings.Contains(path, "/drivers/") {

			}

			if strings.Contains(path, "/team/") {

			}

		}

		fmt.Println("Archive data:", e.DOM)
	})

}

func getUrlEnd(url string) string {

	return url[strings.LastIndex(url, "/")+1:]

}

func handleF1Date(str string) time.Time {

	pd, err := time.Parse("02 Jan 2006", str)
	if err != nil {
		fmt.Println("Error Parsing date")
		pd = time.Time{} // Set empty date
	}

	return pd

}

func handleF1Time(str string) time.Time {

	pt, err := time.Parse("15:04:05.999999999", str)
	if err != nil {
		fmt.Println("Error Parsing date")
		pt = time.Time{} // Set empty time
	}

	return pt

}

func handleF1Driver(el *colly.HTMLElement, querySelectorBase string) models.Driver {

	first_name := el.ChildText(querySelectorBase + ">.hide-for-tablet")
	last_name := el.ChildText(querySelectorBase + ">.hide-for-mobile")
	identifier := el.ChildText(querySelectorBase + ">.hide-for-desktop")

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
		fmt.Println("Error Parsing Int Data")
		i = 0 // Set empty ints
	}
	return i
}

func handleF1Float(str string) float64 {
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println("Error Parsing Int Data")
		f = 0 // Set empty ints
	}
	return f
}

func (s *service) getRaceId(url string) (string, error) {

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
