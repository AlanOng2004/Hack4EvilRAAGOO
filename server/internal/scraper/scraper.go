package scraper

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

// Grant struct matches your Database Schema
type Grant struct {
	Title       string
	Agency      string
	Deadline    string
	Description string
	Quantum     string // e.g., "$20,000"
}

func ScrapeGrants() []Grant {
	var grants []Grant

	c := colly.NewCollector(
		colly.AllowedDomains("www.oursggrants.gov.sg", "oursggrants.gov.sg"),
	)

	// Find the HTML element that holds each grant card
	// [IMPORTANT]: You must inspect the website to get the real class name like ".grant-card"
	c.OnHTML(".grant-listing-item", func(e *colly.HTMLElement) {
		item := Grant{
			Title:    e.ChildText("h2.title"),
			Agency:   e.ChildText(".agency-name"),
			Deadline: e.ChildText(".closing-date"), // e.g., "Closing: 12 Jan 2026"
			// You might need to parse "Quantum" from the description text
			Description: e.ChildText(".description"),
		}
		grants = append(grants, item)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// Start scraping
	err := c.Visit("https://www.oursggrants.gov.sg/grants")
	if err != nil {
		log.Println("Error scraping:", err)
	}

	return grants
}
