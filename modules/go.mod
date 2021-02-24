module modules

go 1.16

replace models => ../models

replace utils => ../utils

replace db => ../db

require (
	github.com/gorilla/mux v1.8.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22 // indirect
	models v0.0.0-00010101000000-000000000000
	utils v0.0.0-00010101000000-000000000000
)
