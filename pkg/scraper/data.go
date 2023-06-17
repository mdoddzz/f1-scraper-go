package scraper

import (
	"fmt"
	"strings"

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

			e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
				tableData := models.DriverStandingsSeason{
					Year:        handleF1Int(getUrlPathByIndex(path, 3)),
					Position:    handleF1Int(el.ChildText("td:nth-child(2)")),
					Driver:      handleF1Driver(el, "td:nth-child(3)"),
					Nationality: el.ChildText("td:nth-child(4)"),
					Car:         el.ChildText("td:nth-child(5)"),
					Points:      handleF1Float(el.ChildText("td:nth-child(6)")),
				}
				s.driver_standings_season.AddDriverStandingSeason(tableData)
			})

		case "team.html":

			e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
				tableData := models.ConstructorStandingsSeason{
					Year:     handleF1Int(getUrlPathByIndex(path, 3)),
					Position: handleF1Int(el.ChildText("td:nth-child(2)")),
					Team:     el.ChildText("td:nth-child(3)"),
					Points:   handleF1Float(el.ChildText("td:nth-child(7)")),
				}
				s.constructor_standings_season.AddConstructorStandingsSeason(tableData)
			})

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
