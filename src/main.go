package main

import (
	"log"
	"os"

	
	"tadrees.ml/reviews_api/src/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadEnvVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Unable to load env vars. Error:" + err.Error())
	}
}

func main() {

	// initialization steps
	loadEnvVariables()
	//config.GetMongoClient()

	PORT := os.Getenv("PORT")

	router := gin.Default()

	router.GET("/", handlers.RootContextEndpointHandler)
	router.GET("/health", handlers.HealthCheckEndpointHandler)

	reviewsGroup := router.Group("/reviews")

	reviewsGroup.GET("/", handlers.GetReviewsEndpointHandler)

	err := router.Run(":" + PORT)

	if err != nil {
		log.Fatalln("Failed to start the server. Error:" + err.Error())
	}
}
