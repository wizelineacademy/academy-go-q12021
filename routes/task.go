//Package routes -> THIS FILE WAS JUST AN EXERCISE STILL INCLUDED TO SHOW THE WORK MADE ON LOOPS THAT I DIDNT HAVE ON THE NORMAL REQUESTS
package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// task structure for the requests
type task struct {
	ID   int
	Name string
	Cont string
}

type allTasks []task

var tasks = allTasks{
	{
		ID:   1,
		Name: "Task One",
		Cont: "Test task",
	},
}

//GetTasks returns all tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)

}

//GetTask returns one task
func GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}
	for _, task := range tasks {
		if task.ID == taskID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
		}
	}
}

//CreateTasks function to create a task
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

//DeleteTask function to delete a task
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}
	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Fprintf(w, "task with id %v was deleted", taskID)
		}
	}
}

//PutTask function to modify task
func PutTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
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
		if task.ID == taskID {
			tasks[i].Name = updatedTask.Name
			tasks[i].Cont = updatedTask.Cont
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(tasks[i])
		}
	}
}
