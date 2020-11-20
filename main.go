package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/josedejesusAmaya/golang-bootcamp-2020/infrastructure/router"
)

func main() {
	http.HandleFunc("/api/read", router.HandleRequest)

	fmt.Println("Service is running")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatalf("Failed to listen on port 80: %v", err)
		return
	}
}
