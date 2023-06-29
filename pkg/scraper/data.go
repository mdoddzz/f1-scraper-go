package scraper

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/mdoddzz/f1-scraper-go/pkg/models"
)

func (s *Service) HandleData() {

	// Create collector for URLs not handled
	var uh []string

	// Set collector HTML
	s.col.OnHTML(".resultsarchive-table > tbody", func(e *colly.HTMLElement) {

		// Get URL path
		path := e.Request.URL.Path

		// Get URL end
		pathEnd := getUrlEnd(path)

		// Get table type based on URL
		switch pathEnd {

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
				err := s.f1Service.AddRace(tableData)
				if err != nil {
					fmt.Println("Unable to save race")
				}
			})

		case "race-result.html":

			// Get race ID from URL
			raceId, err := s.getRaceId(path)
			if err != nil {
				fmt.Println("Unable to get raceID")
				break
			}

			e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
				tableData := models.RaceResult{
					RaceId:   raceId,
					Position: handleF1IntOrString(el.ChildText("td:nth-child(2)")),
					Number:   handleF1Int(el.ChildText("td:nth-child(3)")),
					Driver:   handleF1Driver(el, "td:nth-child(4)"),
					Car:      el.ChildText("td:nth-child(5)"),
					Laps:     handleF1Int(el.ChildText("td:nth-child(6)")),
					Time:     el.ChildText("td:nth-child(7)"),
					Points:   handleF1Float(el.ChildText("td:nth-child(8)")),
				}
				err := s.f1Service.AddRaceResult(tableData)
				if err != nil {
					fmt.Println("Unable to save race result")
				}
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
				err := s.f1Service.AddDriverStandingSeason(tableData)
				if err != nil {
					fmt.Println("Unable to save driver standings season")
				}
			})

		case "team.html":

			e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
				tableData := models.ConstructorStandingsSeason{
					Year:     handleF1Int(getUrlPathByIndex(path, 3)),
					Position: handleF1Int(el.ChildText("td:nth-child(2)")),
					Team:     el.ChildText("td:nth-child(3)"),
					Points:   handleF1Float(el.ChildText("td:nth-child(4)")),
				}
				err := s.f1Service.AddConstructorStandingsSeason(tableData)
				if err != nil {
					fmt.Println("Unable to save constructor standings season")
				}
			})

		case "fastest-laps.html":

			// check if it is for season or round
			if len(strings.Split(path, "/")) == 5 {

				// Get Race ID from GP and Year
				raceId := ""

				// Fastest laps Season / Awards
				e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
					tableData := models.FastestLapAward{
						RaceId: raceId,
						Driver: handleF1Driver(el, "td:nth-child(3)"),
						Car:    el.ChildText("td:nth-child(4)"),
						Time:   *handleF1Time(el.ChildText("td:nth-child(5)"), "time"),
					}
					err := s.f1Service.AddFastestLapAward(tableData)
					if err != nil {
						fmt.Println("Unable to save fastest laps award")
					}
				})

			} else {

				// Get Race ID from URL
				raceId, err := s.getRaceId(path)
				if err != nil {
					fmt.Println("Unable to get raceID")
					break
				}

				// Fastest Laps Round
				e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
					tableData := models.FastestLaps{
						RaceId:    raceId,
						Position:  handleF1Int(el.ChildText("td:nth-child(2)")),
						Number:    handleF1Int(el.ChildText("td:nth-child(3)")),
						Driver:    handleF1Driver(el, "td:nth-child(4)"),
						Lap:       handleF1Int(el.ChildText("td:nth-child(5)")),
						TimeOfDay: *handleF1Time(el.ChildText("td:nth-child(6)"), "time"),
						Time:      *handleF1Time(el.ChildText("td:nth-child(7)"), "time"),
						AvgSpeed:  handleF1Float(el.ChildText("td:nth-child(8)")),
					}
					err := s.f1Service.AddFastestLaps(tableData)
					if err != nil {
						fmt.Println("Unable to save fastest lap")
					}
				})

			}

		case "pit-stop-summary.html":

			// Get race ID from URL
			raceId, err := s.getRaceId(path)
			if err != nil {
				fmt.Println("Unable to get raceID")
				break
			}

			e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
				tableData := models.PitStop{
					RaceId:    raceId,
					Stops:     handleF1Int(el.ChildText("td:nth-child(2)")),
					Number:    handleF1Int(el.ChildText("td:nth-child(3)")),
					Driver:    handleF1Driver(el, "td:nth-child(4)"),
					Car:       el.ChildText("td:nth-child(5)"),
					Lap:       handleF1Int(el.ChildText("td:nth-child(6)")),
					TimeOfDay: *handleF1Time(el.ChildText("td:nth-child(7)"), "time"),
					Time:      *handleF1Time(el.ChildText("td:nth-child(8)"), "time"),
					Total:     *handleF1Time(el.ChildText("td:nth-child(9)"), "time"),
				}
				err := s.f1Service.AddPitStop(tableData)
				if err != nil {
					fmt.Println("Unable to save pit stop")
				}
			})

		case "starting-grid.html":

			// Get race ID from URL
			raceId, err := s.getRaceId(path)
			if err != nil {
				fmt.Println("Unable to get raceID")
				break
			}

			e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
				tableData := models.StartingGrid{
					RaceId:   raceId,
					Position: handleF1Int(el.ChildText("td:nth-child(2)")),
					Number:   handleF1Int(el.ChildText("td:nth-child(3)")),
					Driver:   handleF1Driver(el, "td:nth-child(4)"),
					Car:      el.ChildText("td:nth-child(5)"),
					Time:     *handleF1Time(el.ChildText("td:nth-child(6)"), "time"),
				}
				err := s.f1Service.AddStartingGrid(tableData)
				if err != nil {
					fmt.Println("Unable to save starting grid")
				}
			})

		case "qualifying.html":

			// Get race ID from URL
			raceId, err := s.getRaceId(path)
			if err != nil {
				fmt.Println("Unable to get raceID")
				break
			}

			e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
				tableData := models.Qualifying{
					RaceId:   raceId,
					Position: handleF1Int(el.ChildText("td:nth-child(2)")),
					Number:   handleF1Int(el.ChildText("td:nth-child(3)")),
					Driver:   handleF1Driver(el, "td:nth-child(4)"),
					Car:      el.ChildText("td:nth-child(5)"),
					Q1:       handleF1Time(el.ChildText("td:nth-child(6)"), "time"),
					Q2:       handleF1Time(el.ChildText("td:nth-child(7)"), "time"),
					Q3:       handleF1Time(el.ChildText("td:nth-child(8)"), "time"),
					Laps:     handleF1Int(el.ChildText("td:nth-child(9)")),
				}
				err := s.f1Service.AddQualifyingResult(tableData)
				if err != nil {
					fmt.Println("Unable to save qualifying")
				}
			})

		case "qualifying-0.html", "qualifying-1.html", "qualifying-2.html":

			// Get race ID from URL
			raceId, err := s.getRaceId(path)
			if err != nil {
				fmt.Println("Unable to get raceID")
				break
			}

			session := ""
			if pathEnd == "qualifying-0.html" {
				session = "Overall Qualifying"
			} else {
				session = "Qualifying " + GetStringInBetween(pathEnd, "-", ".")
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
					RaceId:   raceId,
					Session:  session,
					Position: handleF1Int(el.ChildText("td:nth-child(2)")),
					Number:   handleF1Int(el.ChildText("td:nth-child(3)")),
					Driver:   handleF1Driver(el, "td:nth-child(4)"),
					Car:      el.ChildText("td:nth-child(5)"),
					Time:     handleF1Time(el.ChildText("td:nth-child(6)"), "time"),
					Laps:     handleF1Int(el.ChildText("td:nth-child(7)")),
				}
				err := s.f1Service.AddQualifyingResult(tableData)
				if err != nil {
					fmt.Println("Unable to save qualifying")
				}
			})

		case "sprint-grid.html":

		case "sprint-shootout.html":

		case "sprint-results.html":

		case "practice-0.html", "practice-1.html", "practice-2.html", "practice-3.html", "practice-4.html":

			// Get race ID from URL
			raceId, err := s.getRaceId(path)
			if err != nil {
				fmt.Println("Unable to get raceID")
				break
			}

			session := ""
			if pathEnd == "practice-0.html" {
				session = "Warm Up"
			} else {
				session = "Practice " + GetStringInBetween(pathEnd, "-", ".")
			}

			e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
				tableData := models.Practice{
					RaceId:   raceId,
					Session:  session,
					Position: handleF1Int(el.ChildText("td:nth-child(2)")),
					Number:   handleF1Int(el.ChildText("td:nth-child(3)")),
					Driver:   handleF1Driver(el, "td:nth-child(4)"),
					Car:      el.ChildText("td:nth-child(5)"),
					Time:     *handleF1Time(el.ChildText("td:nth-child(6)"), "time"),
					Gap:      el.ChildText("td:nth-child(7)"),
					Laps:     handleF1Int(el.ChildText("td:nth-child(8)")),
				}
				err := s.f1Service.AddPractice(tableData)
				if err != nil {
					fmt.Println("Unable to save practice")
				}
			})

		default:

			if strings.Contains(path, "/drivers/") {

				e.ForEach("tr", func(_ int, el *colly.HTMLElement) {

					// Get Race ID from GP and Date
					gpDate := handleF1Time(el.ChildAttr("td:nth-child(2)", "date"), "date").DateTime
					raceId, err := s.f1Service.GetRaceByGPDate(el.ChildText("td:nth-child(1)"), gpDate)
					if err == nil {

						tableData := models.DriverStandings{
							RaceId:       raceId,
							Driver:       models.Driver{},
							Car:          el.ChildText("td:nth-child(4)"),
							RacePosition: handleF1Int(el.ChildText("td:nth-child(5)")),
							Points:       handleF1Float(el.ChildText("td:nth-child(6)")),
						}
						err := s.f1Service.AddDriverStandings(tableData)
						if err != nil {
							fmt.Println("Unable to save driver standing")
						}

					} else {
						fmt.Println("Unable to get raceID from GP and Date")
					}

				})
				break
			}

			if strings.Contains(path, "/team/") {

				// Get Race ID from GP and Year

				e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
					tableData := models.ConstructorStandings{}
					err := s.f1Service.AddConstructorStandings(tableData)
					if err != nil {
						fmt.Println("Unable to save contractor standings")
					}
				})
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
