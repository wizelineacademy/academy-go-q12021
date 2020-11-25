package presenter

import (
	"encoding/json"
	"fmt"

	"golang-bootcamp-2020/domain/model"
)

// ResponseStudents - Show students
func ResponseStudents(students []model.Student) {
	var count = len(students)

	//for _, v := range students {
	//	fmt.Println(v)
	//}

	slcB, _ := json.Marshal(students)
	fmt.Println(string(slcB))

	fmt.Printf("\n>>>Done ... %v rows were read.\n", count)
}
