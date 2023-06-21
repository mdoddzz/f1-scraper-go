package scraper

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/mdoddzz/f1-scraper-go/pkg/models"
)

func (s *service) HandleData() {

	// Create collector for URLs not handled
	uh := []string{}

	// Set collector HTML
	s.col.OnHTML(".resultsarchive-table > tbody", func(e *colly.HTMLElement) {

		// Get URL path
		path := e.Request.URL.Path

		// Get URL end
		path_end := getUrlEnd(path)

		// Get table type based on URL
		switch path_end {

		case "races.html":

			e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
				tableData := models.Race{
					UrlID:     getIdFromURL(el.ChildAttr("td:nth-child(2) a", "href")),
					GrandPrix: el.ChildText("td:nth-child(2)"),
					Date:      *handleF1Time(el.ChildText("td:nth-child(3)"), "date"),
					Winner:    handleF1Driver(el, "td:nth-child(4)"),
					Car:       el.ChildText("td:nth-child(5)"),
					Laps:      handleF1Int(el.ChildText("td:nth-child(6)")),
					Time:      *handleF1Time(el.ChildText("td:nth-child(7)"), "time"),
				}
				s.f1_service.AddRace(tableData)
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
					Position: handleF1IntOrString(el.ChildText("td:nth-child(2)")),
					Number:   handleF1Int(el.ChildText("td:nth-child(3)")),
					Driver:   handleF1Driver(el, "td:nth-child(4)"),
					Car:      el.ChildText("td:nth-child(5)"),
					Laps:     handleF1Int(el.ChildText("td:nth-child(6)")),
					Time:     el.ChildText("td:nth-child(7)"),
					Points:   handleF1Float(el.ChildText("td:nth-child(8)")),
				}
				s.f1_service.AddRaceResult(tableData)
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
				s.f1_service.AddDriverStandingSeason(tableData)
			})

		case "team.html":

			e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
				tableData := models.ConstructorStandingsSeason{
					Year:     handleF1Int(getUrlPathByIndex(path, 3)),
					Position: handleF1Int(el.ChildText("td:nth-child(2)")),
					Team:     el.ChildText("td:nth-child(3)"),
					Points:   handleF1Float(el.ChildText("td:nth-child(4)")),
				}
				s.f1_service.AddConstructorStandingsSeason(tableData)
			})

		case "fastest-laps.html":

			// check if it is for season or round
			if len(strings.Split(path, "/")) == 5 {

				// Fastest laps Season

			} else {

				// Fastest Laps Round

			}

		case "pit-stop-summary.html":

			// Get race ID from URL
			race_id, err := s.getRaceId(path)
			if err != nil {
				fmt.Println("Unable to get raceID")
				break
			}

			e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
				tableData := models.PitStop{
					RaceId:    race_id,
					Stops:     handleF1Int(el.ChildText("td:nth-child(2)")),
					Number:    handleF1Int(el.ChildText("td:nth-child(3)")),
					Driver:    handleF1Driver(el, "td:nth-child(4)"),
					Car:       el.ChildText("td:nth-child(5)"),
					Lap:       handleF1Int(el.ChildText("td:nth-child(6)")),
					TimeOfDay: *handleF1Time(el.ChildText("td:nth-child(7)"), "time"),
					Time:      *handleF1Time(el.ChildText("td:nth-child(8)"), "time"),
					Total:     *handleF1Time(el.ChildText("td:nth-child(9)"), "time"),
				}
				s.f1_service.AddPitStop(tableData)
			})

		case "starting-grid.html":

			// Get race ID from URL
			race_id, err := s.getRaceId(path)
			if err != nil {
				fmt.Println("Unable to get raceID")
				break
			}

			e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
				tableData := models.StartingGrid{
					RaceId:   race_id,
					Position: handleF1Int(el.ChildText("td:nth-child(2)")),
					Number:   handleF1Int(el.ChildText("td:nth-child(3)")),
					Driver:   handleF1Driver(el, "td:nth-child(4)"),
					Car:      el.ChildText("td:nth-child(5)"),
					Time:     *handleF1Time(el.ChildText("td:nth-child(6)"), "time"),
				}
				s.f1_service.AddStartingGrid(tableData)
			})

		case "qualifying.html":

			// Get race ID from URL
			race_id, err := s.getRaceId(path)
			if err != nil {
				fmt.Println("Unable to get raceID")
				break
			}

			e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
				tableData := models.Qualifying{
					RaceId:   race_id,
					Position: handleF1Int(el.ChildText("td:nth-child(2)")),
					Number:   handleF1Int(el.ChildText("td:nth-child(3)")),
					Driver:   handleF1Driver(el, "td:nth-child(4)"),
					Car:      el.ChildText("td:nth-child(5)"),
					Q1:       handleF1Time(el.ChildText("td:nth-child(6)"), "time"),
					Q2:       handleF1Time(el.ChildText("td:nth-child(7)"), "time"),
					Q3:       handleF1Time(el.ChildText("td:nth-child(8)"), "time"),
					Laps:     handleF1Int(el.ChildText("td:nth-child(9)")),
				}
				s.f1_service.AddQualifyingResult(tableData)
			})

		case "qualifying-0.html", "qualifying-1.html", "qualifying-2.html":

			// Get race ID from URL
			race_id, err := s.getRaceId(path)
			if err != nil {
				fmt.Println("Unable to get raceID")
				break
			}

			session := ""
			if path_end == "qualifying-0.html" {
				session = "Overall Qualifying"
			} else {
				session = "Qualifying " + GetStringInBetween(path_end, "-", ".")
			}

			c := 0
			e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
				c++
			})

			if c == 1 {
				session = "Pole Position"
			}

			e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
				tableData := models.Qualifying{
					RaceId:   race_id,
					Session:  session,
					Position: handleF1Int(el.ChildText("td:nth-child(2)")),
					Number:   handleF1Int(el.ChildText("td:nth-child(3)")),
					Driver:   handleF1Driver(el, "td:nth-child(4)"),
					Car:      el.ChildText("td:nth-child(5)"),
					Time:     handleF1Time(el.ChildText("td:nth-child(6)"), "time"),
					Laps:     handleF1Int(el.ChildText("td:nth-child(7)")),
				}
				s.f1_service.AddQualifyingResult(tableData)
			})

		case "sprint-grid.html":

		case "sprint-shootout.html":

		case "sprint-results.html":

		case "practice-0.html", "practice-1.html", "practice-2.html", "practice-3.html", "practice-4.html":

			// Get race ID from URL
			race_id, err := s.getRaceId(path)
			if err != nil {
				fmt.Println("Unable to get raceID")
				break
			}

			session := ""
			if path_end == "practice-0.html" {
				session = "Warm Up"
			} else {
				session = "Practice " + GetStringInBetween(path_end, "-", ".")
			}

			e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
				tableData := models.Practice{
					RaceId:   race_id,
					Session:  session,
					Position: handleF1Int(el.ChildText("td:nth-child(2)")),
					Number:   handleF1Int(el.ChildText("td:nth-child(3)")),
					Driver:   handleF1Driver(el, "td:nth-child(4)"),
					Car:      el.ChildText("td:nth-child(5)"),
					Time:     *handleF1Time(el.ChildText("td:nth-child(6)"), "time"),
					Gap:      el.ChildText("td:nth-child(7)"),
					Laps:     handleF1Int(el.ChildText("td:nth-child(8)")),
				}
				s.f1_service.AddPractice(tableData)
			})

		default:

			if strings.Contains(path, "/drivers/") {
				fmt.Println("Driver")
				break
			}

			if strings.Contains(path, "/team/") {
				fmt.Println("Team")
				break
			}

			uh = append(uh, path)

		}
	})

	// On finish save to logs urls
	s.col.OnScraped(func(r *colly.Response) {
		fmt.Println("URLs Not Handled: ", uh)
	})

}
