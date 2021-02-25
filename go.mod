module bootcamp

go 1.16

require (
	db v0.0.0-00010101000000-000000000000 // indirect
	model v0.0.0-00010101000000-000000000000 // indirect
	modules v0.0.0-00010101000000-000000000000 // indirect
	routes v0.0.0-00010101000000-000000000000
	utils v0.0.0-00010101000000-000000000000
)

replace model => ./domain/model

replace utils => ./utils

replace modules => ./modules

replace routes => ./routes

replace db => ./db
