package main

import (
	"net/http"

	"github.com/wizelineacademy/academy-go/config"
	"github.com/wizelineacademy/academy-go/constant"
	"github.com/wizelineacademy/academy-go/controller"
	"github.com/wizelineacademy/academy-go/data"
	"github.com/wizelineacademy/academy-go/model"
	"github.com/wizelineacademy/academy-go/service"
)

var serverPort = config.GetEnvVar(constant.PokemonSourceVarName)

func main() {

	for path, handler := range routes {
		http.HandleFunc(path, handler)
	}
	http.ListenAndServe(":"+serverPort, nil)
}
