// Resty Client, service package
package services

import "C"
import (
	"strconv"
	"time"

	"github.com/ruvaz/golang-bootcamp-2020/config"

	"github.com/go-resty/resty/v2"
)

// Client struct
type Client struct {
	client *resty.Client
}

// NewClient: Resty client, Return: *Client
func NewClient() *Client {
	//err := config.ReadConfig("config")
	//if err != nil {
	//	log.Fatal(err)
	//}
	host := config.C.Server.Address + ":" + strconv.Itoa(config.C.Server.Port)
	timeout := config.C.Server.Timeout * time.Second

	client := resty.New().
		SetHostURL(host).
		SetTimeout(timeout)

	return &Client{client: client}
}
