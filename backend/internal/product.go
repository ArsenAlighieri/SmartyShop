package internal

// Product represents a product scraped from an e-commerce site.

type Product struct {
	Title        string  `json:"title"`
	Price        string  `json:"price"`
	Rating       float64 `json:"rating"`
	ReviewsCount int     `json:"reviews_count"`
	URL          string  `json:"url"`
	ImageURL     string  `json:"image_url"`
	Description  string  `json:"description"`
	Site         string  `json:"site"`
}
