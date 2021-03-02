package hello

import (
	"fmt"
	"net/http"
	"bootcamp/usecase/hello"
)

/*
HelloWorld prints a hellow world
*/
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	msg := hello.HelloWorld()
	fmt.Fprintf(w, msg)
}