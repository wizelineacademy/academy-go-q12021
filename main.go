package main

import(
    "net/http"
    "log"
    "io/ioutil"
    "fmt"
)

func main() {
    http.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
        log.Println("Hello World")
        d, err := ioutil.ReadAll(r.Body)

        if err != nil {
            http.Error(rw, "There's an error", http.StatusBadRequest)
            return 
        }

        fmt.Fprintf(rw, "Hello %s", d)
    }) 

    http.ListenAndServe(":9090", nil)
}
