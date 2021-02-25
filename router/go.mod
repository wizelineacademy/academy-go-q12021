module router

go 1.16

require (
	github.com/gorilla/mux v1.8.0
	model v0.0.0-00010101000000-000000000000 // indirect
)

replace model => ../domain/model

replace modules => ../modules
