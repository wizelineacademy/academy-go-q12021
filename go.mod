module bootcamp

go 1.16

require (
	model v0.0.0-00010101000000-000000000000 // indirect
	modules v0.0.0-00010101000000-000000000000 // indirect
	router v0.0.0-00010101000000-000000000000 // indirect
	service/db v0.0.0-00010101000000-000000000000 // indirect
	utils v0.0.0-00010101000000-000000000000
)

replace model => ./domain/model

replace utils => ./utils

replace modules => ./modules

replace service/db => ./service/db

replace router => ./router
