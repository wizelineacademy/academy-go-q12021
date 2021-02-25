module bootcamp

go 1.16

require (
	model v0.0.0-00010101000000-000000000000 // indirect
	modules v0.0.0-00010101000000-000000000000 // indirect
	routes v0.0.0-00010101000000-000000000000
	service/db v0.0.0-00010101000000-000000000000 // indirect
	utils v0.0.0-00010101000000-000000000000
)

replace model => ./domain/model

replace utils => ./utils

replace modules => ./modules

replace routes => ./routes

replace service/db => ./service/db
