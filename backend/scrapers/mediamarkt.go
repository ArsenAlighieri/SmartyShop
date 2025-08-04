package scrapers

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"smartyshop/internal"
)

type MediaMarktScraper struct{}

func (s *MediaMarktScraper) Scrape(query string) ([]internal.Product, error) {
	var products []internal.Product

	c := colly.NewCollector(
		colly.AllowedDomains("www.mediamarkt.com.tr"),
	)

	c.OnHTML("div.sc-43f40bb6-0.QXAyC", func(e *colly.HTMLElement) {
		// Ürün linki
		url := e.ChildAttr("a[data-test='mms-router-link-product-list-item-link']", "href")
		if url == "" {
			log.Println("Ürün linki boş, atlanıyor")
			return
		}
		fullURL := "https://www.mediamarkt.com.tr" + url

		// Başlık
		title := e.ChildText("p[data-test='product-title']")
		if title == "" {
			log.Println("Başlık boş, atlanıyor")
			return
		}

		// Fiyat
		price := e.ChildText("div[data-test='mms-price'] span.sc-5a9f6c31-0.cmEuny")
		price = strings.TrimSpace(price)

		// Rating (puan)
		ratingText := e.ChildAttr("div[data-test='mms-customer-rating']", "aria-label") // "Ortalama ürün değerlendirmesi: 5 yıldız üzerinden 3.5"
		rating := 0.0
		re := regexp.MustCompile(`(\d+(\.\d+)?)`) // Ondalıklı sayı yakalamak için
		matches := re.FindStringSubmatch(ratingText)
		if len(matches) > 0 {
			r, err := strconv.ParseFloat(matches[1], 64)
			if err == nil {
				rating = r
			}
		}

		// Yorum sayısı
		reviewCountStr := e.ChildText("span[data-test='mms-customer-rating-count']")
		reviewCountStr = strings.TrimSpace(reviewCountStr)
		reviewCount := 0
		if reviewCountStr != "" {
			rc, err := strconv.Atoi(reviewCountStr)
			if err == nil {
				reviewCount = rc
			}
		}

		// Resim URL'si
		imageURL := e.ChildAttr("picture img", "src")

		product := internal.Product{
			Title:        title,
			Price:        price,
			URL:          fullURL,
			ImageURL:     imageURL,
			Site:         "MediaMarkt",
			Rating:       rating,
			ReviewsCount: reviewCount,
		}

		products = append(products, product)
	})

	searchURL := fmt.Sprintf("https://www.mediamarkt.com.tr/tr/search.html?query=%s", query)
	err := c.Visit(searchURL)
	if err != nil {
		return nil, err
	}

	return products, nil
}
