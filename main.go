package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
		//	log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(rw, "There's an error", http.StatusBadRequest)
			return
		}

		//log.Println(name)
		log.Println(d)
		content := ReadFile(string(d))

		fmt.Fprintf(rw, "Hello %s", content)
	})

	http.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye World")

	})

	http.ListenAndServe(":9090", nil)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
