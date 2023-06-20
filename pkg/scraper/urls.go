package scraper

import (
	"fmt"

	"github.com/gocolly/colly"
	"golang.org/x/exp/slices"
)

func (s *service) CollectURLs(startURL string) {

	// Set url array
	archive_urls := []string{startURL}

	// Get the URLS from links at top
	s.col.OnHTML(".resultsarchive-filter-item-link", func(e *colly.HTMLElement) {

		// Get url from archive link
		url := e.Attr("href")

		// Check if URL is new
		if !slices.Contains(archive_urls, url) {
			fmt.Println("Found New Archive URL ", e.Attr("href"))
			s.col.Visit(s.baseURL + url)
			archive_urls = append(archive_urls, url)
		}

	})

	// Get URLs from sidebar
	s.col.OnHTML(".side-nav-item-link", func(e *colly.HTMLElement) {

		// Get url from archive link
		url := e.Attr("href")

		// Check if URL is new
		if !slices.Contains(archive_urls, url) {
			fmt.Println("Found New Archive URL ", e.Attr("href"))
			s.col.Visit(s.baseURL + url)
			archive_urls = append(archive_urls, url)
		}

	})

	// On finish save to logs urls

	// Visit first URL
	s.col.Visit(s.baseURL + startURL)

}
