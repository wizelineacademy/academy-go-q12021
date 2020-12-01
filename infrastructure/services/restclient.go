package services

import (
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"golang-bootcamp-2020/config"
	"golang-bootcamp-2020/domain/model"
)


type Client struct{
	client *resty.Client
}

// GET new client Resty
func  NewClient() *Client {
	var (
		host    = config.C.Server.Address + ":" + strconv.Itoa(config.C.Server.Port)
		timeout = config.C.Server.Timeout * time.Second
	)

	client := resty.New().
		SetHostURL(host).
		SetTimeout(timeout)

	return &Client{client: client}
}

func (c *Client) GetStudentsFromCsv() ([]model.Student, error)  {
	c.client.R()
	// endpoint

	return []model.Student{}, nil
}
