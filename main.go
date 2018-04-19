package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"gopkg.in/mgo.v2/bson"
)

type Candidate struct {
	Name string
}

func main() {
	port := os.Getenv("PORT")
	mongoHost := os.Getenv("M_HOST")
	mongoDb := os.Getenv("M_DB")
	mongoUser := os.Getenv("M_USER")
	mongoPassword := os.Getenv("M_PASS")

	if port == "" || mongoHost == "" || mongoDb == "" || mongoUser == "" || mongoPassword == "" {
		log.Fatal("Environment vars must be set")
	}

	mongoDialInfo := &mgo.DialInfo{
		Addrs:    []string{mongoHost},
		Database: mongoDb,
		Username: mongoUser,
		Password: mongoPassword,
	}

	session, err := mgo.DialWithInfo(mongoDialInfo)

	if err != nil {
		log.Fatal(err)
	}

	database := session.DB("heroku_27s1h5cn")

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", func(context *gin.Context) {

		c := database.C("Col1")
		result := Candidate{}

		err = c.Find(bson.M{"name": "Some name"}).One(&result)

		context.JSON(200, result)
	})

	router.GET("/add", func(context *gin.Context) {
		c := database.C("Col1")

		err = c.Insert(&Candidate{Name: "Some name"})

		if err != nil {
			panic(err)
		}

		context.JSON(200, gin.H{
			"status": "OK",
		})
	})

	router.Run(":" + port)
}
