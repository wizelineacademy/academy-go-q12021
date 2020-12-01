package main

import "fmt"

// User user type
type User struct {
	ID   int
	Name string
}

func main() {
	var (
		u1    User
		u2    User
		users = []*User{&u1, &u2}
	)
	u1.ID = 0
	u1.Name = "Juan"
	u2.ID = 1
	u2.Name = "Maria"
	fmt.Println(users)
	for _, user := range users {
		fmt.Println(user)
	}

	//declaracion 1
	rob := User{12,"Rob"}
	fmt.Println("%+v\n", rob) // imprime toda la estructura
	fmt.Println("%v\n", rob.Name) // email val


}


