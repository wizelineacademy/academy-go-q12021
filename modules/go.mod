module modules

go 1.16

replace models => ../models

replace utils => ../utils

require (
	github.com/gorilla/mux v1.8.0
	models v0.0.0-00010101000000-000000000000
	utils v0.0.0-00010101000000-000000000000
)
