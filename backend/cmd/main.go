package main

import (
	"log"
	"os"
	"smartyshop/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file. Assumes the app is run from the 'backend' directory.
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("FATAL: Error loading .env file. Please ensure it's in the 'backend' directory and saved with UTF-8 encoding (without BOM). Error: %v", err)
	}

	// Confirm that the API key is loaded.
	if os.Getenv("GEMINI_API_KEY") == "" {
		log.Fatalf("FATAL: GEMINI_API_KEY is not set in the .env file. The application cannot start.")
	}
	log.Println("INFO: .env file loaded and GEMINI_API_KEY found.")

	r := gin.Default()

	r.Use(cors.Default())

	h := api.NewHandler()

	r.GET("/products", h.GetProducts)
	r.GET("/products/top10", h.GetTop10Products)
	r.POST("/gemini/query", h.GeminiQuery)

	r.Run() // listen and serve on 0.0.0.0:8080
}
