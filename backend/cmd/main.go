package main

import (
	"log"
	"smartyshop/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using environment variables")
	}

	r := gin.Default()

	r.Use(cors.Default())

	h := api.NewHandler()

	r.GET("/products", h.GetProducts)
	r.GET("/products/top10", h.GetTop10Products)
	r.POST("/gemini/query", h.GeminiQuery)

	r.Run() // listen and serve on 0.0.0.0:8080
}
