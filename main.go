package main

import (
  "context"
  "log"
  "fmt"
  "time"
  "github.com/andrecrts/go-bootcamp/infrastructure/router"
  "github.com/andrecrts/go-bootcamp/domain/model"
	"github.com/labstack/echo/v4"

  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
   "go.mongodb.org/mongo-driver/mongo"
   "go.mongodb.org/mongo-driver/mongo/options"
   // "go.mongodb.org/mongo-driver/mongo/readpref"
)


func insertProduct(client *mongo.Client, productId string, name string ) {
  product := model.Product{productId, name}

  collection := client.Database("bootcamp").Collection("products")

  insertResult, err := collection.InsertOne(context.TODO(), product)

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Inserted post with ID:", insertResult.InsertedID)
}


func getPost(client *mongo.Client, id primitive.ObjectID) {

  collection := client.Database("bootcamp").Collection("products")

  filter := bson.D{}

  var product model.Product

  err := collection.FindOne(context.TODO(), filter).Decode(&product)

  if err != nil {

  log.Fatal(err)

  }



  fmt.Println("Found product with name ", product.Name)

}

func handleRequests() {

  client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
  if err != nil {
    log.Fatal(err)
  }

  ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
  err = client.Connect(ctx)
  if err != nil {
    log.Fatal(err)
  }

  defer client.Disconnect(ctx)

  // insertProduct(client, "Test", "123")
  oid, err := primitive.ObjectIDFromHex("5fb760418d726a2670df8b88")
  getPost(client, oid)

  e := echo.New()
  e = router.NewRouter(e)

	if err := e.Start(":1000"); err != nil {
		log.Fatalln(err)
	}
}

func main() {
  handleRequests()
}
