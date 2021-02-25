package hello

import (
	"fmt"
	"bootcamp/usecase/hello"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	msg := hello.HelloWorld()
	fmt.Fprintf(w, msg)
}