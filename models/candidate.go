package models

import (
	"github.com/globalsign/mgo/bson"
)

type Candidate struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name     string        `json:"name" bson:"name"`
	Position string        `json:"position" bson:"position"`
	Salary   int           `json:"salary" bson:"salary"`
	Status   string        `json:"status" bson:"status"`
}
