/**
Resty Client
*/
package services

import (
	"time"

	"golang-bootcamp-2020/config"

	"github.com/go-resty/resty/v2"
)

// Client struct
type Client struct {
	client *resty.Client
}

// NewClient new client Resty
func NewClient() *Client {
	host := config.C.GetServerAddr()
	timeout := config.C.Server.Timeout * time.Second

	client := resty.New().
		SetHostURL(host).
		SetTimeout(timeout)

	return &Client{client: client}
}
