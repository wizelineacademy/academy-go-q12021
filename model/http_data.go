package model

import (
	"fmt"
)

type HttpData struct {
	Url    string
	Method string
	Body   string
}

func (hd HttpData) String() string {
	return fmt.Sprintf("{Url: '%v', Method: '%v', Body: '%v'}", hd.Url, hd.Method, hd.Body)
}
