package datastore

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type firestoreDB struct{}

//Define the path to the json file required for firestore
const (
	credentialsPath string = "GOOGLE_APPLICATION_CREDENTIALS_PATH"
)

//NewFirestoreDB will return the db object
func NewFirestoreDB() Database {
	return &firestoreDB{}
}

//createFirestoreClient is a local function that will return you an client to perform the queries
func createFirestoreClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform credentials.
	options := option.WithCredentialsFile(credentialsPath)

	//Create the app and handle errors
	app, err := firebase.NewApp(ctx, nil, options)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	//Create the client and handle errors
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return client
}

//GetAll function will return all the items in a given table (collection for Firestore)
func (db *firestoreDB) GetAll(tableName string) ([]map[string]interface{}, error) {
	//Create the client and closing it after we are done with it.
	ctx := context.Background()
	client := createFirestoreClient(ctx)
	defer client.Close()

	//Empty slice of map interface to get all the items in any given collection
	docs := []map[string]interface{}{}
	iter := client.Collection(tableName).Documents(ctx)

	//Loop through the items to store them in the map
	for {
		doc, err := iter.Next()

		//Handle errors and finishin the loop
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to get all elements from %v. \nMessage: %v", tableName, err)
			return nil, err
		}

		//Single map to store the current document
		dict := map[string]interface{}{}

		//Loop through all the keys in the document and assign them the correct type
		for key, value := range doc.Data() {

			switch value.(type) {
			case string:
				dict[key] = value.(string)
			case int64:
				dict[key] = value.(int64)
			case float64:
				dict[key] = value.(float64)
			}

		}

		//Return the id too (it's not a field in the current structure)
		dict["id"] = doc.Ref.ID

		//Add the document to the slice
		docs = append(docs, dict)

	}

	return docs, nil

}

//GetItemByID returns only a document specified by the id
func (db *firestoreDB) GetItemByID(tableName string, id string) (map[string]interface{}, error) {
	//Create the client and closing it after we are done with it.
	ctx := context.Background()
	client := createFirestoreClient(ctx)
	defer client.Close()

	//Get the specific document and handle the error
	doc, err := client.Collection(tableName).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	//Get the data from the document and the Id too (it's not a field in the current structure)
	data := doc.Data()
	data["id"] = doc.Ref.ID

	return data, nil
}

//DeleteItem will remove a document from the collection
func (db *firestoreDB) DeleteItem(tableName string, id string) error {
	//Create the client and closing it after we are done with it.
	ctx := context.Background()
	client := createFirestoreClient(ctx)
	defer client.Close()

	//Delete the document in the firestore collection and handle the error if any.
	//NOTE: Currently firestore will not return an error even if the collection doesn't exist.
	_, err := client.Collection(tableName).Doc(id).Delete(ctx)
	if err != nil {
		return err
	}

	return nil

}

func (db *firestoreDB) UpdateItem(tableName string, id string, item map[string]interface{}) (map[string]interface{}, error) {
	//Create the client and closing it after we are done with it.
	ctx := context.Background()
	client := createFirestoreClient(ctx)
	defer client.Close()

	//Update the document with the new item. It will only overwrite the passed parameters and will keep the others not shared.
	_, err := client.Collection(tableName).Doc(id).Set(ctx, item, firestore.MergeAll)

	//Handle errors
	if err != nil {
		return nil, err
	}

	//Return the new updated item
	var result map[string]interface{}

	result, err = db.GetItemByID(tableName, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//AddItem will add a document to the specified collection
func (db *firestoreDB) AddItem(tableName string, item map[string]interface{}) (map[string]interface{}, error) {
	//Create the client and closing it after we are done with it.
	ctx := context.Background()
	client := createFirestoreClient(ctx)
	defer client.Close()

	//Add the document to the collection, get the id and handle the error if any.
	ref, _, err := client.Collection(tableName).Add(ctx, item)

	if err != nil {
		return nil, err
	}

	//Return the document id too.
	item["id"] = ref.ID

	return item, nil
}
