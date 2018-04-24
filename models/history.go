package models

type History struct {
	StatusFrom string `json:"status_from" bson:"status_from"`
	StatusTo   string `json:"status_to" bson:"status_to"`
	Reason     string `json:"reason" bson:"reason"`
	Date       string `json:"date" bson:"date"`
}
