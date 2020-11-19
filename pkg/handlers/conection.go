package handlers

import (
	"context"
	"fmt"
	"github.com/AlonSerrano/GolangBootcamp/pkg/util"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

type PerissonConnectorHandler struct {
	client *mongo.Client
}

func NewPerissonConnectorHandler() *PerissonConnectorHandler {
	p := PerissonConnectorHandler{client: GetConnDB()}
	return &p
}

func (hc *PerissonConnectorHandler) HandleSearchZipCodes(c echo.Context) error {
	zipCode := c.Param("zipCode")
	zipCodes := util.SearchZipCodes(zipCode, hc.client)
	return c.JSON(http.StatusOK, zipCodes)
}

func (hc *PerissonConnectorHandler) HandlePopulateZipCodes(c echo.Context) error {
	zipCodes := util.GetAndSave(hc.client)
	return c.JSON(http.StatusOK, zipCodes)
}

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
