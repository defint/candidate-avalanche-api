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

func getIdFromContext(context *gin.Context) bson.ObjectId {
	id := context.Param("id")

	if !bson.IsObjectIdHex(id) {
		context.JSON(400, "Invalid ID")
		return ""
	}

	return bson.ObjectIdHex(id)
}

func makeModel(id bson.ObjectId, context *gin.Context) models.Candidate {
	name := context.PostForm("name")
	model := models.Candidate{ID: id, Name: name}
	return model
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

	id := getIdFromContext(context)

	if id == "" {
		return
	}

	err := collection.FindId(id).One(&result)

	if err != nil {
		context.JSON(400, "Candidate not found.")
		return
	}

	context.JSON(200, result)
}

func CandidateCreate(context *gin.Context) {
	collection := getCollection()
	id := bson.NewObjectId()

	model := makeModel(id, context)
	err := collection.Insert(&model)

	if err != nil {
		panic(err)
	}

	context.JSON(200, model)
}

func CandidateUpdate(context *gin.Context) {
	collection := getCollection()
	id := getIdFromContext(context)

	if id == "" {
		return
	}

	model := makeModel(id, context)
	err := collection.UpdateId(id, &model)

	if err != nil {
		panic(err)
	}

	context.JSON(200, model)
}

func CandidateDelete(context *gin.Context) {
	collection := getCollection()
	id := getIdFromContext(context)

	if id == "" {
		return
	}

	err := collection.RemoveId(id)

	if err != nil {
		panic(err)
	}

	context.JSON(200, id)
}
