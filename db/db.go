package db

import (
	"github.com/globalsign/mgo"
	"log"
	"os"
)

var (
	Database *mgo.Database
)

func Connect() {
	mongoHost, mongoDb, mongoUser, mongoPassword :=
		os.Getenv("M_HOST"), os.Getenv("M_DB"), os.Getenv("M_USER"), os.Getenv("M_PASS")

	if mongoHost == "" || mongoDb == "" || mongoUser == "" || mongoPassword == "" {
		log.Fatal("MDB vars must be set")
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

	database := session.DB(mongoDb)

	Database = database
}
