package models

import "github.com/globalsign/mgo/bson"

type Candidate struct {
	ID   bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
}
