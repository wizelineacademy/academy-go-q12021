package handlers

import (
	"fmt"
	"net/http"

	"github.com/alexis-aguirre/golang-bootcamp-2020/registry"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	ui := registry.NewUserInteractor()
	u, _ := ui.Get(nil)
	for _, user := range u {
		fmt.Println(user)
	}
	w.WriteHeader(http.StatusFound)
	w.Write([]byte("GetUsers"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	ui := registry.NewUserInteractor()
	u, _ := ui.Create(nil)
	fmt.Println(&u)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("UserCreated"))
}
