package main

import (
	"router"
	"utils"
	"fmt"
	"log"
	"net/http"
)

func main() {
	rt := router.NewRouter()
	port := utils.GetEnvVar("PORT")

	fmt.Println("Server listening in port" + port)

	server := http.ListenAndServe(port, rt)

	log.Fatal(server)
}
