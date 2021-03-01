module FirstDeliverable

go 1.16

require (
	example.com/me/csv v0.0.0
	github.com/gorilla/mux v1.8.0 // indirect
)

replace example.com/me/csv => ../FirstDeliverable/CSV
