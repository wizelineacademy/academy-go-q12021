go test ./infrastructure/newsapi -v -cover   

go tool cover -html=coverage.out
go test ./infrastructure/newsapi -v -cover -coverprofile=coverage.out

mockgen -package mocks -destination utils/mocks/httpclient_mock.go github.com/jesus-mata/academy-go-q12021/infrastructure HTTPClient

 go test ./...