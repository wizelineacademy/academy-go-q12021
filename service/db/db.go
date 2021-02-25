package db

import (
	"utils"
	"gopkg.in/mgo.v2"
)

func GetSession() *mgo.Collection {
	mongoUri := utils.GetEnvVar("MONGO_URL")
	databaseName := utils.GetEnvVar("DB_NAME")
	collectionName := utils.GetEnvVar("COLLECTION")

	session, err := mgo.Dial(mongoUri)

	if err != nil {
		panic(err)
	}

	return session.DB(databaseName).C(collectionName)
}