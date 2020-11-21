package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Todo struct (Model)
type Todo struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Status    string `json:"status"`
	IsDeleted bool   `json:"isDeleted"`
}

var todos []Todo

// GET /todos
func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// GET /todos/{id}
func getTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range todos {
		if id, err := strconv.Atoi(params["id"]); err == nil && item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Todo{})
}

func main() {
	router := mux.NewRouter()

	// TODO get data from DB
	todos = append(todos, Todo{ID: 10, Task: "Wash dishes", Status: "pending", IsDeleted: false})
	todos = append(todos, Todo{ID: 20, Task: "Make report", Status: "pending", IsDeleted: false})

	// Routes
	router.HandleFunc("/todos", getTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", getTodo).Methods("GET")
	// router.HandleFunc("/todos", createTodo).Methods("POST")
	// router.HandleFunc("/todos/{id}/done", markTaskDone).Methods("PUT")
	// router.HandleFunc("/todos/{id}/pending", markTaskPending).Methods("PUT")
	// router.HandleFunc("/todos/{id}/{task}", updateTask).Methods("PUT")
	// router.HandleFunc("/todos/{id}", softDeleteTask).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":3000", router))

}
