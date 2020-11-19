package util

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
)

func GetConnDB() *mongo.Client {
	host := "localhost"
	port := 27017
	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port))
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connections
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Congratulations, you're already connected to MongoDB!")
	return client
}

func Test_getCSVCodes(t *testing.T) {
	// mock de mongo +plus
	client := GetConnDB()
	assert.Condition(t, func() bool {
		if len(GetAndSave(client).InsertedIDs) != 0 {
			return true
		}
		return false
	}, "The process has inserted  rows to the db")
}

func Test_searchZipCodes(t *testing.T) {
	client := GetConnDB()
	//assert.NotEqual(t, ) mejorarlo
	assert.Condition(t, func() bool {
		r := SearchZipCodes("97306", client)
		fmt.Println(r)
		if len(r) != 0 {
			return true
		}
		return false
	}, "Result of search > 0")
}

func Test_dropZipCodes(t *testing.T) {
	client := GetConnDB()
	assert.Equal(t, true, dropZipCodes(client), "")
}

func Test_edoToISO(t *testing.T) {
	assert.Equal(t, "MX-CMX", edoToISO(9), "Success ")
}
