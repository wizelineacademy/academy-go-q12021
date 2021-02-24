package db

import (
	"utils"
	"gopkg.in/mgo.v2"
)

func GetSession() *mgo.Session {
	mongoUri := utils.GetEnvVar("MONGO_URL")
	session, err := mgo.Dial(mongoUri)

	if err != nil {
		panic(err)
	}

	return session
}