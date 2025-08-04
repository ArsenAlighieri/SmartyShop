package scrapers

import "smartyshop/internal"

// Scraper defines the interface for a scraper.
type Scraper interface {
	Scrape(query string) ([]internal.Product, error)
}
