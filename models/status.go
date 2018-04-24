package models

type Status struct {
	Status string `json:"status" bson:"status"`
	Reason string `json:"reason" bson:"reason"`
}
