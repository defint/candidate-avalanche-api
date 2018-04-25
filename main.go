package main

import (
	"os"

	"candidate-avalanche-api/db"
	_ "candidate-avalanche-api/docs"
	"candidate-avalanche-api/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"time"
)

// @title Avalanche Candidates API
// @version 1.0
// @description Avalanche Candidates API

// @contact.name Anton Matrenin
// @contact.email matrenin@ukr.net
func main() {
	port := os.Getenv("PORT")

	db.Connect()

	router := gin.New()
	router.Use(gin.Logger())

	config := cors.Config{
		AllowMethods: []string{"PUT", "PATCH", "POST", "DELETE", "GET", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type"},
		MaxAge:       12 * time.Hour,
	}
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/candidate", services.CandidateList)
	router.PUT("/candidate", services.CandidateCreate)
	router.DELETE("/candidate", services.CandidatesDelete)
	router.GET("/candidate/:id", services.CandidateItem)
	router.POST("/candidate/:id", services.CandidateUpdate)
	router.DELETE("/candidate/:id", services.CandidateDelete)
	router.POST("/candidate/:id/status", services.CandidateStatusUpdate)

	router.Run(":" + port)
}
