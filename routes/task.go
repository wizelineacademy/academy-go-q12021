package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type task struct {
	ID   int    `json:ID`
	Name string `json:Name`
	Cont string `json:Cont`
}

type allTasks []task

var tasks = allTasks{
	{
		ID:   1,
		Name: "Task One",
		Cont: "Test task",
	},
}

//Example returns Success
func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)

}

func GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}
	for _, task := range tasks {
		if task.ID == taskId {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
		}
	}
}

func CreateTasks(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a valid task")
	}
	json.Unmarshal(reqBody, &newTask)

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newTask)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}
	for i, task := range tasks {
		if task.ID == taskId {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Fprintf(w, "task with id %v was deleted", taskId)
		}
	}
}

func PutTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])
	var updatedTask task

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a valid task")
	}
	json.Unmarshal(reqBody, &updatedTask)

	for i, task := range tasks {
		if task.ID == taskId {
			tasks[i].Name = updatedTask.Name
			tasks[i].Cont = updatedTask.Cont
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(tasks[i])
		}
	}
}
