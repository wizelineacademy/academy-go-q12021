/**
Resty Client
*/
package services

import (
	"time"

	"golang-bootcamp-2020/config"

	"github.com/go-resty/resty/v2"
)

// client struct
type Client struct {
	client *resty.Client
}

// GET new client Resty
func NewClient() *Client {
	var (
		host    = config.C.GetServerAddr()
		timeout = config.C.Server.Timeout * time.Second
	)

	client := resty.New().
		SetHostURL(host).
		SetTimeout(timeout)

	return &Client{client: client}
}
