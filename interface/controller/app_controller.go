package controller

// AppController This contain all the controller interfaces that this software accepts
type AppController struct {
	Digimon interface{ DigimonController }
}
