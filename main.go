package main

import (
	"os"

	"candidate-avalanche-api/db"
	"candidate-avalanche-api/services"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	db.Connect()

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", services.CandidateGet)
	router.GET("/add", services.CandidateCreate)

	router.Run(":" + port)
}
