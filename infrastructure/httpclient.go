package infrastructure

import "net/http"

//go:generate mockgen -package mocks -destination $ROOTDIR/mocks/$GOPACKAGE/mock_$GOFILE . HTTPClient
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
