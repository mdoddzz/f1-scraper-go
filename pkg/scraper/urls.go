package scraper

import (
	"fmt"

	"github.com/gocolly/colly"
	"golang.org/x/exp/slices"
)

func (s *service) CollectURLs(startURL string) {

	// Set url array
	visited_urls := []string{startURL}

	// Get the URLS from links at top
	s.col.OnHTML(".resultsarchive-filter-item-link", func(e *colly.HTMLElement) {

		// Get url from archive link
		url := e.Attr("href")

		// Check if URL is new
		if !slices.Contains(visited_urls, url) {
			e.Request.Visit(s.baseURL + url)
			visited_urls = append(visited_urls, url)
		}

	})

	// Get URLs from sidebar
	s.col.OnHTML(".side-nav-item-link", func(e *colly.HTMLElement) {

		// Get url from archive link
		url := e.Attr("href")

		// Check if URL is new
		if !slices.Contains(visited_urls, url) {
			e.Request.Visit(s.baseURL + url)
			visited_urls = append(visited_urls, url)
		}

	})

	// On finish save to logs urls
	s.col.OnScraped(func(r *colly.Response) {
		// Only callback on initial URL
		if r.Request.URL.String() == s.baseURL+startURL {
			fmt.Println("Completed full scrape")
			//fmt.Println("Visited URLs: ", visited_urls)
		}
	})

	// Visit first URL
	s.col.Visit(s.baseURL + startURL)

}
