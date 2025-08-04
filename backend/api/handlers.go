package api

import (
	"fmt"
	"smartyshop/config"
	"smartyshop/gemini"
	"smartyshop/internal"
	"smartyshop/scrapers"
	"sort"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// cacheEntry stores products and an expiration time.
type cacheEntry struct {
	products   []internal.Product
	expiration time.Time
}

// Handler holds the cache and other dependencies.
type Handler struct {
	Cache      map[string]cacheEntry
	CacheMutex *sync.Mutex
}

// NewHandler creates a new handler with an initialized cache.
func NewHandler() *Handler {
	return &Handler{
		Cache:      make(map[string]cacheEntry),
		CacheMutex: &sync.Mutex{},
	}
}

// GetProducts handles the /products endpoint.
func (h *Handler) GetProducts(c *gin.Context) {
	site := c.Query("site")
	query := c.Query("query")

	if site == "" || query == "" {
		c.JSON(400, gin.H{"error": "'site' and 'query' parameters are required"})
		return
	}

	cacheKey := fmt.Sprintf("%s-%s", site, query)

	h.CacheMutex.Lock()
	entry, found := h.Cache[cacheKey]
	h.CacheMutex.Unlock()

	if found && time.Now().Before(entry.expiration) {
		c.JSON(200, entry.products)
		return
	}

	var scraper scrapers.Scraper

	switch site {
	case "trendyol":
		scraper = &scrapers.TrendyolScraper{}
	case "teknosa":
		scraper = &scrapers.TeknosaScraper{}
	case "mediamarkt":
		scraper = &scrapers.MediaMarktScraper{}
	case "amazon":
		scraper = &scrapers.AmazonScraper{}
	default:
		c.JSON(400, gin.H{"error": "invalid site"})
		return
	}

	products, err := scraper.Scrape(query)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.CacheMutex.Lock()
	h.Cache[cacheKey] = cacheEntry{
		products:   products,
		expiration: time.Now().Add(10 * time.Second),
	}
	h.CacheMutex.Unlock()

	c.JSON(200, products)
}

// GetTop10Products handles the /products/top10 endpoint.
func (h *Handler) GetTop10Products(c *gin.Context) {
	site := c.Query("site")
	query := c.Query("query")

	if site == "" || query == "" {
		c.JSON(400, gin.H{"error": "'site' and 'query' parameters are required"})
		return
	}

	cacheKey := fmt.Sprintf("%s-%s", site, query)

	h.CacheMutex.Lock()
	entry, found := h.Cache[cacheKey]
	h.CacheMutex.Unlock()

	if !found || time.Now().After(entry.expiration) {
		c.JSON(400, gin.H{"error": "products not cached, please fetch them first with /products"})
		return
	}

	sortedProducts := entry.products
	sort.Slice(sortedProducts, func(i, j int) bool {
		return sortedProducts[i].Rating > sortedProducts[j].Rating
	})

	if len(sortedProducts) > 10 {
		sortedProducts = sortedProducts[:10]
	}

	c.JSON(200, sortedProducts)
}

// GeminiQuery handles the /gemini/query endpoint.
func (h *Handler) GeminiQuery(c *gin.Context) {
	type GeminiQueryRequest struct {
		Query    string             `json:"query"`
		Products []internal.Product `json:"products"`
	}

	var req GeminiQueryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	apiKey := config.GetGeminiAPIKey()
	if apiKey == "" {
		c.JSON(500, gin.H{"error": "GEMINI_API_KEY not set"})
		return
	}

	resp, err := gemini.GetGeminiProductInsights(req.Products, req.Query, apiKey)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resp)
}
