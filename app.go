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

func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contetn-Type", "application/json")
	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo)
}

func softDeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contetn-Type", "application/json")
	params := mux.Vars(r)
	var todo Todo
	for idx, item := range todos {
		if id, err := strconv.Atoi(params["id"]); err == nil && item.ID == id {
			todos[idx].IsDeleted = true
			todo = todos[idx]
			// todos = append(todos[:idx], todos[idx+1:]...) // hard delete
			break
		}
	}
	json.NewEncoder(w).Encode(todo)
}

func markTodoDone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contetn-Type", "application/json")
	params := mux.Vars(r)
	var todo Todo
	for idx, item := range todos {
		if id, err := strconv.Atoi(params["id"]); err == nil && item.ID == id {
			todos[idx].Status = "done"
			todo = todos[idx]
			break
		}
	}
	json.NewEncoder(w).Encode(todo)
}

func markTodoPending(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contetn-Type", "application/json")
	params := mux.Vars(r)
	var todo Todo
	for idx, item := range todos {
		if id, err := strconv.Atoi(params["id"]); err == nil && item.ID == id {
			todos[idx].Status = "pending"
			todo = todos[idx]
			break
		}
	}
	json.NewEncoder(w).Encode(todo)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contetn-Type", "application/json")
	params := mux.Vars(r)
	var todo Todo
	for idx, item := range todos {
		if id, err := strconv.Atoi(params["id"]); err == nil && item.ID == id {
			todos[idx].Task = params["task"]
			todo = todos[idx]
			break
		}
	}
	json.NewEncoder(w).Encode(todo)
}

func main() {
	router := mux.NewRouter()

	// TODO get data from DB
	todos = append(todos, Todo{ID: 10, Task: "Wash dishes", Status: "pending", IsDeleted: false})
	todos = append(todos, Todo{ID: 20, Task: "Make report", Status: "pending", IsDeleted: false})

	// Routes
	router.HandleFunc("/todos", getTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", getTodo).Methods("GET")
	router.HandleFunc("/todos", createTodo).Methods("POST")
	router.HandleFunc("/todos/{id}/done", markTodoDone).Methods("PUT")
	router.HandleFunc("/todos/{id}/pending", markTodoPending).Methods("PUT")
	router.HandleFunc("/todos/{id}/{task}", updateTask).Methods("PUT")
	router.HandleFunc("/todos/{id}", softDeleteTodo).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":3000", router))

}
