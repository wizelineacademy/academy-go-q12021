module bootcamp

go 1.16

require (
	bootcamp/domain/model v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
	github.com/spf13/viper v1.7.1
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)

replace bootcamp/router => ./router

replace bootcamp/domain/model => ./domain/model
