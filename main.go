package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

func ReadFile(name string) string {
	f, err := os.Open(name)
	check(err)
	o1, err := f.Seek(2, 3)
	check(err)
	b1 := make([]byte, 2)
	n1, err := f.Read(b1)
	fmt.Printf("%v bytes @ %d: ", b1, o1)
	//	return string(b1[:n1])
	//return string(fmt.Sprint(n1))
	return string(b1[:n1])
	//    fmt.Print(string(dat))
}
