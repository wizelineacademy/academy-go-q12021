package main

import (
	"fmt"

	"golang-bootcamp-2020/interface/controller"
	"golang-bootcamp-2020/interface/presenter"
)

func main() {
	fmt.Printf(">First Deliverable\n>>Csv Reader\n")
	students := controller.GetStudents()
	presenter.ResponseStudents(students)
}
