package main

import (
	"fmt"
)

func main() {

	name := "Denis"
	course := "Golang"

	fmt.Println("\nHi", name, "you are currently watching", course)

	fmt.Println("\nCourse was change to", changeCourse(&course)) // & sends the memory address
}

func changeCourse(course *string) string { // receives the contents of the var
	*course = "Docker deep dive" //assign the contents
	fmt.Println("\nChanging Course to", *course)

	return *course //return the contents
}
