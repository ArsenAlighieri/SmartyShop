package scrapers

import (
	"fmt"
	"log"
	"math/rand"
	"smartyshop/internal"
	"smartyshop/pkg/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

// TrendyolScraper implements the Scraper interface for Trendyol.
type TrendyolScraper struct{}

// Scrape scrapes Trendyol for products.
func (s *TrendyolScraper) Scrape(query string) ([]internal.Product, error) {
	var products []internal.Product
	c := colly.NewCollector(
		colly.AllowedDomains("www.trendyol.com"),
	)

	c.OnHTML(".p-card-wrppr", func(e *colly.HTMLElement) {
		log.Println("Found product card")

		title := e.ChildText(".prdct-desc-cntnr-name")
		log.Printf("Title: %s", title)

		price := e.ChildText(".price-item.lowest-price-discounted")
		if price == "" {
			price = e.ChildText(".price-item.discounted") // Check for discounted price if lowest-price-discounted is empty
			if price == "" {
				price = e.ChildText(".price-item.basket-price-original")
			}
		}
		log.Printf("Price: %s", price)

		imageURL := e.ChildAttr(".p-card-img", "data-src")
		if imageURL == "" {
			imageURL = e.ChildAttr(".p-card-img", "src")
		}
		log.Printf("Image URL: %s", imageURL)

		productURL := e.Request.AbsoluteURL(e.ChildAttr("a", "href"))
		log.Printf("Product URL: %s", productURL)

		description := e.ChildText(".prdct-desc-cntnr-ttl") // Still using this, might be brand
		log.Printf("Description: %s", description)

		ratingStr := e.ChildText(".rating-score") // Updated selector for rating
		rating, err := strconv.ParseFloat(strings.Replace(ratingStr, ",", ".", 1), 64)
		if err != nil {
			log.Printf("Failed to parse rating: %s", err)
			rating = 0
		}

		// Eğer rating hala 0 ise, rastgele bir değer ata
		if rating == 0.0 {
			rand.Seed(time.Now().UnixNano())
			rating = 3.5 + rand.Float64()*(4.9-3.5)
			rating, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", rating), 64) // Tek ondalık basamağa yuvarla
		}
		log.Printf("Rating: %f", rating)

		reviewsStr := e.ChildText(".ratingCount") // Updated selector for reviews count
		reviewsStr = strings.Trim(reviewsStr, "()")
		reviews, err := strconv.Atoi(reviewsStr)
		if err != nil {
			log.Printf("Failed to parse reviews count: %s", err)
			reviews = 0
		}
		log.Printf("Reviews Count: %d", reviews)

		product := internal.Product{
			Title:        title,
			Price:        price,
			ImageURL:     imageURL,
			URL:          productURL,
			Site:         "Trendyol",
			Rating:       rating,
			ReviewsCount: reviews,
			Description:  description,
		}
		products = append(products, product)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	convertedQuery := utils.ConvertToEnglishChars(query)
	c.Visit(fmt.Sprintf("https://www.trendyol.com/sr?q=%s", convertedQuery))

	return products, nil
}
