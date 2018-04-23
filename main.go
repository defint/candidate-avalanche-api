package main

import (
	"os"

	"candidate-avalanche-api/db"
	_ "candidate-avalanche-api/docs"
	"candidate-avalanche-api/services"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/candidate", services.CandidateList)
	router.PUT("/candidate", services.CandidateCreate)
	router.GET("/candidate/:id", services.CandidateItem)
	router.POST("/candidate/:id", services.CandidateUpdate)
	router.DELETE("/candidate/:id", services.CandidateDelete)

	router.Run(":" + port)
}
