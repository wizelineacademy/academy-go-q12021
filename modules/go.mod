module modules

go 1.16

replace model => ../domain/model

replace utils => ../utils

replace service/db => ../service/db

require (
	github.com/gorilla/mux v1.8.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22 // indirect
	service/db v0.0.0-00010101000000-000000000000 // indirect
	utils v0.0.0-00010101000000-000000000000
)
