generate-mock:
	go generate -v ./...

/Users/adfernandez/go/bin/mockgen -source=service/csv/csv.go -destination=service/mock/csv_mock.go -package=mock