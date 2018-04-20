package services

import (
	"candidate-avalanche-api/db"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Candidate struct {
	Name string
}

func getCollection() *mgo.Collection {
	return db.Database.C("Col1")
}

func CandidateGet(context *gin.Context) {
	collection := getCollection()
	result := Candidate{}

	var err = collection.Find(bson.M{}).One(&result)

	if err != nil {
		panic(err)
	}

	context.JSON(200, result)
}

func CandidateCreate(context *gin.Context) {
	c := db.Database.C("Col1")

	var err = c.Insert(&Candidate{Name: "Some name"})

	if err != nil {
		panic(err)
	}


	context.JSON(200, gin.H{
		"status": "OK",
	})
}
