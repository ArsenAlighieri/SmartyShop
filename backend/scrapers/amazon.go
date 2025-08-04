package scrapers

import (
	"fmt"
	"smartyshop/internal"
	"smartyshop/pkg/utils"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

type AmazonScraper struct{}

func (s *AmazonScraper) Scrape(query string) ([]internal.Product, error) {
	var products []internal.Product

	c := colly.NewCollector(
		colly.AllowedDomains("www.amazon.com.tr"),
	)

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64)" // Bot algısını azaltır

	c.OnHTML("div[data-component-type=s-search-result]", func(e *colly.HTMLElement) {
		// Başlık (title)
		title := e.ChildAttr("h2", "aria-label")
		if title == "" {
			title = e.ChildText("h2 span")
		}

		// Fiyat (tam fiyat, ₺ hariç)
		price := e.ChildText(".a-price .a-offscreen")
		price = strings.ReplaceAll(price, "TL", "")
		price = strings.TrimSpace(price)

		// Görsel URL
		imageURL := e.ChildAttr("img", "src")

		// Ürün linki
		url := e.Request.AbsoluteURL(e.ChildAttr("h2 a", "href"))

		// Puan (rating)
		rating := 0.0
		ratingRaw := e.ChildText("span.a-icon-alt") // En sağlam yer burası
		if ratingRaw != "" {
			parts := strings.Split(ratingRaw, " ")
			if len(parts) >= 4 {
				ratingStr := strings.Replace(parts[3], ",", ".", 1) // "4,4" → "4.4"
				r, err := strconv.ParseFloat(ratingStr, 64)
				if err == nil {
					rating = r
				}
			}
		}

		// Yorum sayısı (reviews count)
		reviews := 0
		e.DOM.Find("span.a-size-base").EachWithBreak(func(i int, s *goquery.Selection) bool {
			text := strings.TrimSpace(s.Text())
			if isDigitsOnly(text) {
				reviewsStr := strings.ReplaceAll(text, ".", "")
				r, err := strconv.Atoi(reviewsStr)
				if err == nil {
					reviews = r
					return false // bulduk, çık
				}
			}
			return true
		})

		product := internal.Product{
			Title:        title,
			Price:        price,
			ImageURL:     imageURL,
			URL:          url,
			Site:         "Amazon",
			Rating:       rating,
			ReviewsCount: reviews,
		}
		products = append(products, product)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	convertedQuery := utils.ConvertToEnglishChars(query)
	err := c.Visit(fmt.Sprintf("https://www.amazon.com.tr/s?k=%s", convertedQuery))
	if err != nil {
		return nil, err
	}

	return products, nil
}

// Yardımcı: Sadece rakamlardan oluşuyor mu?
func isDigitsOnly(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
