package services

import (
	"candidate-avalanche-api/db"
	"candidate-avalanche-api/models"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func getCollection() *mgo.Collection {
	return db.Database.C("Col1")
}

func CandidateList(context *gin.Context) {
	collection := getCollection()

	var result []models.Candidate
	err := collection.Find(bson.M{}).All(&result)

	if err != nil {
		panic(err)
	}

	context.JSON(200, result)
}

func CandidateItem(context *gin.Context) {
	collection := getCollection()
	result := models.Candidate{}

	id := context.Param("id")

	if !bson.IsObjectIdHex(id) {
		context.JSON(400, "Invalid ID")
		return
	}

	err := collection.FindId(bson.ObjectIdHex(id)).One(&result)

	if err != nil {
		context.JSON(400, "Candidate not found.")
		return
	}

	context.JSON(200, result)
}

func CandidateCreate(context *gin.Context) {
	collection := getCollection()

	name := context.PostForm("name")

	model := models.Candidate{ID: bson.NewObjectId(), Name: name}
	err := collection.Insert(&model)

	if err != nil {
		panic(err)
	}

	context.JSON(200, model)
}
