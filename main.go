package main

import(
    "net/http"
    "log"
    "io/ioutil"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        log.Println("Hello World")
        d, _ := ioutil.ReadAll(r.Body)

        log.Printf("Data %s\n", d)
    }) 

    http.ListenAndServe(":9090", nil)
}
