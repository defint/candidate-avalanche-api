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
	collection := getCollection()
	result := models.Candidate{}
	_ = collection.FindId(id).One(&result)

	model := models.Candidate{}
	context.Bind(&model)
	model.ID = id
	model.History = result.History
	model.Status = result.Status

	return model
}

// ShowCandidates godoc
// @Summary Show a list of candidates
// @Description Show a list of candidates
// @ID candidates-list
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Candidate
// @Failure 400 {string} string
// @Router /candidate [get]
func CandidateList(context *gin.Context) {
	collection := getCollection()

	var result []models.Candidate
	err := collection.Find(bson.M{}).All(&result)

	if err != nil {
		panic(err)
	}

	context.JSON(200, result)
}

// ShowCandidate godoc
// @Summary Show a candidate
// @Description Show a candidate
// @ID candidate-item
// @Accept  json
// @Produce  json
// @Param id path string true "Candidate ID"
// @Success 200 {object} models.Candidate
// @Failure 400 {string} string
// @Router /candidate/{id} [get]
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

// CreateCandidate godoc
// @Summary Create a candidate
// @Description Create a candidate
// @ID candidate-create
// @Accept  json
// @Produce  json
// @Param candidate body models.Candidate true "Candidate"
// @Success 200 {object} models.Candidate
// @Failure 400 {string} string
// @Router /candidate [put]
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

// UpdateCandidate godoc
// @Summary Update a candidate
// @Description Update a candidate
// @ID candidate-update
// @Accept  json
// @Produce  json
// @Param id path string true "Candidate ID"
// @Param candidate body models.Candidate true "Candidate"
// @Success 200 {object} models.Candidate
// @Failure 400 {string} string
// @Router /candidate/{id} [post]
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

// DeleteCandidate godoc
// @Summary Delete a candidate
// @Description Delete a candidate
// @ID candidate-delete
// @Accept  json
// @Produce  json
// @Param id path string true "Candidate ID"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Router /candidate/{id} [delete]
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

// DeleteCandidates godoc
// @Summary Delete all candidates
// @Description Delete all candidates
// @ID candidates-delete
// @Accept  json
// @Produce  json
// @Success 200 {string} string
// @Failure 400 {string} string
// @Router /candidate [delete]
func CandidatesDelete(context *gin.Context) {
	collection := getCollection()

	_, err := collection.RemoveAll(nil)

	if err != nil {
		panic(err)
	}

	context.JSON(200, nil)
}

// UpdateCandidateStatus godoc
// @Summary Update a status of candidate
// @Description Update a status of candidate
// @ID candidate-status-update
// @Accept  json
// @Produce  json
// @Param id path string true "Candidate ID"
// @Param status body models.Status true "Status"
// @Success 200 {object} models.Candidate
// @Failure 400 {string} string
// @Router /candidate/{id}/status [post]
func CandidateStatusUpdate(context *gin.Context) {
	collection := getCollection()
	id := getIdFromContext(context)

	if id == "" {
		return
	}

	result := models.Candidate{}
	err := collection.FindId(id).One(&result)

	if err != nil {
		context.JSON(400, "Candidate not found.")
		return
	}

	statusModel := models.Status{}
	context.Bind(&statusModel)
	result.History = append(result.History, models.History{
		Reason:     statusModel.Reason,
		StatusFrom: result.Status,
		StatusTo:   statusModel.Status,
	})

	result.Status = statusModel.Status
	err = collection.UpdateId(id, &result)

	if err != nil {
		context.JSON(400, "Server error")
		return
	}

	context.JSON(200, result)
}
