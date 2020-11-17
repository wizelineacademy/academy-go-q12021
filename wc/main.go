package main

import "fmt"


func printValue() func() int {
	i := 0
    return func() (r int) {
        r = i
		i++
        return 
    }
}

func main() {
	f := printValue()
	for i := 0; i < 10; i++ {
		fmt.Println(i, f())
	}
}
