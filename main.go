package main

import (
	"routes"
	"utils"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := routes.NewRouter()

	port := utils.GetEnvVar("PORT")
	fmt.Println("Server listening in port" + port)
	server := http.ListenAndServe(port, router)

	log.Fatal(server)
}
