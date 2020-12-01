package datastore

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewDb() (*mongo.Client, context.Context) {
	fmt.Println("Connecting mongo...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://dbHX5amongo:5ZrpEm0hHX5aEG5pE@cluster0.6uk2m.mongodb.net/?retryWrites=true&w=majority",
	))
	check(err)
	//defer close(client)

	// Check the connections
	err = client.Ping(context.TODO(), nil)
	check(err)

	err = client.Ping(ctx, readpref.Primary())
	check(err)
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	check(err)
	fmt.Println(databases)

	fmt.Println("MongoDB is ready!")

	return client, ctx
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

