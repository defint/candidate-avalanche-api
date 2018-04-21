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

	router.GET("/candidate", services.CandidateList)
	router.PUT("/candidate", services.CandidateCreate)
	router.GET("/candidate/:id", services.CandidateItem)
	router.POST("/candidate/:id", services.CandidateUpdate)
	router.DELETE("/candidate/:id", services.CandidateDelete)

	router.Run(":" + port)
}
