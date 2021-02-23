module bootcamp

go 1.16

require (
	github.com/gorilla/mux v1.8.0
	models v0.0.0-00010101000000-000000000000
	modules v0.0.0-00010101000000-000000000000 // indirect
	routes v0.0.0-00010101000000-000000000000 // indirect
	utils v0.0.0-00010101000000-000000000000
)

replace models => ./models

replace utils => ./utils

replace modules => ./modules

replace routes => ./routes
