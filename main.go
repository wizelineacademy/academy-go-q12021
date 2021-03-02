package main

import (
	"fmt"
	"log"
	"net/http"
	"bootcamp/router"
	"bootcamp/utils"
)

func main() {
	rt := router.NewRouter()
	port, _ := utils.GetEnvVar("PORT")

	fmt.Println("Server listening in port" + port)

	server := http.ListenAndServe(port, rt)

	log.Fatal(server)
}
